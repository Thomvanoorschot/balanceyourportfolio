package ishares

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	NewTransaction(ctx context.Context) (pgx.Tx, error)
	RollBack(tx pgx.Tx, ctx context.Context)

	UpsertFund(ctx context.Context, f model.Fund, tx pgx.Tx) (uuid.UUID, error)
	UpsertHoldings(ctx context.Context, holdings []model.Holding, tx pgx.Tx) (map[string]uuid.UUID, error)
	UpsertFundHoldings(ctx context.Context, s []model.FundHolding, tx pgx.Tx) error
	UpsertFundListings(ctx context.Context, s []model.FundListing, tx pgx.Tx) error
	UpsertFigiISINMapping(ctx context.Context, s []model.FigiMapping, tx pgx.Tx) error
}
