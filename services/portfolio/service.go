package portfolio

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePortfolio(ctx context.Context, req CreatePortfolioRequest) error {
	userId := uuid.MustParse("b21b14c9-70bb-4336-a35c-7a69396ffbd8")

	return s.repo.CreatePortfolio(ctx, model.Portfolio{
		UserID: &userId,
		Name:   &req.Name,
	})
}
