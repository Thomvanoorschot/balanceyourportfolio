package pgrepo

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) UpsertFundHoldings(ctx context.Context, s []model.FundHolding, tx pgx.Tx) error {
	sql, args := FundHolding.
		INSERT(FundHolding.MutableColumns).
		MODELS(s).
		ON_CONFLICT(FundHolding.FundID, FundHolding.HoldingID).
		DO_UPDATE(
			SET(
				FundHolding.MarketValue.SET(FundHolding.EXCLUDED.MarketValue),
				FundHolding.PercentageOfTotal.SET(FundHolding.EXCLUDED.PercentageOfTotal),
				FundHolding.Amount.SET(FundHolding.EXCLUDED.Amount),
			),
		).
		Sql()

	_, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

//Select all funds that have ASML in top 10
//SELECT rank_filter.* FROM (
//SELECT h.ticker,fh.*,
//rank() OVER (
//PARTITION BY fh.fund_id
//ORDER BY percentage_of_total DESC
//)
//FROM fund_holding fh
//INNER JOIN holding h ON h.id = fh.holding_id
//) rank_filter WHERE ticker = 'ASML' AND rank < 10;