package pgrepo

import (
	"context"
	"time"

	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
	"etfinsight/services/portfolio"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

func (r *Repository) GetPortfolios(ctx context.Context, userId string) (portfolio.Models, error) {
	sql, args := SELECT(
		Portfolio.ID,
		Portfolio.Name,
		PortfolioFund.ID,
		PortfolioFund.Amount,
		Fund.ID,
		Fund.Name,
	).
		FROM(Portfolio.
			INNER_JOIN(PortfolioFund, PortfolioFund.PortfolioID.EQ(Portfolio.ID)).
			INNER_JOIN(Fund, Fund.ID.EQ(PortfolioFund.FundID)),
		).
		ORDER_BY(Portfolio.CreatedAt.ASC()).
		WHERE(Portfolio.UserID.EQ(String(userId))).
		Sql()

	var models portfolio.Models
	rows, err := r.ConnectionPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			portfolioId         uuid.UUID
			portfolioName       string
			portfolioFundId     uuid.UUID
			portfolioFundAmount float64
			fundID              uuid.UUID
			fundName            string
		)
		err := rows.Scan(&portfolioId,
			&portfolioName,
			&portfolioFundId,
			&portfolioFundAmount,
			&fundID,
			&fundName,
		)
		if err != nil {
			return nil, err
		}
		if len(models) == 0 || models[len(models)-1].Id != portfolioId {
			models = append(models, portfolio.Model{
				Id:   portfolioId,
				Name: portfolioName,
				Items: []portfolio.ListItem{{
					Id:     portfolioFundId,
					Amount: portfolioFundAmount,
					FundID: fundID,
					Name:   fundName,
				}},
			})
			continue
		}
		models[len(models)-1].Items = append(models[len(models)-1].Items, portfolio.ListItem{
			Id:     portfolioFundId,
			Amount: portfolioFundAmount,
			FundID: fundID,
			Name:   fundName,
		})
	}
	return models, nil
}

func (r *Repository) UpsertPortfolio(ctx context.Context,
	p model.Portfolio,
	tx pgx.Tx) error {
	now := time.Now()
	p.UpdatedAt = &now
	sql, args := Portfolio.
		INSERT(Portfolio.AllColumns).
		MODEL(p).
		ON_CONFLICT(Portfolio.ID).
		DO_UPDATE(
			SET(
				Portfolio.Name.SET(Portfolio.EXCLUDED.Name),
				Portfolio.UpdatedAt.SET(Portfolio.EXCLUDED.UpdatedAt),
			),
		).
		Sql()

	_, err := tx.Exec(ctx, sql, args...)
	return err
}
func (r *Repository) GetPortfolioOwner(ctx context.Context, portfolioId uuid.UUID) (string, error) {
	sql, args :=
		SELECT(Portfolio.UserID).
			FROM(Portfolio).
			WHERE(Portfolio.ID.EQ(UUID(portfolioId))).
			Sql()

	var userId string
	row := r.ConnectionPool.QueryRow(ctx, sql, args...)
	err := row.Scan(&userId)
	if err != nil {
		return "", err
	}
	return userId, nil
}
