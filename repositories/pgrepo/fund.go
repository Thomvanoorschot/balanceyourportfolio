package pgrepo

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
	"etfinsight/services/fund"

	"github.com/georgysavva/scany/v2/pgxscan"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetFunds(ctx context.Context, searchTerm string) ([]fund.Fund, error) {
	sql, args := SELECT(Fund.ID, Fund.Name).
		FROM(Fund.
			INNER_JOIN(FundListing, FundListing.FundID.EQ(Fund.ID)),
		).
		WHERE(ILike(Fund.Name, searchTerm).
			OR(ILike(FundListing.Ticker, searchTerm)),
		).
		DISTINCT(Fund.ID).
		LIMIT(int64(10)).
		Sql()

	var f []fund.Fund
	err := pgxscan.Select(ctx, r.ConnectionPool, &f, sql, args...)
	if err != nil {
		return nil, err
	}
	return f, nil
}
func (r *Repository) GetFund(ctx context.Context, fundId uuid.UUID) (fund.Information, error) {
	sql, args := SELECT(
		Fund.Name,
		Fund.OutstandingShares,
		Fund.EffectiveDate,
	).
		FROM(Fund).
		WHERE(Fund.ID.EQ(UUID(fundId))).
		Sql()

	var fi fund.Information
	err := pgxscan.Get(ctx, r.ConnectionPool, &fi, sql, args...)
	if err != nil {
		return fund.Information{}, err
	}
	return fi, nil
}

func (r *Repository) UpsertFund(ctx context.Context, f model.Fund, tx pgx.Tx) (uuid.UUID, error) {
	sql, args := Fund.
		INSERT(Fund.MutableColumns).
		MODEL(f).
		ON_CONFLICT(Fund.Isin).
		DO_UPDATE(
			SET(
				Fund.TotalHoldings.SET(Fund.EXCLUDED.TotalHoldings),
				Fund.Price.SET(Fund.EXCLUDED.Price),
			),
		).
		ON_CONFLICT(Fund.ExternalIdentifier).
		DO_UPDATE(
			SET(
				Fund.TotalHoldings.SET(Fund.EXCLUDED.TotalHoldings),
				Fund.Price.SET(Fund.EXCLUDED.Price),
			),
		).
		RETURNING(Fund.ID).
		Sql()

	row := tx.QueryRow(ctx, sql, args...)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
