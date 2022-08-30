package pgx

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPGX(ctx context.Context, pgc *Config) (*pgxpool.Pool, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cfg, err := pgxpool.ParseConfig(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pgc.Host, pgc.Port, pgc.User, pgc.Password, pgc.Database, pgc.Sslmode))
	if err != nil {
		return nil, err
	}

	cfg.MaxConnIdleTime = time.Duration(pgc.MaxConnIdleTimeSec) * time.Second
	cfg.MaxConnLifetime = time.Duration(pgc.MaxConnLifetimeSec) * time.Second
	cfg.MinConns = pgc.MinConns
	cfg.MaxConns = pgc.MaxConns

	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	if err := conn.Conn().Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
