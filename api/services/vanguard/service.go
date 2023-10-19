package vanguard

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"etfinsight/generated/jet_gen/postgres/public/model"
	"etfinsight/services/fund"
	"etfinsight/utils/isinutils"
	"etfinsight/utils/stringutils"

	"github.com/google/uuid"
)

type Service struct {
	client EtfIssuerClient
	repo   Repository
}

func NewService(client EtfIssuerClient, repo Repository) *Service {
	return &Service{client: client, repo: repo}
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
	batches = append(batches, ei)
	for _, batch := range batches {
		return func() error {
			defer func() {
				a := recover()

				if a != nil {
					fmt.Println("Recovered", a)
				}
			}()
			return s.upsertFunds(ctx, batch)
		}()
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
		fmt.Println(fmt.Sprintf("Converted %s", f.Profile.FundFullName))
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
	holdings, fundHoldings := convertToHoldings(fund.Holdings.OriginalHoldings.Items)
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
		upsertedHoldings, err := s.repo.UpsertHoldings(ctx, holdings, tx)
		if err != nil {
			return err
		}
		var aggFundHoldings []model.FundHolding
		for fundHoldingTicker, fh := range fundHoldings {
			upsertedHoldingId := upsertedHoldings[fundHoldingTicker]
			fh.FundID = &fundId
			fh.HoldingID = &upsertedHoldingId
			aggFundHoldings = append(aggFundHoldings, fh)
		}
		err = s.repo.UpsertFundHoldings(ctx, aggFundHoldings, tx)
	}
	if err != nil {
		return err
	}
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

func convertToHoldings(HoldingItems []HoldingsItem) ([]model.Holding, map[string]model.FundHolding) {
	var holdings []model.Holding
	holdingMap := map[string]model.Holding{}
	fundHoldingsMap := map[string]model.FundHolding{}
	for _, hi := range HoldingItems {
		var ISIN *string
		if hi.SEDOL != nil && (hi.CountryCode == "GB" || hi.CountryCode == "IE") {
			result, err := isinutils.SEDOLtoISIN(*hi.SEDOL, hi.CountryCode)
			if err == nil {
				ISIN = &result
			}
		} else if hi.CUSIP != nil && (hi.CountryCode == "US" || hi.CountryCode == "CA" || hi.CountryCode == "BM") {
			result, err := isinutils.CUSIPtoISIN(*hi.CUSIP, hi.CountryCode)
			if err == nil {
				ISIN = &result
			}
		}
		if hi.IssueTypename == fund.Currency {
			hi.Ticker = hi.Name
		}
		if hi.Ticker == "" {
			str := ""
			hi.Ticker = "UNKNOWN"
			hi.Name = "Unknown company"
			hi.SEDOL = &str
			hi.CUSIP = &str
			hi.IssueTypename = fund.Unknown
		}

		ticker := hi.Ticker
		name := hi.Name
		SEDOL := hi.SEDOL
		CUSIP := hi.CUSIP
		sector, ok := sectorMap[hi.SectorName]
		if !ok {
			sector = fund.UnknownSector
		}
		sectorString := string(sector)
		issueTypeName := string(hi.IssueTypename)
		holding := model.Holding{
			Ticker: &ticker,
			Type:   &issueTypeName,
			Isin:   ISIN,
			Name:   &name,
			Sedol:  SEDOL,
			Cusip:  CUSIP,
			Sector: &sectorString,
		}
		amount := hi.NumberOfShares
		percentageOfTotal := hi.MarketValPercent
		marketValue := hi.MarketValue
		if _, ok := holdingMap[ticker]; ok {
			matchingFundHolding := fundHoldingsMap[ticker]
			newAmount := amount + *matchingFundHolding.Amount
			newPercentageOfTotal := percentageOfTotal + *matchingFundHolding.PercentageOfTotal
			newMarketValue := marketValue + *matchingFundHolding.MarketValue
			fundHoldingsMap[ticker] = model.FundHolding{
				Amount:            &newAmount,
				PercentageOfTotal: &newPercentageOfTotal,
				MarketValue:       &newMarketValue,
			}
			continue
		}
		fundHoldingsMap[ticker] = model.FundHolding{
			Amount:            &amount,
			PercentageOfTotal: &percentageOfTotal,
			MarketValue:       &marketValue,
		}
		holdingMap[ticker] = holding
		holdings = append(holdings, holding)
	}
	return holdings, fundHoldingsMap
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
