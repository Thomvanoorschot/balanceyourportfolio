package portfolio

import (
	"context"
	"errors"
	"time"

	"etfinsight/generated/jet_gen/postgres/public/model"
	"etfinsight/generated/proto"
	"etfinsight/services/fund"
	"etfinsight/utils/concurrencyutils"
	"etfinsight/utils/stringutils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetPortfolioDetails(ctx context.Context, req *proto.PortfolioDetailsRequest) (*proto.PortfolioDetailsResponse, error) {
	portfolioId := stringutils.ConvertToUUID(req.PortfolioId)
	ownerCh := concurrencyutils.Async2(func() (string, error) {
		return s.repo.GetPortfolioOwner(ctx, portfolioId)
	})
	ratioCh := concurrencyutils.Async2(func() (map[uuid.UUID]float64, error) {
		return s.repo.GetRatio(ctx, portfolioId)
	})
	portfolioSectorCh := concurrencyutils.Async2(func() (fund.SectorNames, error) {
		return s.repo.GetPortfolioFundSectors(ctx, portfolioId)
	})
	informationCh := concurrencyutils.Async2(func() (fund.InformationList, error) {
		return s.repo.GetPortfolioFunds(ctx, portfolioId)
	})
	relativeWeightingsCh := concurrencyutils.Async2(func() (RelativeSectorWeightings, error) {
		return s.repo.GetPortfolioFundRelativeWeightings(ctx, portfolioId)
	})
	portfolioFundHoldingsCh := concurrencyutils.Async2(func() (FundHoldings, error) {
		return s.repo.GetPortfolioFundHoldings(ctx, portfolioId, "", []string{}, 20, 0)
	})
	ratioResult := <-ratioCh
	portfolioSectorResult := <-portfolioSectorCh
	informationResult := <-informationCh
	relativeWeightingsResult := <-relativeWeightingsCh
	portfolioFundHoldingsResult := <-portfolioFundHoldingsCh
	ownerResult := <-ownerCh
	if ratioResult.Error != nil {
		return nil, ratioResult.Error
	}
	if portfolioSectorResult.Error != nil {
		return nil, portfolioSectorResult.Error
	}
	if informationResult.Error != nil {
		return nil, informationResult.Error
	}
	if relativeWeightingsResult.Error != nil {
		return nil, relativeWeightingsResult.Error
	}
	if portfolioFundHoldingsResult.Error != nil {
		return nil, portfolioFundHoldingsResult.Error
	}
	if ownerResult.Error != nil {
		return nil, ownerResult.Error
	}
	if ownerResult.Value != req.UserId {
		// TODO Return GRPC unauthenticated
		return nil, errors.New("unauthorized")
	}
	return &proto.PortfolioDetailsResponse{
		FundInformation:               informationResult.Value.ConvertToResponse(),
		Sectors:                       portfolioSectorResult.Value.ConvertToResponse(),
		PortfolioFundSectorWeightings: relativeWeightingsResult.Value.ConvertToResponse(ratioResult.Value),
		PortfolioFundHoldings:         portfolioFundHoldingsResult.Value.ConvertToResponse(),
	}, nil
}
func (s *Service) GetPortfolios(ctx context.Context, req *proto.PortfoliosRequest) (*proto.PortfoliosResponse, error) {
	p, err := s.repo.GetPortfolios(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return p.ConvertToResponse(), nil
}
func (s *Service) UpsertPortfolio(
	ctx context.Context,
	req *proto.UpsertPortfolioRequest,
) (resp *proto.UpsertPortfolioResponse, err error) {
	tx, err := s.repo.NewTransaction(ctx)
	if err != nil {
		return resp, err
	}
	defer s.repo.RollBack(tx, ctx)

	p := ConvertToModel(req.Portfolio)
	portfolioModel := model.Portfolio{
		ID:     p.Id,
		UserID: &req.UserId,
		Name:   &p.Name,
	}
	if p.Id == uuid.Nil {
		p.Id = uuid.New()
		now := time.Now()
		portfolioModel.CreatedAt = &now
		portfolioModel.ID = p.Id
	} else {
		err = s.checkAndDeleteFunds(ctx, p, tx)
		if err != nil {
			return nil, err
		}
	}
	err = s.repo.UpsertPortfolio(ctx, portfolioModel, tx)
	if err != nil {
		return resp, err
	}
	pfs := p.Items.ConvertToDbModel(p.Id)
	err = s.repo.UpsertPortfolioListItems(ctx, pfs, tx)
	if err != nil {
		return resp, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return &proto.UpsertPortfolioResponse{Portfolio: p.ConvertToResponse()}, nil
}

func (s *Service) checkAndDeleteFunds(ctx context.Context, model Model, tx pgx.Tx) error {
	li, err := s.repo.GetListItems(ctx, model.Id)
	if err != nil {
		return err
	}
	var itemsToDelete []uuid.UUID
	comparisonLoop := func(dbItem ListItem) bool {
		for _, newItem := range model.Items {
			if dbItem.Id == newItem.Id {
				return true
			}
		}
		return false
	}
	for _, dbItem := range li {
		match := comparisonLoop(dbItem)
		if !match {
			itemsToDelete = append(itemsToDelete, dbItem.Id)
		}
	}
	if len(itemsToDelete) > 0 {
		err = s.repo.DeleteListItems(ctx, itemsToDelete, tx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) FilterPortfolioHoldings(ctx context.Context, filter *proto.FilterPortfolioFundHoldingsRequest) (*proto.FilterPortfolioFundHoldingsResponse, error) {
	fundHoldings, err := s.repo.GetPortfolioFundHoldings(ctx,
		uuid.MustParse(filter.PortfolioId),
		filter.SearchTerm,
		filter.SelectedSectors,
		filter.Limit,
		filter.Offset,
	)
	if err != nil {
		return nil, err
	}
	return &proto.FilterPortfolioFundHoldingsResponse{Entries: fundHoldings.ConvertToResponse()}, nil
}

func (s *Service) UpdatePortfolioFundAmount(ctx context.Context, req *proto.UpdatePortfolioFundAmountRequest) (*proto.Empty, error) {
	err := s.repo.UpdatePortfolioFundAmount(ctx,
		stringutils.ConvertToUUID(req.PortfolioId),
		stringutils.ConvertToUUID(req.FundId),
		req.Amount,
	)
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}
