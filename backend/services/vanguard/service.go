package vanguard

import (
	"context"
	"encoding/json"
	"errors"
	"etfinsight/generated/jet_gen/postgres/public/model"
	"etfinsight/services/fund"
	"etfinsight/utils/stringutils"
	"fmt"
	"github.com/jackc/pgx/v5"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	client     EtfIssuerClient
	figiClient FigiClient
	repo       Repository
}

type FigiResp struct {
	Warning string      `json:"warning"`
	Data    []FigiValue `json:"data"`
}
type FigiValue struct {
	Figi   string `json:"figi"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}

type FigiPayload struct {
	IdType       string       `json:"idType"`
	IdValue      string       `json:"idValue"`
	HoldingsItem HoldingsItem `json:"-"`
}

var cusipFigiMap = map[string]model.FigiMapping{}
var sedolFigiMap = map[string]model.FigiMapping{}

func NewService(client EtfIssuerClient, repo Repository, figiClient FigiClient) *Service {
	return &Service{client: client, repo: repo, figiClient: figiClient}
}

func (s *Service) UpsertFunds(ctx context.Context) error {
	ei, err := getVanguardExternalIds()
	if err != nil {
		return err
	}
	batchSize := 3
	batches := make([][]string, 0, (len(ei)+batchSize-1)/batchSize)

	for batchSize < len(ei) {
		ei, batches = ei[batchSize:], append(batches, ei[0:batchSize:batchSize])
	}
	if err != nil {
		return err
	}
	batches = append(batches, ei)
	cusipM, sedolM, err := s.repo.GetFigiMappings(ctx)
	if err != nil {
		return err
	}
	cusipFigiMap = cusipM
	sedolFigiMap = sedolM
	for _, batch := range batches {
		err := func() error {
			defer func() {
				a := recover()

				if a != nil {
					fmt.Println("Recovered", a)
				}
			}()
			return s.upsertFunds(ctx, batch)
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) upsertFunds(ctx context.Context, ei []string) error {
	fundingBytes, err := s.client.GetFunds(ei)
	if err != nil {
		return err
	}

	fundsResponse := FundsResponse{}
	err = json.Unmarshal(fundingBytes, &fundsResponse)
	if err != nil {
		return err
	}
	for i, f := range fundsResponse.Data.Funds {
		err := s.convertFund(ctx, f, fundsResponse.Data.PolarisAnalyticsHistories[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func getVanguardExternalIds() ([]string, error) {
	get, err := http.Get("https://www.vanguard.co.uk/professional/product?asset-class=ETF")
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	result, found := stringutils.GetStringInBetweenTwoString(string(b), "portIds&q;:&q;", "&q;")
	if !found {
		return nil, errors.New("no external ids found")
	}
	return strings.Split(result, ","), nil
}

func getISIN(identifiers []Identifier) string {
	for _, vhfli := range identifiers {
		if vhfli.AltId == "ISIN" {
			return vhfli.AltIdValue
		}
	}
	return ""
}

func (s *Service) convertFund(ctx context.Context, fund Fund, polarisHistory PolarisAnalyticsHistory) error {
	tx, err := s.repo.NewTransaction(ctx)
	if err != nil {
		return err
	}
	defer s.repo.RollBack(tx, ctx)
	holdings, fundHoldings := s.convertToHoldings(ctx, fund.Holdings.OriginalHoldings.Items, tx)
	ISIN := getISIN(fund.Profile.Identifiers)
	name := fund.Profile.FundFullName
	currency := fund.Profile.FundCurrency
	externalIdentifier := fund.Profile.PortId
	totalHoldings := fund.Holdings.OriginalHoldings.TotalHoldings
	var outstandingShares float64
	var effectiveDate time.Time
	for _, item := range polarisHistory.AnalyticsMonthly.Valuation.Fund.Items {
		for _, oscltshqty := range item.OSCLTSHQTY {
			outstandingShares = oscltshqty.OutstandingShare
			parsedDate, err := time.Parse("2006-01-02", oscltshqty.EffectiveDate)
			if err != nil {
				return err
			}
			effectiveDate = parsedDate
		}
	}
	var price float64
	if len(fund.PricingDetails.NavPrices.Items) > 0 {
		price = fund.PricingDetails.NavPrices.Items[0].Price
	}
	provider := "Vanguard"
	f := model.Fund{
		Name:               &name,
		Currency:           &currency,
		Isin:               &ISIN,
		ExternalIdentifier: &externalIdentifier,
		TotalHoldings:      &totalHoldings,
		Price:              &price,
		Provider:           &provider,
		OutstandingShares:  &outstandingShares,
		EffectiveDate:      &effectiveDate,
	}
	fundId, err := s.repo.UpsertFund(ctx, f, tx)
	if err != nil {
		return err
	}
	fundListings := convertToListings(fundId, fund.Profile.Listings)
	if len(fundListings) > 0 {
		err = s.repo.UpsertFundListings(ctx, fundListings, tx)
		if err != nil {
			return err
		}
	}
	if len(holdings) > 0 {
		var holdingsSlice []model.Holding
		for _, hv := range holdings {
			holdingsSlice = append(holdingsSlice, hv)
		}
		upsertedHoldings, err := s.repo.UpsertHoldings(ctx, holdingsSlice, tx)
		if err != nil {
			return err
		}
		var aggFundHoldings []model.FundHolding
		for fundHoldingFigi, fh := range fundHoldings {
			upsertedHoldingId := upsertedHoldings[fundHoldingFigi]
			fh.FundID = &fundId
			fh.HoldingID = &upsertedHoldingId
			aggFundHoldings = append(aggFundHoldings, fh)
		}
		batchSize := 5000
		batches := make([][]model.FundHolding, 0, (len(aggFundHoldings)+batchSize-1)/batchSize)

		for batchSize < len(aggFundHoldings) {
			aggFundHoldings, batches = aggFundHoldings[batchSize:], append(batches, aggFundHoldings[0:batchSize:batchSize])
		}

		batches = append(batches, aggFundHoldings)
		for _, batch := range batches {
			err = s.repo.UpsertFundHoldings(ctx, batch, tx)
		}
	}
	if err != nil {
		return err
	}
	fmt.Printf("Converted fund %s\n", fund.Profile.FundFullName)
	return tx.Commit(ctx)
}

func convertToListings(fundId uuid.UUID, listings []FundInformationListing) []model.FundListing {
	var fundListings []model.FundListing
	for _, listing := range listings {
		if len(listing.Identifiers) != 1 {
			continue
		}
		fundListings = append(fundListings, model.FundListing{
			FundID: &fundId,
			Ticker: &listing.Identifiers[0].AltIdValue,
		})
	}
	return fundListings
}

func (s *Service) convertToHoldings(ctx context.Context, HoldingItems []HoldingsItem, tx pgx.Tx) (map[string]model.Holding, map[string]model.FundHolding) {
	holdingsMap := map[string]model.Holding{}
	fundHoldingsMap := map[string]model.FundHolding{}

	batchSize := 100
	batches := make([][]HoldingsItem, 0, (len(HoldingItems)+batchSize-1)/batchSize)

	for batchSize < len(HoldingItems) {
		HoldingItems, batches = HoldingItems[batchSize:], append(batches, HoldingItems[0:batchSize:batchSize])
	}

	batches = append(batches, HoldingItems)
	for _, batch := range batches {
		err := s.convertBatch(ctx, batch, holdingsMap, fundHoldingsMap, tx)
		if err != nil {
			fmt.Println(err)
			return nil, nil
		}
	}

	return holdingsMap, fundHoldingsMap
}

func (s *Service) convertBatch(ctx context.Context,
	batch []HoldingsItem,
	holdingsMap map[string]model.Holding,
	fundHoldingsMap map[string]model.FundHolding,
	tx pgx.Tx) error {
	var figiPayload []FigiPayload
	for _, hi := range batch {
		if hi.SEDOL != nil {
			foundFigi, ok := sedolFigiMap[*hi.SEDOL]
			if ok {
				figiCopy := foundFigi.Figi
				h := model.Holding{Figi: &figiCopy}
				setHoldingValues(hi, &h)
				holdingsMap[figiCopy] = h
				addToFundHoldingsMap(figiCopy, hi, fundHoldingsMap)
			} else {
				if hi.CUSIP != nil {
					foundFigi, ok := cusipFigiMap[*hi.CUSIP]
					if ok {
						figiCopy := foundFigi.Figi
						h := model.Holding{Figi: &figiCopy}
						setHoldingValues(hi, &h)
						holdingsMap[figiCopy] = h
						addToFundHoldingsMap(figiCopy, hi, fundHoldingsMap)
					} else {
						figiPayload = append(figiPayload, FigiPayload{
							IdType:       "ID_SEDOL",
							IdValue:      *hi.SEDOL,
							HoldingsItem: hi,
						})
					}
				}
			}
		} else if hi.CUSIP != nil {
			foundFigi, ok := cusipFigiMap[*hi.CUSIP]
			if ok {
				figiCopy := foundFigi.Figi
				h := model.Holding{Figi: &figiCopy}
				setHoldingValues(hi, &h)
				holdingsMap[figiCopy] = h
				addToFundHoldingsMap(figiCopy, hi, fundHoldingsMap)
			} else {
				figiPayload = append(figiPayload, FigiPayload{
					IdType:       "ID_CUSIP",
					IdValue:      *hi.CUSIP,
					HoldingsItem: hi,
				})
			}
		}
	}
	if len(figiPayload) > 0 {
		// TODO Deal with currency
		figisResp, err := s.figiClient.GetFigi(figiPayload)
		if err != nil {
			return nil
		}
		var mappings []model.FigiMapping
		for i, figiR := range figisResp {
			if len(figiR.Data) == 0 || figiR.Warning != "" {
				var retryFigiPayload []FigiPayload
				if figiPayload[i].HoldingsItem.CUSIP != nil {
					retryCusipCopy := figiPayload[i].HoldingsItem.CUSIP
					retryFigiPl := FigiPayload{
						IdType:       "ID_CUSIP",
						IdValue:      *retryCusipCopy,
						HoldingsItem: figiPayload[i].HoldingsItem,
					}
					figiPayload[i] = retryFigiPl
					retryFigiPayload = append(retryFigiPayload, retryFigiPl)
				}
				if figiPayload[i].HoldingsItem.SEDOL != nil {
					retrySedolCopy := figiPayload[i].HoldingsItem.SEDOL
					retryFigiPayload = append(retryFigiPayload, FigiPayload{
						IdType:       "ID_SEDOL",
						IdValue:      *retrySedolCopy,
						HoldingsItem: figiPayload[i].HoldingsItem,
					})
				}
				if figiPayload[i].HoldingsItem.Ticker != "" {
					retryTickerCopy := figiPayload[i].HoldingsItem.Ticker
					retryFigiPayload = append(retryFigiPayload, FigiPayload{
						IdType:       "TICKER",
						IdValue:      retryTickerCopy,
						HoldingsItem: figiPayload[i].HoldingsItem,
					})
				}
				retryFigi, err := s.figiClient.GetFigi(retryFigiPayload)
				if err != nil {
					return err
				}
				for _, rfgi := range retryFigi {
					if len(rfgi.Data) > 0 {
						figiR = rfgi
						break
					}
				}
			}
			if len(figiR.Data) == 0 {
				fmt.Println("NO FIGI FOUND")
				continue
			}
			figiCopy := figiR.Data[0].Figi
			shouldAdd := true
			for _, mi := range mappings {
				if mi.Figi == figiCopy {
					shouldAdd = false
				}
			}
			if shouldAdd {
				tickerCopy := figiR.Data[0].Ticker
				nameCopy := figiR.Data[0].Name
				var cusipCopy *string
				var sedolCopy *string
				cusipCopy = figiPayload[i].HoldingsItem.CUSIP
				sedolCopy = figiPayload[i].HoldingsItem.SEDOL
				mappings = append(mappings, model.FigiMapping{
					Figi:   figiCopy,
					Ticker: &tickerCopy,
					Name:   &nameCopy,
					Sedol:  sedolCopy,
					Cusip:  cusipCopy,
				})
			}
		}
		if len(mappings) == 0 {
			fmt.Println("No mappings found")
			return nil
		}
		err = s.repo.UpsertFigiMapping(ctx, mappings, tx)
		if err != nil {
			return err
		}
		for i, upsertedFigi := range mappings {
			if upsertedFigi.Cusip != nil {
				cusipFigiMap[upsertedFigi.Figi] = upsertedFigi
			}
			if upsertedFigi.Sedol != nil {
				sedolFigiMap[upsertedFigi.Figi] = upsertedFigi
			}
			upsertedFigiCopy := upsertedFigi.Figi
			h := model.Holding{Figi: &upsertedFigiCopy}
			setHoldingValues(figiPayload[i].HoldingsItem, &h)
			holdingsMap[upsertedFigiCopy] = h
			addToFundHoldingsMap(upsertedFigi.Figi, figiPayload[i].HoldingsItem, fundHoldingsMap)
		}
	}
	return nil
}

func addToFundHoldingsMap(figi string, hi HoldingsItem, fundHoldingsMap map[string]model.FundHolding) {
	amount := hi.NumberOfShares
	percentageOfTotal := hi.MarketValPercent
	marketValue := hi.MarketValue
	fundHoldingsMap[figi] = model.FundHolding{
		Amount:            &amount,
		PercentageOfTotal: &percentageOfTotal,
		MarketValue:       &marketValue,
	}
}

func setHoldingValues(hi HoldingsItem, h *model.Holding) {
	sector, ok := sectorMap[hi.SectorName]
	if !ok {
		sector = fund.UnknownSector
	}
	holdingType, ok := typeMap[hi.IssueTypename]
	if !ok {
		holdingType = fund.UnknownType
	}

	switch holdingType {
	case fund.BondsType:
		sector = fund.BondsSector
	case fund.NotesType:
		sector = fund.NotesSector
	case fund.CashType:
		sector = fund.CashSector
		hi.Ticker = hi.Name
	}

	sectorString := string(sector)
	issueTypeName := string(holdingType)
	h.Sector = &sectorString
	h.Type = &issueTypeName
}

var typeMap = map[string]fund.HoldingType{
	"Limited Partnership":         fund.UnknownType,
	"Bonds":                       fund.BondsType,
	"Equity":                      fund.Stocks,
	"Cash":                        fund.CashType,
	"Unknown":                     fund.UnknownType,
	"Money Market":                fund.MoneyMarketType,
	"Treasury Bill":               fund.TreasuryType,
	"Government Bond":             fund.BondsType,
	"Closed End Funds":            fund.ClosedEndFundType,
	"Cash Collateral and Margins": fund.CashType,
	"Subscription Rights":         fund.UnknownType,
	"Financial Futures":           fund.FuturesType,
	"Commercial Paper":            fund.BondsType,
	"Warrants":                    fund.BondsType,
	"Municipal Instrument":        fund.BondsType,
	"Asset Backed Security":       fund.UnknownType,
	"Currency":                    fund.CashType,
	"Installment Bonds":           fund.BondsType,
	"Common Stock":                fund.Stocks,
	"Futures":                     fund.FuturesType,
	"Index Future":                fund.FuturesType,
	"Treasury Note":               fund.NotesType,
	"Preferred Stock":             fund.Stocks,
	"Mortgage Backed Security":    fund.UnknownType,
	"Treasury Bond":               fund.BondsType,
	"Other Equity":                fund.UnknownType,
	"Corporate Bond":              fund.BondsType,
	"Mutual Fund":                 fund.MutualFundType,
}

var sectorMap = map[string]fund.SectorName{
	"Infrastructure REITs":                     fund.RealEstateSector,
	"Computer Services":                        fund.TechnologySector,
	"Medical Services":                         fund.HealthCareSector,
	"Funeral Parlors and Cemetery":             fund.HealthCareSector,
	"Cannabis Producers":                       fund.HealthCareSector,
	"Full Line Insurance":                      fund.FinancialsSector,
	"Health Care Facilities":                   fund.HealthCareSector,
	"Entertainment":                            fund.ConsumerDiscretionarySector,
	"Home Improvement Retailers":               fund.ConsumerDiscretionarySector,
	"Delivery Services":                        fund.IndustrialsSector,
	"Forms and Bulk Printing Services":         fund.IndustrialsSector,
	"Business Training and Employment Agencie": fund.IndustrialsSector,
	"Electronic Components":                    fund.TechnologySector,
	"Pipelines":                                fund.EnergySector,
	"Diversified REITs":                        fund.RealEstateSector,
	"Defense":                                  fund.IndustrialsSector,
	"Machinery: Agricultural":                  fund.IndustrialsSector,
	"Tires":                                    fund.ConsumerDiscretionarySector,
	"Conventional Electricity":                 fund.UtilitiesSector,
	"Education Services":                       fund.ConsumerDiscretionarySector,
	"Electronic Equipment: Pollution Control":  fund.TechnologySector,
	"Engineering and Contracting Services":     fund.IndustrialsSector,
	"Restaurants and Bars":                     fund.ConsumerDiscretionarySector,
	"Health Care REITs":                        fund.RealEstateSector,
	"Other Specialty REITs":                    fund.RealEstateSector,
	"Coal":                                     fund.EnergySector,
	"Hotel and Lodging REITs":                  fund.RealEstateSector,
	"Asset Managers and Custodians":            fund.FinancialsSector,
	"Paints and Coatings":                      fund.MaterialsSector,
	"Waste and Disposal Services":              fund.IndustrialsSector,
	"Specialty Retailers":                      fund.ConsumerDiscretionarySector,
	"Commercial Vehicle-Equipment Leasing":     fund.IndustrialsSector,
	"Consumer Digital Services":                fund.TechnologySector,
	"Health Care Management Services":          fund.HealthCareSector,
	"Machinery: Specialty":                     fund.IndustrialsSector,
	"Soft Drinks":                              fund.ConsumerStaplesSector,
	"Brewers":                                  fund.ConsumerStaplesSector,
	"Water":                                    fund.UtilitiesSector,
	"Cosmetics":                                fund.ConsumerStaplesSector,
	"Glass":                                    fund.MaterialsSector,
	"Storage Facilities":                       fund.RealEstateSector,
	"Oil: Crude Producers":                     fund.EnergySector,
	"Banks":                                    fund.FinancialsSector,
	"Containers and Packaging":                 fund.MaterialsSector,
	"Household Furnishings":                    fund.ConsumerDiscretionarySector,
	"Industrial Suppliers":                     fund.IndustrialsSector,
	"Life Insurance":                           fund.FinancialsSector,
	"Aerospace":                                fund.IndustrialsSector,
	"Chemicals and Synthetic Fibers":           fund.MaterialsSector,
	"Real Estate Holding and Development":      fund.RealEstateSector,
	"Apparel Retailers":                        fund.ConsumerDiscretionarySector,
	"Nonferrous Metals":                        fund.MaterialsSector,
	"Footwear":                                 fund.ConsumerDiscretionarySector,
	"Railroad Equipment":                       fund.IndustrialsSector,
	"Fertilizers":                              fund.MaterialsSector,
	"Commercial Vehicles and Parts":            fund.IndustrialsSector,
	"Specialty Chemicals":                      fund.MaterialsSector,
	"Publishing":                               fund.ConsumerDiscretionarySector,
	"Machinery: Industrial":                    fund.IndustrialsSector,
	"Professional Business Support Services":   fund.IndustrialsSector,
	"Cable Television Services":                fund.ConsumerDiscretionarySector,
	"Property and Casualty Insurance":          fund.FinancialsSector,
	"Railroads":                                fund.IndustrialsSector,
	"Integrated Oil and Gas":                   fund.EnergySector,
	"Telecommunications Services":              fund.TelecommunicationSector,
	"Closed End Investments":                   fund.FinancialsSector,
	"Drug Retailers":                           fund.HealthCareSector,
	"Transportation Services":                  fund.IndustrialsSector,
	"Alternative Fuels":                        fund.EnergySector,
	"Clothing and Accessories":                 fund.ConsumerDiscretionarySector,
	"Open End and Miscellaneous Investment Ve": fund.FinancialsSector,
	"Personal Products":                        fund.ConsumerStaplesSector,
	"Security Services":                        fund.IndustrialsSector,
	"Electronic Entertainment":                 fund.TechnologySector,
	"Vending and Catering Service":             fund.ConsumerDiscretionarySector,
	"Household Equipment and Products":         fund.ConsumerDiscretionarySector,
	"Real Estate Services":                     fund.RealEstateSector,
	"Offshore Drilling and Other Services":     fund.EnergySector,
	"Machinery: Construction and Handling":     fund.IndustrialsSector,
	"Distillers and Vintners":                  fund.ConsumerStaplesSector,
	"Electronic Equipment: Control and Filter": fund.TechnologySector,
	"Fruit and Grain Processing":               fund.ConsumerStaplesSector,
	"Software":                                 fund.TechnologySector,
	"Platinum and Precious Metals":             fund.MaterialsSector,
	"Pharmaceuticals":                          fund.HealthCareSector,
	"Consumer Services: Misc.":                 fund.ConsumerDiscretionarySector,
	"Reinsurance":                              fund.FinancialsSector,
	"Electronic Equipment: Gauges and Meters":  fund.TechnologySector,
	"Auto Services":                            fund.ConsumerDiscretionarySector,
	"Luxury Items":                             fund.ConsumerDiscretionarySector,
	"Construction":                             fund.IndustrialsSector,
	"General Mining":                           fund.MaterialsSector,
	"Recreational Vehicles and Boats":          fund.ConsumerDiscretionarySector,
	"Household Appliance":                      fund.ConsumerDiscretionarySector,
	"Residential REITs":                        fund.RealEstateSector,
	"Casinos and Gambling":                     fund.ConsumerDiscretionarySector,
	"Electronic Office Equipment":              fund.TechnologySector,
	"Telecommunications Equipment":             fund.TechnologySector,
	"Electrical Components":                    fund.TechnologySector,
	"Health Care Services":                     fund.HealthCareSector,
	"Hotels and Motels":                        fund.ConsumerDiscretionarySector,
	"Building: Climate Control":                fund.IndustrialsSector,
	"Travel and Tourism":                       fund.ConsumerDiscretionarySector,
	"Timber REITs":                             fund.RealEstateSector,
	"Mortgage REITs: Residential":              fund.RealEstateSector,
	"Computer Hardware":                        fund.TechnologySector,
	"Automobiles":                              fund.ConsumerDiscretionarySector,
	"Iron and Steel":                           fund.MaterialsSector,
	"Marine Transportation":                    fund.IndustrialsSector,
	"Consumer Lending":                         fund.FinancialsSector,
	"Building Roofing/Wallboard and Plumbing":  fund.IndustrialsSector,
	"Mortgage REITs: Commercial":               fund.RealEstateSector,
	"Cement":                                   fund.MaterialsSector,
	"Radio and TV Broadcasters":                fund.ConsumerDiscretionarySector,
	"Machinery: Tools":                         fund.IndustrialsSector,
	"Home Construction":                        fund.ConsumerDiscretionarySector,
	"Gas Distribution":                         fund.UtilitiesSector,
	"Airlines":                                 fund.IndustrialsSector,
	"Toys":                                     fund.ConsumerDiscretionarySector,
	"Transaction Processing Services":          fund.TechnologySector,
	"Industrial REITs":                         fund.RealEstateSector,
	"Photography":                              fund.ConsumerDiscretionarySector,
	"Trucking":                                 fund.IndustrialsSector,
	"Diversified Materials":                    fund.MaterialsSector,
	"Insurance Brokers":                        fund.FinancialsSector,
	"Recreational Products":                    fund.ConsumerDiscretionarySector,
	"Diversified Industrials":                  fund.IndustrialsSector,
	"Diversified Financial Services":           fund.FinancialsSector,
	"Biotechnology":                            fund.HealthCareSector,
	"Mortgage REITs: Diversified":              fund.RealEstateSector,
	"Food Retailers and Wholesalers":           fund.ConsumerStaplesSector,
	"Medical Equipment":                        fund.HealthCareSector,
	"Chemicals: Diversified":                   fund.MaterialsSector,
	"Metal Fabricating":                        fund.MaterialsSector,
	"Multi-Utilities":                          fund.UtilitiesSector,
	"Miscellaneous Consumer Staple Goods":      fund.ConsumerStaplesSector,
	"Printing and Copying Services":            fund.IndustrialsSector,
	"Diversified Retailers":                    fund.ConsumerDiscretionarySector,
	"Plastics":                                 fund.MaterialsSector,
	"Retail REITs":                             fund.RealEstateSector,
	"Copper":                                   fund.MaterialsSector,
	"Auto Parts":                               fund.ConsumerDiscretionarySector,
	"Building Materials: Other":                fund.MaterialsSector,
	"Mortgage Finance":                         fund.FinancialsSector,
	"Tobacco":                                  fund.ConsumerStaplesSector,
	"Paper":                                    fund.MaterialsSector,
	"Media Agencies":                           fund.ConsumerDiscretionarySector,
	"Production Technology Equipment":          fund.IndustrialsSector,
	"Textile Products":                         fund.ConsumerDiscretionarySector,
	"Consumer Electronics":                     fund.TechnologySector,
	"Office REITs":                             fund.RealEstateSector,
	"Investment Services":                      fund.FinancialsSector,
	"Rental and Leasing Services: Consumer":    fund.ConsumerDiscretionarySector,
	"Recreational Services":                    fund.ConsumerDiscretionarySector,
	"Financial Data Providers":                 fund.FinancialsSector,
	"Aluminum":                                 fund.MaterialsSector,
	"Forestry":                                 fund.MaterialsSector,
	"Health Care: Misc.":                       fund.HealthCareSector,
	"Renewable Energy Equipment":               fund.EnergySector,
	"Alternative Electricity":                  fund.EnergySector,
	"Machinery: Engines":                       fund.IndustrialsSector,
	"Sugar":                                    fund.ConsumerStaplesSector,
	"Medical Supplies":                         fund.HealthCareSector,
	"Gold Mining":                              fund.MaterialsSector,
	"Storage REITs":                            fund.RealEstateSector,
	"Farming Fishing Ranching and Plantatio":   fund.ConsumerStaplesSector,
	"Nondurable Household Products":            fund.ConsumerStaplesSector,
	"Electronic Equipment: Other":              fund.TechnologySector,
	"Diamonds and Gemstones":                   fund.MaterialsSector,
	"Oil Equipment and Services":               fund.EnergySector,
	"Semiconductors":                           fund.TechnologySector,
	"Oil Refining and Marketing":               fund.EnergySector,
	"Food Products":                            fund.ConsumerStaplesSector,
}
