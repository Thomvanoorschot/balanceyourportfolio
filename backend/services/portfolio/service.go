package portfolio

import (
	"context"

	"etfinsight/api/contracts"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetPortfolios(ctx context.Context, userID uuid.UUID) ([]contracts.Portfolio, error) {
	p, err := s.repo.GetPortfolios(ctx, userID)
	if err != nil {
		return nil, err
	}
	return p.ConvertToResponse(), nil
}
func (s *Service) UpsertPortfolio(ctx context.Context,
	userID uuid.UUID,
	req contracts.Portfolio) (resp contracts.Portfolio, err error) {
	tx, err := s.repo.NewTransaction(ctx)
	if err != nil {
		return resp, err
	}
	defer s.repo.RollBack(tx, ctx)

	p := ConvertToModel(req)
	p, err = s.repo.UpsertPortfolio(ctx, userID, p, tx)
	if err != nil {
		return resp, err
	}
	pli, err := s.repo.UpsertPortfolioListItems(ctx, p.ID, p.Items, tx)
	if err != nil {
		return resp, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return contracts.Portfolio{}, err
	}
	p.Items = pli
	resp = p.ConvertToResponse()
	return resp, nil
}
