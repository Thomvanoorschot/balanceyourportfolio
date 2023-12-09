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

func (r *Repository) GetRatio(ctx context.Context, portfolioId uuid.UUID) (map[uuid.UUID]float64, error) {
	sql, args := SELECT(Fund.ID, Fund.Price.MUL(PortfolioFund.Amount).DIV(SUMf(Fund.Price.MUL(PortfolioFund.Amount)).OVER()).AS("ratio")).
		FROM(PortfolioFund.
			INNER_JOIN(Fund, Fund.ID.EQ(PortfolioFund.FundID))).
		WHERE(PortfolioFund.PortfolioID.EQ(UUID(portfolioId))).
		Sql()
	rows, err := r.ConnectionPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ratios := map[uuid.UUID]float64{}
	for rows.Next() {
		var (
			fundId uuid.UUID
			ratio  float64
		)
		err = rows.Scan(
			&fundId,
			&ratio,
		)
		if err != nil {
			return nil, err
		}
		ratios[fundId] = ratio
	}
	return ratios, nil
}
func (r *Repository) GetPortfolioFundRelativeWeightings(ctx context.Context, portfolioId uuid.UUID) (portfolio.RelativeSectorWeightings, error) {
	sql, args := SELECT(
		PortfolioFund.FundID,
		Fund.Name,
		Holding.Sector,
		SUM(FundHolding.PercentageOfTotal).AS("percentage_sum"),
	).
		FROM(Holding.
			INNER_JOIN(FundHolding, FundHolding.HoldingID.EQ(Holding.ID)).
			INNER_JOIN(PortfolioFund, PortfolioFund.FundID.EQ(FundHolding.FundID)).
			INNER_JOIN(Fund, PortfolioFund.FundID.EQ(Fund.ID)),
		).
		WHERE(PortfolioFund.PortfolioID.EQ(UUID(portfolioId))).
		GROUP_BY(
			PortfolioFund.FundID,
			PortfolioFund.Amount,
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
			sectorName                 fund.SectorName
			sectorWeightingsPercentage float64
		)
		err = rows.Scan(
			&fundID,
			&fundName,
			&sectorName,
			&sectorWeightingsPercentage,
		)
		if err != nil {
			return nil, err
		}
		if len(rw) == 0 || rw[len(rw)-1].FundID != fundID {
			rw = append(rw, portfolio.RelativeSectorWeighting{
				FundID:   fundID,
				FundName: fundName,
				SectorWeightings: []portfolio.SectorWeighting{
					{
						SectorWeighting: fund.SectorWeighting{
							SectorName: sectorName,
							Percentage: sectorWeightingsPercentage,
						},
					},
				},
			})
			continue
		}
		rw[len(rw)-1].SectorWeightings = append(rw[len(rw)-1].SectorWeightings, portfolio.SectorWeighting{
			SectorWeighting: fund.SectorWeighting{
				SectorName: sectorName,
				Percentage: sectorWeightingsPercentage,
			},
		})
	}
	return rw, nil
}

func (r *Repository) GetPortfolioFunds(ctx context.Context, portfolioId uuid.UUID) ([]fund.Information, error) {
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
		WHERE(PortfolioFund.PortfolioID.EQ(UUID(portfolioId))).
		Sql()

	var fi []fund.Information
	err := pgxscan.Select(ctx, r.ConnectionPool, &fi, sql, args...)
	if err != nil {
		return nil, err
	}
	return fi, nil
}
func (r *Repository) GetPortfolioFundHoldings(ctx context.Context, portfolioId uuid.UUID, limit int64, offset int64) (portfolio.FundHoldings, error) {
	sql, args := RawStatement(
		`WITH ratio_cte AS (
     SELECT fund.id AS "fund.id",
          ((fund.price * portfolio_fund.amount) / SUM(fund.price * portfolio_fund.amount) OVER ()) AS "ratio"
     FROM public.portfolio_fund
          INNER JOIN public.fund ON (fund.id = portfolio_fund.fund_id)
     WHERE portfolio_fund.portfolio_id = :portfolioId
	),
	relative_weightings_cte AS (
		 SELECT H.Id as "holdingId", H.ticker, H."name", (ratio_cte.ratio * FH.percentage_of_total) as ratiodPercentage, F.id as "fundId"
		 FROM holding H
			  INNER JOIN fund_holding FH ON (FH.holding_id = H.id)
			  INNER JOIN portfolio_fund PF ON (PF.fund_id = FH.fund_id)
			  INNER JOIN fund F ON (PF.fund_id = F.id)
			  INNER JOIN ratio_cte ON (ratio_cte."fund.id" = F.id)
		 WHERE PF.portfolio_id = :portfolioId
	), limiting_cte as (
		SELECT distinct("holdingId"), (SUM(ratiodPercentage) OVER(PARTITION BY ticker)) as cumulativePercentage 
		from relative_weightings_cte 
		order by cumulativePercentage desc
		limit :limit 
		offset :offset
	) SELECT *, (SUM(ratiodPercentage) OVER(PARTITION BY ticker)) as cumulativePercentage from relative_weightings_cte 
	where "holdingId" in (select "holdingId" from limiting_cte)
	order by cumulativePercentage desc`, RawArgs{":portfolioId": portfolioId, ":limit": limit, ":offset": offset},
	).
		Sql()

	rows, err := r.ConnectionPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var fh portfolio.FundHoldings
	for rows.Next() {
		var (
			holdingId            uuid.UUID
			ticker               string
			holdingName          string
			ratiodPercentage     float64
			fundId               uuid.UUID
			cumulativePercentage float64
		)
		err = rows.Scan(
			&holdingId,
			&ticker,
			&holdingName,
			&ratiodPercentage,
			&fundId,
			&cumulativePercentage,
		)
		if err != nil {
			return nil, err
		}
		if len(fh) == 0 || fh[len(fh)-1].Ticker != ticker {
			fh = append(fh, portfolio.FundHolding{
				Ticker:               ticker,
				HoldingId:            holdingId,
				HoldingName:          holdingName,
				CumulativePercentage: cumulativePercentage,
				Funds: []portfolio.FundHoldingEntry{
					{
						FundId:          fundId,
						RatiodPerentage: ratiodPercentage,
					},
				},
			})
			continue
		}
		fh[len(fh)-1].Funds = append(fh[len(fh)-1].Funds, portfolio.FundHoldingEntry{
			FundId:          fundId,
			RatiodPerentage: ratiodPercentage,
		})
	}
	return fh, nil
}
