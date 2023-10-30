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

func (r *Repository) GetFundHoldings(ctx context.Context, fundId uuid.UUID, limit int64) (fund.Holdings, error) {
	return r.getHoldings(ctx, func(statement SelectStatement) SelectStatement {
		return statement.
			WHERE(Fund.ID.EQ(UUID(fundId))).
			ORDER_BY(FundHolding.MarketValue.DESC()).
			LIMIT(limit)
	})
}
func (r *Repository) FilterHoldings(ctx context.Context, filter fund.HoldingsFilter) (fund.Holdings, error) {
	var filterExp BoolExpression
	if filter.FundID != uuid.Nil {
		if filterExp == nil {
			filterExp = Fund.ID.EQ(UUID(filter.FundID))
		}
	}
	if filter.SectorName != "" {
		if filterExp == nil {
			filterExp = Holding.Sector.EQ(String(filter.SectorName))
		} else {
			filterExp = filterExp.AND(Holding.Sector.EQ(String(filter.SectorName)))
		}
	}
	if filter.SearchTerm != "" {
		if filterExp == nil {
			filterExp = ILike(Holding.Ticker, filter.SearchTerm).OR(ILike(Holding.Name, filter.SearchTerm))
		} else {
			filterExp = filterExp.AND(ILike(Holding.Ticker, filter.SearchTerm).OR(ILike(Holding.Name, filter.SearchTerm)))
		}
	}
	return r.getHoldings(ctx, func(statement SelectStatement) SelectStatement {
		return statement.
			WHERE(filterExp).
			ORDER_BY(FundHolding.MarketValue.DESC()).
			LIMIT(filter.Limit).
			OFFSET(filter.Offset)
	})
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

func (r *Repository) UpsertHoldings(ctx context.Context, holdings []model.Holding, tx pgx.Tx) (map[string]uuid.UUID, error) {
	insertCte := CTE("insert_cte")

	insertStmt := Holding.
		INSERT(Holding.MutableColumns).
		MODELS(holdings).
		ON_CONFLICT(Holding.Ticker).
		DO_NOTHING().
		RETURNING(Holding.ID, Holding.Ticker)
	withStmt := WITH(
		insertCte.AS(
			insertStmt,
		),
	)
	var tickerExpr []Expression
	for _, h := range holdings {
		tickerExpr = append(tickerExpr, String(*h.Ticker))
	}
	selectStmt := SELECT(Holding.ID, Holding.Ticker).FROM(Holding).WHERE(Holding.Ticker.IN(tickerExpr...))
	sql, args := withStmt(UNION_ALL(SELECT(STAR).FROM(insertCte), selectStmt)).
		Sql()
	rows, err := tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	holdingMap := map[string]uuid.UUID{}
	_, err = pgx.CollectRows(rows, func(row pgx.CollectableRow) (uuid.UUID, error) {
		var id uuid.UUID
		var ticker string
		err := row.Scan(&id, &ticker)
		holdingMap[ticker] = id
		return id, err
	})
	if err != nil {
		return nil, err
	}

	return holdingMap, nil
}

func (r *Repository) getHoldings(ctx context.Context, stmt func(SelectStatement) SelectStatement) (fund.Holdings, error) {
	sql, args := stmt(SELECT(
		FundHolding.Amount,
		FundHolding.MarketValue,
		FundHolding.PercentageOfTotal,
		Holding.Ticker,
		Holding.Name,
		Holding.Type,
		Holding.Sector,
	).
		FROM(Fund.
			INNER_JOIN(FundHolding, FundHolding.FundID.EQ(Fund.ID)).
			INNER_JOIN(Holding, Holding.ID.EQ(FundHolding.HoldingID)),
		)).Sql()

	var h []fund.Holding
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
