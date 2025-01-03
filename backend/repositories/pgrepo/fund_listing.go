package pgrepo

import (
	"context"

	"balanceyourportfolio/generated/jet_gen/postgres/public/model"
	. "balanceyourportfolio/generated/jet_gen/postgres/public/table"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) UpsertFundListings(ctx context.Context, s []model.FundListing, tx pgx.Tx) error {
	sql, args := FundListing.
		INSERT(FundListing.MutableColumns).
		MODELS(s).
		ON_CONFLICT(FundListing.FundID, FundListing.Ticker).
		DO_NOTHING().
		Sql()

	_, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
