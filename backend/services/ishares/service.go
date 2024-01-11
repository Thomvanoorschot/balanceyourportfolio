package ishares

import (
	"bytes"
	"context"
	"encoding/json"
	"etfinsight/generated/jet_gen/postgres/public/model"
	"etfinsight/services/fund"
)

type Service struct {
	client EtfIssuerClient
	repo   Repository
}

func NewService(client EtfIssuerClient, repo Repository) *Service {
	return &Service{client: client, repo: repo}
}

func (s *Service) UpsertFunds(ctx context.Context) error {
	fundingBytes, err := s.client.GetFunds()
	if err != nil {
		return err
	}
	fundingBytes = bytes.TrimPrefix(fundingBytes, []byte("\xef\xbb\xbf"))

	holdingsResponse := HoldingsResponse{}
	err = json.Unmarshal(fundingBytes, &holdingsResponse)
	if err != nil {
		return err
	}
	tx, err := s.repo.NewTransaction(ctx)
	if err != nil {
		return err
	}
	defer s.repo.RollBack(tx, ctx)

	var holdings []model.Holding
	for _, entry := range holdingsResponse.AaData {
		h := convertToHoldings(entry)
		holdings = append(holdings, h)
	}
	//_, err = s.repo.UpsertHoldings(ctx, holdings, tx)
	//if err != nil {
	//	return err
	//}
	return nil
	//return tx.Commit(ctx)
}

var test = map[string]struct{}{}

func convertToHoldings(entry []interface{}) model.Holding {

	ticker := entry[0].(string)
	iSector := entry[2].(string)
	hType := entry[3].(string)
	//weight := entry[5].(NumberValue)
	isin := entry[8].(string)
	sector, ok := sectorMap[iSector]
	if !ok {
		sector = fund.UnknownSector
	}
	sectorStr := string(sector)
	test[hType] = struct{}{}
	return model.Holding{
		Ticker: &ticker,
		Type:   &hType,
		Isin:   &isin,
		Sector: &sectorStr,
	}
}

// TODO Fix Types
// Equity
// Cash
// Money Market
// Cash Collateral and Margins
// Futures
var typeMap = map[string]fund.IssueTypeName{}

var sectorMap = map[string]fund.SectorName{
	"Financials":              fund.FinancialsSector,
	"Health Care":             fund.HealthCareSector,
	"Consumer Staples":        fund.ConsumerStaplesSector,
	"Utilities":               fund.UtilitiesSector,
	"Real Estate":             fund.RealEstateSector,
	"Cash and/or Derivatives": fund.CashSector,
	"Information Technology":  fund.TechnologySector,
	"Consumer Discretionary":  fund.ConsumerDiscretionarySector,
	"Communication":           fund.TelecommunicationSector,
	"Energy":                  fund.EnergySector,
	"Materials":               fund.MaterialsSector,
	"Industrials":             fund.IndustrialsSector,
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
