package fund

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	NewTransaction(ctx context.Context) (pgx.Tx, error)
	RollBack(tx pgx.Tx, ctx context.Context)

	GetFundHoldings(ctx context.Context, fundID uuid.UUID, limit int64) (Holdings, error)
	GetFundSectors(ctx context.Context, fundID uuid.UUID) ([]SectorName, error)
	GetFundSectorWeightings(ctx context.Context, fundID uuid.UUID) ([]SectorWeighting, error)
	FilterHoldings(ctx context.Context, filter HoldingsFilter) (Holdings, error)
	GetFund(ctx context.Context, fundID uuid.UUID) (Information, error)
	FilterFunds(ctx context.Context, filter FundsFilter) (Funds, error)
}
