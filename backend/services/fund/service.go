package fund

import (
	"context"
	"etfinsight/generated/proto"
	"etfinsight/utils/concurrencyutils"
	"etfinsight/utils/stringutils"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetFundsWithTickers(ctx context.Context, searchTerm string) (*proto.SearchFundsResponse, error) {
	funds, err := s.repo.FilterFunds(ctx, FundsFilter{
		SearchTerm: searchTerm,
		Limit:      5,
		Offset:     0,
	})
	if err != nil {
		return nil, err
	}
	return &proto.SearchFundsResponse{Entries: funds.ConvertToResponse()}, nil
}

func (s *Service) FilterFunds(ctx context.Context, filter *proto.FilterFundsRequest) (*proto.FilterFundsResponse, error) {
	funds, err := s.repo.FilterFunds(ctx, FundsFilter{
		SearchTerm: filter.SearchTerm,
		Providers:  filter.Providers,
		Limit:      filter.Limit,
		Offset:     filter.Offset,
	})
	if err != nil {
		return nil, err
	}
	return &proto.FilterFundsResponse{Entries: funds.ConvertToResponse()}, nil
}

func (s *Service) GetDetails(ctx context.Context, fundId uuid.UUID) (*proto.FundDetailsResponse, error) {
	fundSectorCh := concurrencyutils.Async2(func() (SectorNames, error) {
		return s.repo.GetFundSectors(ctx, fundId)
	})
	fundCh := concurrencyutils.Async2(func() (Information, error) {
		return s.repo.GetFund(ctx, fundId)
	})
	sectorWeightingsCh := concurrencyutils.Async2(func() (SectorWeightings, error) {
		return s.repo.GetFundSectorWeightings(ctx, fundId)
	})
	fundHoldingsCh := concurrencyutils.Async2(func() (Holdings, error) {
		return s.repo.FilterHoldings(ctx, HoldingsFilter{
			FundId:          fundId,
			SearchTerm:      "",
			SelectedSectors: nil,
			Limit:           20,
			Offset:          0,
		})
	})
	fundSectorResult := <-fundSectorCh
	fundResult := <-fundCh
	sectorWeightingsResult := <-sectorWeightingsCh
	fundHoldingsResult := <-fundHoldingsCh
	if fundSectorResult.Error != nil {
		return nil, fundSectorResult.Error
	}
	if fundResult.Error != nil {
		return nil, fundResult.Error
	}
	if sectorWeightingsResult.Error != nil {
		return nil, sectorWeightingsResult.Error
	}
	if fundHoldingsResult.Error != nil {
		return nil, fundHoldingsResult.Error
	}

	return &proto.FundDetailsResponse{
		Sectors:          fundSectorResult.Value.ConvertToResponse(),
		Information:      fundResult.Value.ConvertToResponse(),
		SectorWeightings: sectorWeightingsResult.Value.ConvertToResponse(),
		FundHoldings:     fundHoldingsResult.Value.ConvertToResponse(),
	}, nil
}

func (s *Service) FilterHoldings(ctx context.Context, filter *proto.FilterFundHoldingsRequest) (*proto.FilterFundHoldingsResponse, error) {
	fundHoldings, err := s.repo.FilterHoldings(ctx, HoldingsFilter{
		FundId:          stringutils.ConvertToUUID(filter.FundId),
		SearchTerm:      filter.SearchTerm,
		SelectedSectors: filter.SelectedSectors,
		Limit:           filter.Limit,
		Offset:          filter.Offset,
	})
	if err != nil {
		return nil, err
	}
	return &proto.FilterFundHoldingsResponse{Entries: fundHoldings.ConvertToResponse()}, nil
}

func (s *Service) CompareFunds(ctx context.Context, req *proto.CompareFundRequest) (*proto.CompareFundResponse, error) {
	fundOne := stringutils.ConvertToUUID(req.FundOne)
	fundTwo := stringutils.ConvertToUUID(req.FundTwo)
	totalOverlapCh := concurrencyutils.Async2(func() (OverlappingFunds, error) {
		return s.repo.GetTotalOverlap(ctx, fundOne, fundTwo)
	})
	overlappingHoldingsCh := concurrencyutils.Async2(func() (OverlappingHoldings, error) {
		return s.repo.GetOverlappingHoldings(ctx, fundOne, fundTwo)
	})
	fundOneNonOverlappingHoldingsCh := concurrencyutils.Async2(func() (NonOverlappingHoldings, error) {
		return s.repo.GetNonOverlappingHoldings(ctx, fundOne, fundTwo)
	})
	fundTwoNonOverlappingHoldingsCh := concurrencyutils.Async2(func() (NonOverlappingHoldings, error) {
		return s.repo.GetNonOverlappingHoldings(ctx, fundTwo, fundOne)
	})
	sectorWeightingsCh := concurrencyutils.Async2(func() (SectorWeightings, error) {
		return s.repo.GetFundsSectorWeightings(ctx, fundTwo, fundOne)
	})
	sectorWeightingsResult := <-sectorWeightingsCh
	totalOverlapResult := <-totalOverlapCh
	overlappingHoldingsResult := <-overlappingHoldingsCh
	fundOneNonOverlappingHoldingsResult := <-fundOneNonOverlappingHoldingsCh
	fundTwoNonOverlappingHoldingsResult := <-fundTwoNonOverlappingHoldingsCh

	if totalOverlapResult.Error != nil {
		return nil, totalOverlapResult.Error
	}
	if overlappingHoldingsResult.Error != nil {
		return nil, overlappingHoldingsResult.Error
	}
	if fundOneNonOverlappingHoldingsResult.Error != nil {
		return nil, fundOneNonOverlappingHoldingsResult.Error
	}
	if fundTwoNonOverlappingHoldingsResult.Error != nil {
		return nil, fundTwoNonOverlappingHoldingsResult.Error
	}
	if sectorWeightingsResult.Error != nil {
		return nil, sectorWeightingsResult.Error
	}

	return &proto.CompareFundResponse{
		TotalOverlappingPercentage:        totalOverlapResult.Value.TotalOverlappingPercentage,
		OverlappingHoldingsCount:          totalOverlapResult.Value.OverlappingHoldingsCount,
		FundOneHoldingCount:               totalOverlapResult.Value.FundOneHoldingCount,
		FundOneOverlappingCountPercentage: totalOverlapResult.Value.FundOneOverlappingCountPercentage,
		FundTwoHoldingCount:               totalOverlapResult.Value.FundTwoHoldingCount,
		FundTwoOverlappingCountPercentage: totalOverlapResult.Value.FundTwoOverlappingCountPercentage,
		FundOneName:                       totalOverlapResult.Value.FundOneName,
		FundTwoName:                       totalOverlapResult.Value.FundTwoName,
		OverlappingHoldings:               overlappingHoldingsResult.Value.ConvertToResponse(),
		FundOneNonOverlappingHoldings:     fundOneNonOverlappingHoldingsResult.Value.ConvertToResponse(),
		FundTwoNonOverlappingHoldings:     fundTwoNonOverlappingHoldingsResult.Value.ConvertToResponse(),
		SectorWeightings:                  sectorWeightingsResult.Value.ConvertToResponse(),
	}, nil
}
