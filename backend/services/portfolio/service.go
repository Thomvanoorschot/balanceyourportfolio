package portfolio

import (
	"context"
	"fmt"
	"time"

	"etfinsight/generated/jet_gen/postgres/public/model"
	"etfinsight/generated/proto"
	"etfinsight/services/fund"
	"etfinsight/utils/concurrencyutils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetPortfolioDetails(ctx context.Context, userId uuid.UUID, portfolioId uuid.UUID) (*proto.PortfolioDetailsResponse, error) {
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
	portfolioFundHoldingsCh := concurrencyutils.Async2(func() (fund.InformationList, error) {
		return s.repo.GetPortfolioFundHoldings(ctx, portfolioId)
	})
	ratioResult := <-ratioCh
	portfolioSectorResult := <-portfolioSectorCh
	informationResult := <-informationCh
	relativeWeightingsResult := <-relativeWeightingsCh
	test := <-portfolioFundHoldingsCh
	fmt.Println(test)

	return &proto.PortfolioDetailsResponse{
		FundInformation:               informationResult.Value.ConvertToResponse(),
		Sectors:                       portfolioSectorResult.Value.ConvertToResponse(),
		PortfolioFundSectorWeightings: relativeWeightingsResult.Value.ConvertToResponse(ratioResult.Value),
	}, nil
}
func (s *Service) GetPortfolios(ctx context.Context, userId uuid.UUID) (*proto.PortfoliosResponse, error) {
	p, err := s.repo.GetPortfolios(ctx, userId)
	if err != nil {
		return nil, err
	}
	return p.ConvertToResponse(), nil
}
func (s *Service) UpsertPortfolio(ctx context.Context,
	userID uuid.UUID,
	req *proto.Portfolio) (resp *proto.UpsertPortfolioResponse, err error) {
	tx, err := s.repo.NewTransaction(ctx)
	if err != nil {
		return resp, err
	}
	defer s.repo.RollBack(tx, ctx)

	p := ConvertToModel(req)
	portfolioModel := model.Portfolio{
		ID:     p.Id,
		UserID: &userID,
		Name:   &p.Name,
	}
	if p.Id == uuid.Nil {
		p.Id = uuid.New()
		now := time.Now()
		portfolioModel.CreatedAt = &now
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
