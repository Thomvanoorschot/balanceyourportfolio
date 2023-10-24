package fund

import (
	"context"

	"etfinsight/api/contracts"
	"etfinsight/utils/concurrencyutils"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}
func (s *Service) GetFunds(ctx context.Context, searchTerm string) ([]contracts.Fund, error) {
	funds, err := s.repo.GetFunds(ctx, searchTerm)
	if err != nil {
		return nil, err
	}
	return funds.ConvertToResponse(), nil
}
func (s *Service) GetFundsWithTickers(ctx context.Context, searchTerm string) ([]contracts.Fund, error) {
	funds, err := s.repo.GetFundsWithTickers(ctx, searchTerm)
	if err != nil {
		return nil, err
	}
	return funds.ConvertToResponse(), nil
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
func (s *Service) GetDetails(ctx context.Context, fundID uuid.UUID) (contracts.FundDetails, error) {
	fundSectorCh := concurrencyutils.Async2(func() (SectorNames, error) {
		return s.repo.GetFundSectors(ctx, fundID)
	})
	fundCh := concurrencyutils.Async2(func() (Information, error) {
		return s.repo.GetFund(ctx, fundID)
	})
	sectorWeightingsCh := concurrencyutils.Async2(func() (SectorWeightings, error) {
		return s.repo.GetFundSectorWeightings(ctx, fundID)
	})
	fundSectorResult := <-fundSectorCh
	fundResult := <-fundCh
	sectorWeightingsResult := <-sectorWeightingsCh
	if fundSectorResult.Error != nil {
		return contracts.FundDetails{}, fundSectorResult.Error
	}
	if fundResult.Error != nil {
		return contracts.FundDetails{}, fundResult.Error
	}
	if sectorWeightingsResult.Error != nil {
		return contracts.FundDetails{}, sectorWeightingsResult.Error
	}
	fundSectorResult.Value = append([]SectorName{AnySector}, fundSectorResult.Value...)

	return contracts.FundDetails{
		Sectors:          fundSectorResult.Value.ConvertToResponse(),
		Information:      fundResult.Value.ConvertToResponse(),
		SectorWeightings: sectorWeightingsResult.Value.ConvertToResponse(),
	}, nil
}

func (s *Service) FilterHoldings(ctx context.Context, filter contracts.FundHoldingsFilter) ([]contracts.FundHolding, error) {
	if filter.SectorName == string(AnySector) {
		filter.SectorName = ""
	}
	fundHoldings, err := s.repo.FilterHoldings(ctx, ConvertToHoldingsFilter(filter))
	if err != nil {
		return nil, err
	}
	return fundHoldings.ConvertToResponse(), nil
}
