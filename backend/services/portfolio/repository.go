package portfolio

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
)

type Repository interface {
	CreatePortfolio(ctx context.Context, portfolio model.Portfolio) error
	UpsertPortfolioFunds(ctx context.Context, pf []model.PortfolioFund) error
}
