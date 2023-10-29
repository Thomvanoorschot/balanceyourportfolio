package portfolio

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	NewTransaction(ctx context.Context) (pgx.Tx, error)
	RollBack(tx pgx.Tx, ctx context.Context)

	GetPortfolios(ctx context.Context, userID uuid.UUID) (Models, error)
	UpsertPortfolio(ctx context.Context, userID uuid.UUID, portfolio Model, tx pgx.Tx) (Model, error)
	UpsertPortfolioListItems(ctx context.Context, portfolioID uuid.UUID, listItems ListItems, tx pgx.Tx) (ListItems, error)
}
