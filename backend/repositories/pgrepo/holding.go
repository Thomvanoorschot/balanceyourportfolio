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
	if filter.FundId != uuid.Nil {
		if filterExp == nil {
			filterExp = Fund.ID.EQ(UUID(filter.FundId))
		}
	}
	if len(filter.SelectedSectors) > 0 {
		var sectorExpression []Expression
		for _, selectedSector := range filter.SelectedSectors {
			sectorExpression = append(sectorExpression, String(selectedSector))
		}
		if filterExp == nil {
			filterExp = Holding.Sector.IN(sectorExpression...)
		} else {
			filterExp = filterExp.AND(Holding.Sector.IN(sectorExpression...))
		}
	}
	if filter.SearchTerm != "" {
		if filterExp == nil {
			filterExp = ILike(FigiMapping.Ticker, filter.SearchTerm).OR(ILike(FigiMapping.Name, filter.SearchTerm))
		} else {
			filterExp = filterExp.AND(ILike(FigiMapping.Ticker, filter.SearchTerm).OR(ILike(FigiMapping.Name, filter.SearchTerm)))
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

func (r *Repository) UpsertHoldings(ctx context.Context, holdings []model.Holding, tx pgx.Tx) (map[string]uuid.UUID, error) {
	insertCte := CTE("insert_cte")

	insertStmt := Holding.
		INSERT(Holding.MutableColumns).
		MODELS(holdings).
		ON_CONFLICT(Holding.Figi).
		DO_NOTHING().
		RETURNING(Holding.ID, Holding.Figi)
	withStmt := WITH(
		insertCte.AS(
			insertStmt,
		),
	)
	var figiExpr []Expression
	for _, h := range holdings {
		figiExpr = append(figiExpr, String(*h.Figi))
	}
	selectStmt := SELECT(Holding.ID, Holding.Figi).FROM(Holding).WHERE(Holding.Figi.IN(figiExpr...))
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
		var figi string
		err := row.Scan(&id, &figi)
		holdingMap[figi] = id
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
		FigiMapping.Ticker,
		FigiMapping.Name,
		Holding.Type,
		Holding.Sector,
	).
		FROM(Fund.
			INNER_JOIN(FundHolding, FundHolding.FundID.EQ(Fund.ID)).
			INNER_JOIN(Holding, Holding.ID.EQ(FundHolding.HoldingID)).
			INNER_JOIN(FigiMapping, FigiMapping.Figi.EQ(Holding.Figi)),
		)).Sql()

	var h []fund.Holding
	err := pgxscan.Select(ctx, r.ConnectionPool, &h, sql, args...)
	if err != nil {
		return nil, err
	}
	return h, nil
}
