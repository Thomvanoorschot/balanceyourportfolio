package pgrepo

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
	"etfinsight/services/portfolio"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

func (r *Repository) GetPortfolios(ctx context.Context, userID uuid.UUID) (portfolio.Models, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) UpsertPortfolio(ctx context.Context,
	userID uuid.UUID,
	p portfolio.Model,
	tx pgx.Tx) (portfolio.Model, error) {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	pm := model.Portfolio{
		ID:     p.ID,
		UserID: &userID,
		Name:   &p.Name,
	}
	sql, args := Portfolio.
		INSERT(Portfolio.AllColumns).
		MODEL(pm).
		ON_CONFLICT(Portfolio.ID).
		DO_UPDATE(
			SET(
				Portfolio.Name.SET(Portfolio.EXCLUDED.Name),
			),
		).
		Sql()

	_, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return portfolio.Model{}, err
	}

	return p, err
}
