package pgrepo

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
)

func (r *Repository) CreateUser(ctx context.Context, u model.User) error {
	sql, args := User.
		INSERT(User.MutableColumns).
		MODEL(u).
		Sql()

	_, err := r.ConnectionPool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
