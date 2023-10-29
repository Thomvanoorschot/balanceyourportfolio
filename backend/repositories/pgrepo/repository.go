package pgrepo

import (
	"context"
	"errors"
	"log"
	"sync"

	"etfinsight/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	ConnectionPool *pgxpool.Pool
}

var repository *Repository
var once sync.Once

func NewRepository(config *config.Config) *Repository {
	once.Do(func() {
		pool, err := pgxpool.New(context.Background(), config.DbConnectionString)
		if err != nil {
			log.Fatal(err)
		}
		repository = &Repository{ConnectionPool: pool}
	})

	return repository
}

func (r *Repository) NewTransaction(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.ConnectionPool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadUncommitted})
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (r *Repository) RollBack(tx pgx.Tx, ctx context.Context) {
	errTxRollBack := tx.Rollback(ctx)
	if errTxRollBack != nil && !errors.Is(errTxRollBack, pgx.ErrTxClosed) {
		log.Fatal(errTxRollBack)
	}
}
