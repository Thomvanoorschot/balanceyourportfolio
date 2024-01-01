package pgrepo

import (
	"context"

	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
	"etfinsight/services/fund"

	"github.com/georgysavva/scany/v2/pgxscan"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) FilterFunds(ctx context.Context, filter fund.FundsFilter) (fund.Funds, error) {
	queryCte := CTE("queryCte")
	queryStmt := SELECT(Fund.ID,
		Fund.Name,
		Fund.Currency,
		FundListing.Ticker,
		Fund.Price.MUL(Fund.OutstandingShares).AS("marketCap"),
		Fund.Price.MUL(Fund.OutstandingShares).MUL(Currency.ExchangeRate).AS("relativeMarketCap"),
	).
		FROM(Fund.
			INNER_JOIN(Currency, Currency.Code.EQ(Fund.Currency)).
			LEFT_JOIN(FundListing, FundListing.FundID.EQ(Fund.ID)),
		).
		WHERE(
			ILike(Fund.Name, filter.SearchTerm).
				OR(ILike(FundListing.Ticker, filter.SearchTerm).
					OR(Fund.Isin.EQ(String(filter.SearchTerm)))),
		).ORDER_BY(Raw(`"relativeMarketCap"`).DESC())

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
		)
		err := rows.Scan(&fundId,
			&fundName,
			&fundCurrency,
			&fundTicker,
			&fundMarketCap,
			&relativeMarketCap,
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
