package pgrepo

import (
	"context"

	"balanceyourportfolio/generated/jet_gen/postgres/public/model"
	. "balanceyourportfolio/generated/jet_gen/postgres/public/table"
	"balanceyourportfolio/services/fund"

	"github.com/georgysavva/scany/v2/pgxscan"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) FilterFunds(ctx context.Context, filter fund.FundsFilter) (fund.Funds, error) {
	queryCte := CTE("queryCte")
	var providerExpression []Expression
	for _, p := range filter.Providers {
		providerExpression = append(providerExpression, String(p))
	}
	whereExpr := Fund.TotalHoldings.GT(Float(0)).AND(ILike(Fund.Name, filter.SearchTerm).
		OR(ILike(FundListing.Ticker, filter.SearchTerm).
			OR(Fund.Isin.EQ(String(filter.SearchTerm)))))
	if len(providerExpression) > 0 {
		whereExpr = whereExpr.
			AND(Fund.Provider.IN(providerExpression...))
	}
	queryStmt := SELECT(Fund.ID,
		Fund.Name,
		Fund.Currency,
		FundListing.Ticker,
		Fund.Price.MUL(Fund.OutstandingShares).AS("marketCap"),
		Fund.Price.MUL(Fund.OutstandingShares).MUL(Currency.ExchangeRate).AS("relativeMarketCap"),
		Fund.Provider,
	).
		FROM(Fund.
			INNER_JOIN(Currency, Currency.Code.EQ(Fund.Currency)).
			LEFT_JOIN(FundListing, FundListing.FundID.EQ(Fund.ID)),
		).
		WHERE(whereExpr).
		ORDER_BY(Raw(`"relativeMarketCap"`).DESC())

	limitCte := CTE("limitCte")
	limitStmt := SELECT(DISTINCT(StringColumn("fund.id")), Raw(`"relativeMarketCap"`)).
		FROM(queryCte).
		ORDER_BY(Raw(`"relativeMarketCap"`).DESC()).
		LIMIT(filter.Limit).
		OFFSET(filter.Offset)

	withStmt := WITH(
		queryCte.AS(
			queryStmt,
		),
		limitCte.
			AS(
				limitStmt,
			),
	)
	sql, args := withStmt(SELECT(STAR).FROM(queryCte).WHERE(StringColumn("fund.id").IN(SELECT(StringColumn("fund.id")).FROM(limitCte)))).
		Sql()
	var funds []fund.Fund
	rows, err := r.ConnectionPool.Query(ctx, sql, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			fundId            uuid.UUID
			fundName          string
			fundCurrency      string
			fundTicker        *string
			fundMarketCap     float64
			relativeMarketCap float64
			provider          string
		)
		err := rows.Scan(&fundId,
			&fundName,
			&fundCurrency,
			&fundTicker,
			&fundMarketCap,
			&relativeMarketCap,
			&provider,
		)
		if err != nil {
			return nil, err
		}
		if len(funds) == 0 || funds[len(funds)-1].ID != fundId {
			newFund := fund.Fund{
				ID:        fundId,
				Name:      fundName,
				Currency:  fundCurrency,
				MarketCap: fundMarketCap,
				Provider:  provider,
			}
			if fundTicker != nil {
				newFund.Tickers = append(newFund.Tickers, *fundTicker)
			}
			funds = append(funds, newFund)
			continue
		}
		f := &funds[len(funds)-1]
		if fundTicker != nil {
			f.Tickers = append(f.Tickers, *fundTicker)
		}
	}
	return funds, nil
}

func (r *Repository) GetFund(ctx context.Context, fundId uuid.UUID) (fund.Information, error) {
	sql, args := SELECT(
		Fund.ID,
		Fund.Name,
		Fund.OutstandingShares,
		Fund.EffectiveDate,
	).
		FROM(Fund).
		WHERE(Fund.ID.EQ(UUID(fundId))).
		Sql()

	var fi fund.Information
	err := pgxscan.Get(ctx, r.ConnectionPool, &fi, sql, args...)
	if err != nil {
		return fund.Information{}, err
	}
	return fi, nil
}

func (r *Repository) UpsertFund(ctx context.Context, f model.Fund, tx pgx.Tx) (uuid.UUID, error) {
	sql, args := Fund.
		INSERT(Fund.MutableColumns).
		MODEL(f).
		ON_CONFLICT(Fund.Isin).
		DO_UPDATE(
			SET(
				Fund.TotalHoldings.SET(Fund.EXCLUDED.TotalHoldings),
				Fund.Price.SET(Fund.EXCLUDED.Price),
				Fund.OutstandingShares.SET(Fund.EXCLUDED.OutstandingShares),
				Fund.EffectiveDate.SET(Fund.EXCLUDED.EffectiveDate),
				Fund.Name.SET(Fund.EXCLUDED.Name),
				Fund.Currency.SET(Fund.EXCLUDED.Currency),
			),
		).
		ON_CONFLICT(Fund.ExternalIdentifier).
		DO_UPDATE(
			SET(
				Fund.TotalHoldings.SET(Fund.EXCLUDED.TotalHoldings),
				Fund.Price.SET(Fund.EXCLUDED.Price),
				Fund.OutstandingShares.SET(Fund.EXCLUDED.OutstandingShares),
				Fund.EffectiveDate.SET(Fund.EXCLUDED.EffectiveDate),
				Fund.Name.SET(Fund.EXCLUDED.Name),
				Fund.Currency.SET(Fund.EXCLUDED.Currency),
			),
		).
		RETURNING(Fund.ID).
		Sql()

	row := tx.QueryRow(ctx, sql, args...)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

func (r *Repository) GetFundSectors(ctx context.Context, fundId uuid.UUID) ([]fund.SectorName, error) {
	sql, args := SELECT(DISTINCT(Holding.Sector)).
		FROM(Holding.
			INNER_JOIN(FundHolding, FundHolding.HoldingID.EQ(Holding.ID)),
		).
		WHERE(FundHolding.FundID.EQ(UUID(fundId))).
		Sql()
	var h []fund.SectorName
	err := pgxscan.Select(ctx, r.ConnectionPool, &h, sql, args...)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (r *Repository) GetFundSectorWeightings(ctx context.Context, fundId uuid.UUID) ([]fund.SectorWeighting, error) {
	sql, args := SELECT(Holding.Sector, SUM(FundHolding.PercentageOfTotal).AS("percentage_sum")).
		FROM(Holding.
			INNER_JOIN(FundHolding, FundHolding.HoldingID.EQ(Holding.ID)),
		).
		WHERE(FundHolding.FundID.EQ(UUID(fundId))).
		GROUP_BY(Holding.Sector).
		ORDER_BY(Raw("percentage_sum").DESC()).
		Sql()
	var sw []fund.SectorWeighting
	err := pgxscan.Select(ctx, r.ConnectionPool, &sw, sql, args...)
	if err != nil {
		return nil, err
	}
	return sw, nil
}

func (r *Repository) GetFundsSectorWeightings(ctx context.Context, fundOne, fundTwo uuid.UUID) (resp []fund.SectorWeighting, err error) {
	sql, args := RawStatement(
		`
		WITH SectorSums AS (
			SELECT 
				 fund_holding.fund_id AS "fund_holding.fund_id",
				 holding.sector AS "holding.sector",
				 SUM(fund_holding.percentage_of_total) AS "percentage_sum"
			FROM 
				 public.holding
				 INNER JOIN public.fund_holding ON fund_holding.holding_id = holding.id
			WHERE 
				 fund_holding.fund_id IN (:fund1, :fund2)
			GROUP BY 
				 fund_holding.fund_id, holding.sector
		),
		MaxPercentagePerSector AS (
			SELECT 
				 "holding.sector",
				 MAX("percentage_sum") OVER (PARTITION BY "holding.sector") as max_percentage_in_sector
			FROM 
				 SectorSums
		)
		SELECT distinct on (mps.max_percentage_in_sector, ss."fund_holding.fund_id", ss."holding.sector")
			ss."fund_holding.fund_id",
			ss."holding.sector",
			ss."percentage_sum"
		FROM 
			SectorSums ss
			INNER JOIN MaxPercentagePerSector mps ON ss."holding.sector" = mps."holding.sector"
		ORDER BY 
			mps.max_percentage_in_sector DESC,
			ss."fund_holding.fund_id", 
			ss."holding.sector";
`,
		RawArgs{
			":fund1": fundOne,
			":fund2": fundTwo,
		},
	).Sql()

	rows, err := r.ConnectionPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		var (
			fundId     uuid.UUID
			sector     string
			percentage float64
		)
		err = rows.Scan(
			&fundId,
			&sector,
			&percentage,
		)
		if err != nil {
			return nil, err
		}

		resp = append(resp, fund.SectorWeighting{
			FundId:     fundId,
			SectorName: fund.SectorName(sector),
			Percentage: percentage,
		})
	}
	return resp, nil
}

// Get highest ranking sector for fund
//WITH R as (
//select FH.fund_id, H.sector, RANK() over (partition by FH.fund_id order by SUM(FH.percentage_of_total) desc)
//FROM holding H
//INNER JOIN fund_holding FH on FH.holding_id = H.id
//group by H.sector, FH.fund_id
//order by FH.fund_id
//)
//SELECT F."name", R.sector from fund F
//INNER JOIN R on R.fund_id = F.id AND R.RANK = 1;
