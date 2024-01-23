package pgrepo

import (
	"context"
	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
	"etfinsight/services/fund"
	"github.com/google/uuid"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) UpsertFundHoldings(ctx context.Context, s []model.FundHolding, tx pgx.Tx) error {
	sql, args := FundHolding.
		INSERT(FundHolding.MutableColumns).
		MODELS(s).
		ON_CONFLICT(FundHolding.FundID, FundHolding.HoldingID).
		DO_UPDATE(
			SET(
				FundHolding.MarketValue.SET(FundHolding.EXCLUDED.MarketValue),
				FundHolding.PercentageOfTotal.SET(FundHolding.EXCLUDED.PercentageOfTotal),
				FundHolding.Amount.SET(FundHolding.EXCLUDED.Amount),
			),
		).
		Sql()

	_, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) GetTotalOverlap(ctx context.Context, fundOne, fundTwo uuid.UUID) (resp fund.OverlappingFunds, err error) {
	sql, args := RawStatement(
		`
	WITH CTE AS (
	SELECT 
		fh1.fund_id, 
		COUNT(DISTINCT fh1.holding_id) AS overlapping_holdings_count
	FROM fund_holding fh1
	INNER JOIN fund_holding fh2 ON fh1.holding_id = fh2.holding_id
	WHERE fh1.fund_id = :fund1
	AND fh2.fund_id = :fund2
	GROUP BY fh1.fund_id
	)
	SELECT
		SUM(COALESCE(
			LEAST(coalesce(fh1.percentage_of_total, 0), coalesce(fh2.percentage_of_total, 0)),
			0
		))AS total_overlap,
		c.overlapping_holdings_count,
		f1.total_holdings,
		c.overlapping_holdings_count / f1.total_holdings,
		f2.total_holdings,
		c.overlapping_holdings_count / f2.total_holdings,
		f1."name",
		f2."name"
		FROM
			fund_holding fh1
			LEFT JOIN fund_holding fh2 ON fh1.holding_id = fh2.holding_id 
			INNER JOIN fund f1 ON f1.id = fh1.fund_id
			INNER JOIN fund f2 ON f2.id = fh2.fund_id
			INNER JOIN CTE c ON c.fund_id = f1.id
		WHERE
			fh1.fund_id = :fund1 AND fh2.fund_id = :fund2
			group by f1.total_holdings, f2.total_holdings, c.overlapping_holdings_count, f1."name", f2."name";
`,
		RawArgs{
			":fund1": fundOne,
			":fund2": fundTwo,
		},
	).Sql()

	rows, err := r.ConnectionPool.Query(ctx, sql, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&resp.TotalOverlappingPercentage,
			&resp.OverlappingHoldingsCount,
			&resp.FundOneHoldingCount,
			&resp.FundOneOverlappingCountPercentage,
			&resp.FundTwoHoldingCount,
			&resp.FundTwoOverlappingCountPercentage,
			&resp.FundOneName,
			&resp.FundTwoName,
		)
		if err != nil {
			return resp, err
		}
		return resp, nil
	}
	return resp, nil
}
func (r *Repository) GetOverlappingHoldings(ctx context.Context, fundOne, fundTwo uuid.UUID) (fund.OverlappingHoldings, error) {
	sql, args := RawStatement(
		`
		WITH OverlappingHoldings AS (
		SELECT
			 fh1.holding_id,
			 COALESCE(
				  LEAST(coalesce(fh1.percentage_of_total, 0), coalesce(fh2.percentage_of_total, 0)),
				  0
			 ) AS min_weight
		FROM
			 fund_holding fh1
			 LEFT JOIN fund_holding fh2 ON fh1.holding_id = fh2.holding_id AND fh2.fund_id = :fund2
		WHERE
			 fh1.fund_id = :fund1
		)
		SELECT
		h.id,
		h."name",
	    oh.min_weight AS weighted_overlap_percentage,
	    fh1.percentage_of_total,
	    fh2.percentage_of_total
		FROM
			 OverlappingHoldings oh
			 left JOIN fund_holding fh1 ON oh.holding_id = fh1.holding_id AND fh1.fund_id = :fund1
			 left JOIN fund_holding fh2 ON oh.holding_id = fh2.holding_id AND fh2.fund_id = :fund2
			 join holding h on oh.holding_id = h.id
			 group by h.id, h."name", oh.min_weight, fh1.percentage_of_total, fh2.percentage_of_total
			 order by oh.min_weight desc
			 limit 20;`,
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

	var overlappingHoldings []fund.OverlappingHolding
	for rows.Next() {

		var (
			holdingId             uuid.UUID
			holdingName           string
			overlappingPercentage float64
			fundOnePercentage     float64
			fundTwoPercentage     float64
		)
		err = rows.Scan(
			&holdingId,
			&holdingName,
			&overlappingPercentage,
			&fundOnePercentage,
			&fundTwoPercentage,
		)
		if err != nil {
			return nil, err
		}
		overlappingHoldings = append(overlappingHoldings, fund.OverlappingHolding{
			HoldingId:             holdingId,
			HoldingName:           holdingName,
			OverlappingPercentage: overlappingPercentage,
			FundOnePercentage:     fundOnePercentage,
			FundTwoPercentage:     fundTwoPercentage,
		})
	}
	return overlappingHoldings, nil
}

//WITH OverlappingHoldings AS (
//SELECT
//fh1.holding_id,
//COALESCE(
//LEAST(coalesce(fh1.percentage_of_total, 0), coalesce(fh2.percentage_of_total, 0)),
//0
//) AS min_weight
//FROM
//fund_holding fh1
//LEFT JOIN fund_holding fh2 ON fh1.holding_id = fh2.holding_id AND fh2.fund_id = 'b72bec51-e1ad-4a0e-9a98-98e943419c1f'
//WHERE
//fh1.fund_id = 'f168c64a-282d-40f0-80ef-74eb95de2bb8'
//)
//SELECT
//h."name",
//COALESCE(
//(SUM(oh.min_weight) * 100) /
//NULLIF(SUM(fh1.percentage_of_total) + SUM(fh2.percentage_of_total) - SUM(oh.min_weight), 0),
//0
//) AS weighted_overlap_percentage
//FROM
//OverlappingHoldings oh
//left JOIN fund_holding fh1 ON oh.holding_id = fh1.holding_id AND fh1.fund_id = 'f168c64a-282d-40f0-80ef-74eb95de2bb8'
//left JOIN fund_holding fh2 ON oh.holding_id = fh2.holding_id AND fh2.fund_id = 'b72bec51-e1ad-4a0e-9a98-98e943419c1f'
//join holding h on oh.holding_id = h.id
//group by h.id, h."name", oh.min_weight
//order by oh.min_weight desc;
//

//Select all funds that have ASML in top 10
//SELECT rank_filter.* FROM (
//SELECT h.ticker,fh.*,
//rank() OVER (
//PARTITION BY fh.fund_id
//ORDER BY percentage_of_total DESC
//)
//FROM fund_holding fh
//INNER JOIN holding h ON h.id = fh.holding_id
//) rank_filter WHERE ticker = 'ASML' AND rank < 10;
