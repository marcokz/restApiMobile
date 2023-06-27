package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"log"
)

func NewPostgres(ctx context.Context, postgresURL string) (*pgxpool.Pool, error) { //универсальная для постгрес соединения
	poolConfig, err := pgxpool.ParseConfig(postgresURL)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrapf(err, "postgresclient.NewPostgres.pgxpool.ParseConfig, failed to parse postgres url")
	}
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrapf(err, "postgresclient.NewPostgres.pgxpool.Connect, failed to connect to postgres, postgresURL: %s", postgresURL)
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.Wrapf(err, "postgresclient.NewPostgres.pool.Ping, failed to ping postgres")
	}
	return pool, nil
}
