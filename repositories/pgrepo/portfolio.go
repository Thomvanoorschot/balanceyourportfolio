package pgrepo

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
)

func (r *Repository) CreatePortfolio(ctx context.Context, p model.Portfolio) error {
	sql, args := Portfolio.
		INSERT(Portfolio.MutableColumns).
		MODEL(p).
		Sql()

	_, err := r.ConnectionPool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
