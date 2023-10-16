package fund

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	NewTransaction(ctx context.Context) (pgx.Tx, error)
	RollBack(tx pgx.Tx, ctx context.Context)

	GetFundHoldings(ctx context.Context, fundId uuid.UUID, limit int64) ([]Holding, error)
	GetFundSectors(ctx context.Context, fundId uuid.UUID) ([]SectorName, error)
	GetFundSectorWeightings(ctx context.Context, fundId uuid.UUID) ([]SectorWeighting, error)
	FilterHoldings(ctx context.Context, filter HoldingsFilter) ([]Holding, error)
	GetFunds(ctx context.Context, searchTerm string) ([]Fund, error)
	GetFundsWithTickers(ctx context.Context, searchTerm string) ([]Fund, error)
	GetFund(ctx context.Context, fundId uuid.UUID) (Information, error)
}
