package pgrepo

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
	"etfinsight/services/fund"
	"etfinsight/services/portfolio"

	"github.com/georgysavva/scany/v2/pgxscan"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

func (r *Repository) GetListItems(ctx context.Context,
	portfolioID uuid.UUID,
) (portfolio.ListItems, error) {
	sql, args := SELECT(PortfolioFund.ID, PortfolioFund.Amount, Fund.ID, Fund.Name).
		FROM(PortfolioFund.
			INNER_JOIN(Fund, Fund.ID.EQ(PortfolioFund.FundID)),
		).
		WHERE(PortfolioFund.PortfolioID.EQ(UUID(portfolioID))).
		Sql()

	var li portfolio.ListItems
	err := pgxscan.Select(ctx, r.ConnectionPool, &li, sql, args...)
	if err != nil {
		return nil, err
	}
	return li, nil
}
func (r *Repository) DeleteListItems(ctx context.Context,
	ids []uuid.UUID,
	tx pgx.Tx,
) error {
	var IDExpression []Expression
	for _, h := range ids {
		IDExpression = append(IDExpression, UUID(h))
	}
	sql, args := PortfolioFund.
		DELETE().
		WHERE(PortfolioFund.ID.IN(IDExpression...)).
		Sql()

	_, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) UpsertPortfolioListItems(ctx context.Context,
	pfs []model.PortfolioFund,
	tx pgx.Tx,
) error {
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
	return err
}

func (r *Repository) GetPortfolioFundSectors(ctx context.Context, portfolioID uuid.UUID) ([]fund.SectorName, error) {
	sql, args := SELECT(DISTINCT(Holding.Sector)).
		FROM(Holding.
			INNER_JOIN(FundHolding, FundHolding.HoldingID.EQ(Holding.ID)).
			INNER_JOIN(PortfolioFund, PortfolioFund.FundID.EQ(FundHolding.FundID)),
		).
		WHERE(PortfolioFund.PortfolioID.EQ(UUID(portfolioID))).
		Sql()
	var h []fund.SectorName
	err := pgxscan.Select(ctx, r.ConnectionPool, &h, sql, args...)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (r *Repository) GetPortfolioFundRelativeWeightings(ctx context.Context, portfolioID uuid.UUID) (portfolio.RelativeSectorWeightings, error) {
	sql, args := SELECT(
		PortfolioFund.FundID,
		PortfolioFund.Amount,
		Fund.Price,
		Fund.Name,
		Holding.Sector,
		SUM(FundHolding.PercentageOfTotal).AS("percentage_sum"),
	).
		FROM(Holding.
			INNER_JOIN(FundHolding, FundHolding.HoldingID.EQ(Holding.ID)).
			INNER_JOIN(PortfolioFund, PortfolioFund.FundID.EQ(FundHolding.FundID)).
			INNER_JOIN(Fund, PortfolioFund.FundID.EQ(Fund.ID)),
		).
		WHERE(PortfolioFund.PortfolioID.EQ(UUID(portfolioID))).
		GROUP_BY(
			PortfolioFund.FundID,
			PortfolioFund.Amount,
			Fund.Price,
			Fund.Name,
			Holding.Sector,
		).
		ORDER_BY(PortfolioFund.FundID, Raw("percentage_sum").DESC()).
		Sql()
	var rw portfolio.RelativeSectorWeightings
	rows, err := r.ConnectionPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var (
			fundID                     uuid.UUID
			fundName                   string
			portfolioFundAmount        float64
			fundPrice                  float64
			sectorName                 fund.SectorName
			sectorWeightingsPercentage float64
		)
		err = rows.Scan(
			&fundID,
			&portfolioFundAmount,
			&fundPrice,
			&fundName,
			&sectorName,
			&sectorWeightingsPercentage,
		)
		if err != nil {
			return nil, err
		}
		if len(rw) == 0 || rw[len(rw)-1].FundID != fundID {
			rw = append(rw, portfolio.RelativeSectorWeighting{
				FundID:              fundID,
				FundName:            fundName,
				FundPrice:           fundPrice,
				PortfolioFundAmount: portfolioFundAmount,
				SectorWeightings: []fund.SectorWeighting{
					{
						SectorName: sectorName,
						Percentage: sectorWeightingsPercentage,
					},
				},
			})
			continue
		}
		rw[len(rw)-1].SectorWeightings = append(rw[len(rw)-1].SectorWeightings, fund.SectorWeighting{
			SectorName: sectorName,
			Percentage: sectorWeightingsPercentage,
		})
	}
	return rw, nil
}

func (r *Repository) GetPortfolioFunds(ctx context.Context, portfolioID uuid.UUID) ([]fund.Information, error) {
	sql, args := SELECT(
		Fund.ID,
		Fund.Name,
		Fund.OutstandingShares,
		Fund.EffectiveDate,
	).
		DISTINCT(Fund.ID).
		FROM(Fund.
			INNER_JOIN(PortfolioFund, PortfolioFund.FundID.EQ(Fund.ID)),
		).
		WHERE(PortfolioFund.PortfolioID.EQ(UUID(portfolioID))).
		Sql()

	var fi []fund.Information
	err := pgxscan.Select(ctx, r.ConnectionPool, &fi, sql, args...)
	if err != nil {
		return nil, err
	}
	return fi, nil
}
