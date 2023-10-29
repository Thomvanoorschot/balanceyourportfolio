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

func (r *Repository) UpsertPortfolioListItems(ctx context.Context,
	portfolioID uuid.UUID,
	li portfolio.ListItems,
	tx pgx.Tx,
) (portfolio.ListItems, error) {
	var pfs []model.PortfolioFund
	for i := range li {
		if li[i].ID == uuid.Nil {
			li[i].ID = uuid.New()
		}
		pf := model.PortfolioFund{
			ID:          li[i].ID,
			PortfolioID: &portfolioID,
			FundID:      &li[i].FundID,
			Amount:      &li[i].Amount,
		}
		pfs = append(pfs, pf)
	}

	sql, args := PortfolioFund.
		INSERT(PortfolioFund.AllColumns).
		MODELS(pfs).
		ON_CONFLICT(PortfolioFund.ID).
		DO_UPDATE(
			SET(
				PortfolioFund.FundID.SET(PortfolioFund.EXCLUDED.FundID),
				PortfolioFund.Amount.SET(PortfolioFund.EXCLUDED.Amount),
			),
		).
		Sql()

	_, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return portfolio.ListItems{}, err
	}
	return li, err
}
