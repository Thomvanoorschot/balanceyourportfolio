package ishares

import (
	"context"
	"errors"
	"etfinsight/generated/jet_gen/postgres/public/model"
	"etfinsight/services/fund"
	"fmt"
	"github.com/google/uuid"
)

type Service struct {
	client EtfIssuerClient
	repo   Repository
}

func NewService(client EtfIssuerClient, repo Repository) *Service {
	return &Service{client: client, repo: repo}
}

var (
	iShares = "IShares"
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
	holdings, fundHoldings, err := convertToHoldings(fetchedFund)
	if err != nil {
		return err
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

func convertToHoldings(fundResponse FundResponse) ([]model.Holding, map[string]model.FundHolding, error) {
	var holdings []model.Holding
	holdingMap := map[string]model.Holding{}
	fundHoldingsMap := map[string]model.FundHolding{}

	for _, entry := range fundResponse.Holdings.AaData {
		ticker, ok := entry[fundResponse.HoldingsTableIndex["colTicker"]].(string)
		if !ok {
			return nil, nil, errors.New("could not find ticker")
		}
		name, ok := entry[fundResponse.HoldingsTableIndex["colIssueName"]].(string)
		if !ok {
			return nil, nil, errors.New("could not find name")
		}
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
		isin, ok := entry[fundResponse.HoldingsTableIndex["colIsin"]].(string)
		if !ok {
			return nil, nil, errors.New("could not find isin")
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
		holding := model.Holding{
			Ticker: &ticker,
			Type:   &typeStr,
			Isin:   &isin,
			Sector: &sectorStr,
			Name:   &name,
		}
		if _, ok := holdingMap[ticker]; ok {
			matchingFundHolding := fundHoldingsMap[ticker]
			newPercentageOfTotal := weightRaw + *matchingFundHolding.PercentageOfTotal
			newMarketValue := marketValueRaw + *matchingFundHolding.MarketValue
			fundHoldingsMap[ticker] = model.FundHolding{
				PercentageOfTotal: &newPercentageOfTotal,
				MarketValue:       &newMarketValue,
				Amount:            &amountRaw,
			}
			continue
		}
		fundHoldingsMap[ticker] = model.FundHolding{
			PercentageOfTotal: &weightRaw,
			MarketValue:       &marketValueRaw,
			Amount:            &amountRaw,
		}
		holdingMap[ticker] = holding
		holdings = append(holdings, holding)
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
