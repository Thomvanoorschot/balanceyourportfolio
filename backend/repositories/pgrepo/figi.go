package pgrepo

import (
	"context"
	"etfinsight/generated/jet_gen/postgres/public/model"
	. "etfinsight/generated/jet_gen/postgres/public/table"
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) UpsertFigiMapping(ctx context.Context, s []model.FigiMapping, tx pgx.Tx) error {
	sql, args := FigiMapping.
		INSERT(FigiMapping.AllColumns).
		MODELS(s).
		ON_CONFLICT(FigiMapping.Figi).
		DO_UPDATE(
			SET(
				FigiMapping.Sedol.SET(FigiMapping.EXCLUDED.Sedol),
				FigiMapping.Cusip.SET(FigiMapping.EXCLUDED.Cusip),
				FigiMapping.Isin.SET(FigiMapping.EXCLUDED.Isin),
				FigiMapping.Name.SET(FigiMapping.EXCLUDED.Name),
				FigiMapping.Ticker.SET(FigiMapping.EXCLUDED.Ticker),
			),
		).
		Sql()

	_, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (r *Repository) UpsertFigiISINMapping(ctx context.Context, s []model.FigiMapping, tx pgx.Tx) error {
	sql, args := FigiMapping.
		INSERT(FigiMapping.AllColumns).
		MODELS(s).
		ON_CONFLICT(FigiMapping.Figi).
		DO_UPDATE(
			SET(
				FigiMapping.Isin.SET(FigiMapping.EXCLUDED.Isin),
			),
		).
		Sql()

	_, err := tx.Exec(ctx, sql, args...)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
func (r *Repository) GetFigiMappings(ctx context.Context) (map[string]model.FigiMapping, map[string]model.FigiMapping, error) {
	sql, args := SELECT(FigiMapping.AllColumns).
		FROM(FigiMapping).
		Sql()
	rows, err := r.ConnectionPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	cusipMapping := map[string]model.FigiMapping{}
	sedolMapping := map[string]model.FigiMapping{}
	for rows.Next() {

		var (
			figi   string
			ticker string
			name   string
			isin   *string
			sedol  *string
			cusip  *string
		)
		err = rows.Scan(
			&figi,
			&ticker,
			&name,
			&isin,
			&sedol,
			&cusip,
		)
		if err != nil {
			return nil, nil, err
		}
		figiMapping := model.FigiMapping{
			Figi:   figi,
			Ticker: &ticker,
		}
		if cusip != nil {
			figiMapping.Cusip = cusip
			cusipMapping[*cusip] = figiMapping
		}
		if sedol != nil {
			figiMapping.Sedol = sedol
			sedolMapping[*sedol] = figiMapping
		}
	}
	return cusipMapping, sedolMapping, nil
}
