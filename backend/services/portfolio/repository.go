package portfolio

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
	"etfinsight/services/fund"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	NewTransaction(ctx context.Context) (pgx.Tx, error)
	RollBack(tx pgx.Tx, ctx context.Context)

	GetPortfolios(ctx context.Context, userID uuid.UUID) (Models, error)
	GetListItems(ctx context.Context, portfolioID uuid.UUID) (ListItems, error)
	DeleteListItems(ctx context.Context, ids []uuid.UUID, tx pgx.Tx) error
	UpsertPortfolio(ctx context.Context, portfolio model.Portfolio, tx pgx.Tx) error
	UpsertPortfolioListItems(ctx context.Context, listItems []model.PortfolioFund, tx pgx.Tx) error

	GetRatio(ctx context.Context, portfolioId uuid.UUID) (map[uuid.UUID]float64, error)
	GetPortfolioFundSectors(ctx context.Context, portfolioID uuid.UUID) ([]fund.SectorName, error)
	GetPortfolioFundRelativeWeightings(ctx context.Context, portfolioID uuid.UUID) (RelativeSectorWeightings, error)
	GetPortfolioFunds(ctx context.Context, portfolioID uuid.UUID) ([]fund.Information, error)
	GetPortfolioFundHoldings(ctx context.Context, portfolioId uuid.UUID, searchTerm string, selectedSectors []string, limit int64, offset int64) (FundHoldings, error)
}
