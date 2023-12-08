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

func (r *Repository) GetPortfolios(ctx context.Context, userID uuid.UUID) (portfolio.Models, error) {
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
		WHERE(Portfolio.UserID.EQ(UUID(userID))).
		Sql()

	var models portfolio.Models
	rows, err := r.ConnectionPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			portfolioID         uuid.UUID
			portfolioName       string
			portfolioFundID     uuid.UUID
			portfolioFundAmount float64
			fundID              uuid.UUID
			fundName            string
		)
		err := rows.Scan(&portfolioID,
			&portfolioName,
			&portfolioFundID,
			&portfolioFundAmount,
			&fundID,
			&fundName,
		)
		if err != nil {
			return nil, err
		}
		if len(models) == 0 || models[len(models)-1].Id != portfolioID {
			models = append(models, portfolio.Model{
				Id:   portfolioID,
				Name: portfolioName,
				Items: []portfolio.ListItem{{
					Id:     portfolioFundID,
					Amount: portfolioFundAmount,
					FundID: fundID,
					Name:   fundName,
				}},
			})
			continue
		}
		models[len(models)-1].Items = append(models[len(models)-1].Items, portfolio.ListItem{
			Id:     portfolioFundID,
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
