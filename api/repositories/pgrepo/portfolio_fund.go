package pgrepo

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
)

func (r *Repository) UpsertPortfolioFunds(ctx context.Context, pf []model.PortfolioFund) error {
	sql, args := PortfolioFund.
		INSERT(Portfolio.MutableColumns).
		MODELS(pf).
		Sql()

	_, err := r.ConnectionPool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
