package fund

import (
	"context"

	"etfinsight/utils/concurrencyutils"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
func (s *Service) GetFunds(ctx context.Context, searchTerm string) ([]Fund, error) {
	funds, err := s.repo.GetFunds(ctx, searchTerm)
	if err != nil {
		return nil, err
	}
	return funds, nil
}
func (s *Service) GetFundsWithTickers(ctx context.Context, searchTerm string) ([]Fund, error) {
	funds, err := s.repo.GetFundsWithTickers(ctx, searchTerm)
	if err != nil {
		return nil, err
	}
	return funds, nil
}

func (s *Service) GetEffectiveShares(ctx context.Context, fundId uuid.UUID) ([]EffectiveShare, error) {
	//fundHoldings, err := s.repo.GetFundHoldings(ctx, fundId)
	//if err != nil {
	//	return nil, err
	//}
	//var effectiveShares []EffectiveShare
	//for _, fh := range fundHoldings {
	//	totalInvestment := 2065 * fh.FundPrice
	//	//if  hi.MarketValue == 0 {
	//	//	continue
	//	//}
	//	itemTotalInvestment := totalInvestment * (fh.PercentageOfTotal / 100)
	//	itemEffectiveShares := (itemTotalInvestment / fh.MarketValue) * fh.Amount
	//	effectiveShares = append(effectiveShares, EffectiveShare{
	//		Ticker: fh.Ticker,
	//		Name:   fh.Name,
	//		Amount: fmt.Sprintf("%.2f", itemEffectiveShares),
	//	})
	//}
	return nil, nil
}
func (s *Service) GetFundDetails(ctx context.Context, fundId uuid.UUID, limit int64) (Details, error) {
	holdingsCh := concurrencyutils.Async2(func() ([]Holding, error) {
		return s.repo.GetFundHoldings(ctx, fundId, limit)
	})
	fundSectorCh := concurrencyutils.Async2(func() ([]SectorName, error) {
		return s.repo.GetFundSectors(ctx, fundId)
	})
	fundCh := concurrencyutils.Async2(func() (Information, error) {
		return s.repo.GetFund(ctx, fundId)
	})
	sectorWeightingsCh := concurrencyutils.Async2(func() ([]SectorWeighting, error) {
		return s.repo.GetFundSectorWeightings(ctx, fundId)
	})
	holdingsResult := <-holdingsCh
	fundSectorResult := <-fundSectorCh
	fundResult := <-fundCh
	sectorWeightingsResult := <-sectorWeightingsCh
	if holdingsResult.Error != nil {
		return Details{}, holdingsResult.Error
	}
	if fundSectorResult.Error != nil {
		return Details{}, fundSectorResult.Error
	}
	if fundResult.Error != nil {
		return Details{}, fundResult.Error
	}
	if sectorWeightingsResult.Error != nil {
		return Details{}, sectorWeightingsResult.Error
	}
	fundSectorResult.Value = append([]SectorName{AnySector}, fundSectorResult.Value...)

	var fh []Holding
	for _, item := range holdingsResult.Value {
		fh = append(fh, Holding{
			Ticker:            item.Ticker,
			Name:              item.Name,
			Type:              item.Type,
			PercentageOfTotal: item.PercentageOfTotal,
		})
	}

	return Details{
		Holdings:         fh,
		Sectors:          fundSectorResult.Value,
		Information:      fundResult.Value,
		SectorWeightings: sectorWeightingsResult.Value,
	}, nil
}
func (s *Service) FilterHoldings(ctx context.Context, filter HoldingsFilter) ([]Holding, error) {
	fundHoldings, err := s.repo.FilterHoldings(ctx, filter)
	if err != nil {
		return nil, err
	}
	return fundHoldings, nil
}
