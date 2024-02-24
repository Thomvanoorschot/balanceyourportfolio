package ishares

import (
	"balanceyourportfolio/generated/jet_gen/postgres/public/model"
	"balanceyourportfolio/services/fund"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	client     EtfIssuerClient
	repo       Repository
	figiClient FigiClient
}

func NewService(client EtfIssuerClient, repo Repository, figiClient FigiClient) *Service {
	return &Service{client: client, repo: repo, figiClient: figiClient}
}

var (
	iShares     = "IShares"
	isinFigiMap = map[string]model.FigiMapping{}
)

func (s *Service) UpsertFunds(ctx context.Context) error {
	for i := 0; i < 1500; i += 20 {
		f, err := s.client.GetFunds(20, i)
		if err != nil {
			return err
		}
		for i := range f {
			err := s.convertFund(ctx, f[i])
			if err != nil {
				fmt.Printf("could not convert fund %s %s\n", f[i].FundName, err.Error())
			}
			fmt.Printf("converted fund %s\n", f[i].FundName)
		}
	}
	return nil
}
func (s *Service) convertFund(ctx context.Context, fetchedFund FundResponse) error {
	tx, err := s.repo.NewTransaction(ctx)
	if err != nil {
		return err
	}
	defer s.repo.RollBack(tx, ctx)
	totalHoldings := float64(len(fetchedFund.Holdings.AaData))
	if fetchedFund.NetAssets == 0 || fetchedFund.Price == 0 {
		return nil
	}
	outstandingShares := fetchedFund.NetAssets / fetchedFund.Price
	f := model.Fund{
		Name:               &fetchedFund.FundName,
		Currency:           &fetchedFund.Currency,
		Isin:               &fetchedFund.ISIN,
		TotalHoldings:      &totalHoldings,
		Price:              &fetchedFund.Price,
		Provider:           &iShares,
		ExternalIdentifier: &fetchedFund.ExternalIdentifier,
		OutstandingShares:  &outstandingShares,
		EffectiveDate:      &fetchedFund.EffectiveDate,
	}
	fundId, err := s.repo.UpsertFund(ctx, f, tx)
	if err != nil {
		return err
	}
	fundListings := convertToListings(fundId, fetchedFund.Tickers)
	if len(fundListings) > 0 {
		err = s.repo.UpsertFundListings(ctx, fundListings, tx)
		if err != nil {
			return err
		}
	}
	holdings, fundHoldings, err := s.convertToHoldings(fetchedFund, tx)
	if err != nil {
		return err
	}
	if len(holdings) > 0 {
		upsertedHoldings, err := s.repo.UpsertHoldings(ctx, holdings, tx)
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
		err = s.repo.UpsertFundHoldings(ctx, aggFundHoldings, tx)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func convertToListings(fundId uuid.UUID, listings []string) []model.FundListing {
	var fundListings []model.FundListing
	for _, l := range listings {
		shadowedFundId := fundId
		shadowedListing := l
		fundListings = append(fundListings, model.FundListing{
			FundID: &shadowedFundId,
			Ticker: &shadowedListing,
		})
	}
	return fundListings
}

var tempUnknownSectorMap = map[string]interface{}{}

func (s *Service) convertToHoldings(fundResponse FundResponse, tx pgx.Tx) ([]model.Holding, map[string]model.FundHolding, error) {
	var holdings []model.Holding
	holdingMap := map[string]model.Holding{}
	fundHoldingsMap := map[string]model.FundHolding{}

	batchSize := 100
	batches := make([][][]interface{}, 0, (len(fundResponse.Holdings.AaData)+batchSize-1)/batchSize)

	for batchSize < len(fundResponse.Holdings.AaData) {
		fundResponse.Holdings.AaData, batches = fundResponse.Holdings.AaData[batchSize:], append(batches, fundResponse.Holdings.AaData[0:batchSize:batchSize])
	}

	batches = append(batches, fundResponse.Holdings.AaData)
	for _, batch := range batches {
		var figiPayload []FigiPayload

		for _, entry := range batch {
			isin, ok := entry[fundResponse.HoldingsTableIndex["colIsin"]].(string)
			if !ok {
				return nil, nil, errors.New("could not find isin")
			}
			figi, ok := isinFigiMap[isin]
			if ok {
				iSector, ok := entry[fundResponse.HoldingsTableIndex["colSectorName"]].(string)
				if !ok {
					return nil, nil, errors.New("could not find sector name")
				}
				hType, ok := entry[fundResponse.HoldingsTableIndex["colAssetClass"]].(string)
				if !ok {
					return nil, nil, errors.New("could not find holding type")
				}
				marketValue, ok := entry[fundResponse.HoldingsTableIndex["colMarketValue"]].(map[string]interface{})
				if !ok {
					return nil, nil, errors.New("could not find marketValue")
				}
				weight, ok := entry[fundResponse.HoldingsTableIndex["colHoldingPercent"]].(map[string]interface{})
				if !ok {
					return nil, nil, errors.New("could not find weight")
				}
				amount, ok := entry[fundResponse.HoldingsTableIndex["colUnitsHeld"]].(map[string]interface{})
				if !ok {
					amount, ok = entry[fundResponse.HoldingsTableIndex["colParValue"]].(map[string]interface{})
					if !ok {
						return nil, nil, errors.New("could not find amount")
					}
				}
				sector, ok := sectorMap[iSector]
				weightRaw := weight["raw"].(float64)
				marketValueRaw := marketValue["raw"].(float64)
				amountRaw := amount["raw"].(float64)
				if !ok {
					tempUnknownSectorMap[iSector] = nil
					sector = fund.UnknownSector
				}
				holdingType, ok := typeMap[hType]
				if !ok {
					holdingType = fund.UnknownType
				}
				sectorStr := string(sector)
				typeStr := string(holdingType)
				figiCopy := figi.Figi
				holding := model.Holding{
					Figi:   &figiCopy,
					Type:   &typeStr,
					Sector: &sectorStr,
				}
				if _, ok := holdingMap[figiCopy]; ok {
					matchingFundHolding := fundHoldingsMap[figiCopy]
					newPercentageOfTotal := weightRaw + *matchingFundHolding.PercentageOfTotal
					newMarketValue := marketValueRaw + *matchingFundHolding.MarketValue
					fundHoldingsMap[figiCopy] = model.FundHolding{
						PercentageOfTotal: &newPercentageOfTotal,
						MarketValue:       &newMarketValue,
						Amount:            &amountRaw,
					}
					continue
				}
				fundHoldingsMap[figiCopy] = model.FundHolding{
					PercentageOfTotal: &weightRaw,
					MarketValue:       &marketValueRaw,
					Amount:            &amountRaw,
				}
				holdingMap[figiCopy] = holding
				holdings = append(holdings, holding)
			} else if isin != "-" {

				figiPayload = append(figiPayload, FigiPayload{
					IdType:  "ID_ISIN",
					IdValue: isin,
				})
			}
		}
		if len(figiPayload) == 0 {
			continue
		}
		// TODO Deal with currency
		figisResp, err := s.figiClient.GetFigi(figiPayload)
		if err != nil {
			fmt.Println(err)
			return nil, nil, nil
		}
		var mappings []model.FigiMapping
		for _, entry := range batch {
			isin, ok := entry[fundResponse.HoldingsTableIndex["colIsin"]].(string)
			if !ok {
				return nil, nil, errors.New("could not find isin")
			}
			var figiMapping model.FigiMapping
			for i, figiR := range figisResp {
				if figiR.Warning != "" || len(figiR.Data) == 0 {
					fmt.Println("NO FIGI FOUND")
					continue
				} else if figiPayload[i].IdValue == isin {
					tickerCopy := figiR.Data[0].Ticker
					nameCopy := figiR.Data[0].Name
					isinCopy := isin
					figiCopy := figiR.Data[0].Figi
					shareClassFigi := figiR.Data[0].ShareClassFigi
					if shareClassFigi != nil {
						figiCopy = *shareClassFigi
					}
					figiMapping = model.FigiMapping{
						Figi:   figiCopy,
						Ticker: &tickerCopy,
						Name:   &nameCopy,
						Isin:   &isinCopy,
					}
					mappings = append(mappings, figiMapping)
					isinFigiMap[isin] = figiMapping
					iSector, ok := entry[fundResponse.HoldingsTableIndex["colSectorName"]].(string)
					if !ok {
						return nil, nil, errors.New("could not find sector name")
					}
					hType, ok := entry[fundResponse.HoldingsTableIndex["colAssetClass"]].(string)
					if !ok {
						return nil, nil, errors.New("could not find holding type")
					}
					marketValue, ok := entry[fundResponse.HoldingsTableIndex["colMarketValue"]].(map[string]interface{})
					if !ok {
						return nil, nil, errors.New("could not find marketValue")
					}
					weight, ok := entry[fundResponse.HoldingsTableIndex["colHoldingPercent"]].(map[string]interface{})
					if !ok {
						return nil, nil, errors.New("could not find weight")
					}
					amount, ok := entry[fundResponse.HoldingsTableIndex["colUnitsHeld"]].(map[string]interface{})
					if !ok {
						amount, ok = entry[fundResponse.HoldingsTableIndex["colParValue"]].(map[string]interface{})
						if !ok {
							return nil, nil, errors.New("could not find amount")
						}
					}
					sector, ok := sectorMap[iSector]
					weightRaw := weight["raw"].(float64)
					marketValueRaw := marketValue["raw"].(float64)
					amountRaw := amount["raw"].(float64)
					if !ok {
						tempUnknownSectorMap[iSector] = nil
						sector = fund.UnknownSector
					}
					holdingType, ok := typeMap[hType]
					if !ok {
						holdingType = fund.UnknownType
					}
					sectorStr := string(sector)
					typeStr := string(holdingType)
					figiMappingCopy := figiMapping.Figi
					holding := model.Holding{
						Figi:   &figiMappingCopy,
						Type:   &typeStr,
						Sector: &sectorStr,
					}
					if _, ok := holdingMap[figiMappingCopy]; ok {
						matchingFundHolding := fundHoldingsMap[figiMappingCopy]
						newPercentageOfTotal := weightRaw + *matchingFundHolding.PercentageOfTotal
						newMarketValue := marketValueRaw + *matchingFundHolding.MarketValue
						fundHoldingsMap[figiMappingCopy] = model.FundHolding{
							PercentageOfTotal: &newPercentageOfTotal,
							MarketValue:       &newMarketValue,
							Amount:            &amountRaw,
						}
						continue
					}
					fundHoldingsMap[figiMappingCopy] = model.FundHolding{
						PercentageOfTotal: &weightRaw,
						MarketValue:       &marketValueRaw,
						Amount:            &amountRaw,
					}
					holdingMap[figiMappingCopy] = holding
					holdings = append(holdings, holding)
				}
			}
		}
		if len(mappings) > 0 {
			err = s.repo.UpsertFigiISINMapping(context.Background(), mappings, tx)
			if err != nil {
				return nil, nil, err
			}
		}
	}
	return holdings, fundHoldingsMap, nil
}

var typeMap = map[string]fund.HoldingType{
	"Equity":                      fund.Stocks,
	"Cash":                        fund.CashType,
	"FX":                          fund.CashType,
	"Money Market":                fund.MoneyMarketType,
	"Cash Collateral and Margins": fund.CashType,
	"Futures":                     fund.FuturesType,
	"Treasuries":                  fund.TreasuryType,
	"Fixed Income":                fund.FixedIncomeType,
}

var sectorMap = map[string]fund.SectorName{
	"Financials":                         fund.FinancialsSector,
	"Health Care":                        fund.HealthCareSector,
	"Consumer Staples":                   fund.ConsumerStaplesSector,
	"Utilities":                          fund.UtilitiesSector,
	"Real Estate":                        fund.RealEstateSector,
	"Cash and/or Derivatives":            fund.CashSector,
	"Information Technology":             fund.TechnologySector,
	"Technology":                         fund.TechnologySector,
	"Electric":                           fund.TechnologySector,
	"Consumer Discretionary":             fund.ConsumerDiscretionarySector,
	"Communication":                      fund.TelecommunicationSector,
	"Communications":                     fund.TelecommunicationSector,
	"Energy":                             fund.EnergySector,
	"Materials":                          fund.MaterialsSector,
	"Industrials":                        fund.IndustrialsSector,
	"Capital Goods":                      fund.IndustrialsSector,
	"Treasuries":                         fund.BondsSector,
	"Treasury":                           fund.BondsSector,
	"Consumer Non-Cyclical":              fund.ConsumerDiscretionarySector,
	"Consumer Cyclical":                  fund.ConsumerCyclicalSector,
	"Banking":                            fund.FinancialsSector,
	"Finance Companies":                  fund.FinancialsSector,
	"Insurance":                          fund.InsuranceSector,
	"Transportation":                     fund.IndustrialsSector,
	"MBS Pass-Through":                   fund.MortgageBackedSecuritySector,
	"Owned No Guarantee":                 fund.BondsSector,
	"Basic Industry":                     fund.IndustrialsSector,
	"Natural Gas":                        fund.EnergySector,
	"Financial Institutions":             fund.FinancialsSector,
	"Precious Metals":                    fund.MaterialsSector,
	"Securitized":                        fund.BondsSector,
	"Agriculture":                        fund.ConsumerStaplesSector,
	"Industrial Metals":                  fund.MaterialsSector,
	"Other":                              fund.UnknownSector,
	"Financial Other":                    fund.FinancialsSector,
	"Industrial":                         fund.IndustrialsSector,
	"Utility":                            fund.UtilitiesSector,
	"Covered":                            fund.BondsSector,
	"Brokerage/Asset Managers/Exchanges": fund.FinancialsSector,
	"ABS":                                fund.BondsSector,
	"Reits":                              fund.RealEstateSector,
	"CMBS":                               fund.MortgageBackedSecuritySector,
	"Whole Business":                     fund.IndustrialsSector,
	"Industrial Other":                   fund.IndustrialsSector,
	"Utility Other":                      fund.UtilitiesSector,
	"Corporates":                         fund.IndustrialsSector,
	"Livestock":                          fund.ConsumerStaplesSector,
	"Government Sponsored":               fund.BondsSector,
	"Government Guaranteed":              fund.BondsSector,
	"Agency":                             fund.BondsSector,
	"Supranational":                      fund.BondsSector,
	"Local Authority":                    fund.BondsSector,
	"Government Related":                 fund.BondsSector,
	"Sovereign":                          fund.BondsSector,
	"Covered Other":                      fund.BondsSector,
	"Public Sector Collateralized":       fund.BondsSector,
	"Agency Fixed Rate":                  fund.BondsSector,
	"Non-Agency CMBS":                    fund.MortgageBackedSecuritySector,
	"Agency CMBS":                        fund.RealEstateSector,
	"Stranded Cost Utility":              fund.UtilitiesSector,
	"Mortgage Collateralized":            fund.MortgageBackedSecuritySector,
	"Hybrid Collateralized":              fund.MortgageBackedSecuritySector,
}

//UnknownSector               SectorName = "Unknown"
//TechnologySector            SectorName = "Technology"
//HealthCareSector            SectorName = "HealthCare"
//FinancialsSector            SectorName = "Financials"
//RealEstateSector            SectorName = "RealEstate"
//EnergySector                SectorName = "Energy"
//MaterialsSector             SectorName = "Materials"
//ConsumerDiscretionarySector SectorName = "Consumer Discretionary"
//IndustrialsSector           SectorName = "Industrials"
//UtilitiesSector             SectorName = "Utilities"
//ConsumerStaplesSector       SectorName = "Consumer Staples"
//TelecommunicationSector     SectorName = "Telecommunication"
//BondsSector                 SectorName = "Bonds"
//NotesSector                 SectorName = "Notes"
//CashSector                  SectorName = "Cash"
