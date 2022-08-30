package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	userservice "github.com/manc88/demo/internal/user_service"
)

var _ userservice.IStorage = (*PgStorage)(nil)

type PgStorage struct {
	ctx    context.Context
	logger *log.Logger
	pool   *pgxpool.Pool
}

func NewPgStorage(ctx context.Context, p *pgxpool.Pool) *PgStorage {
	return &PgStorage{
		ctx:  ctx,
		pool: p,
	}
}

func (p *PgStorage) SetLogger(l *log.Logger) {
	p.logger = l
}

func (p *PgStorage) Close() error {
	//why u dont need to close pgx pool
	//https://github.com/jackc/pgx/issues/802#issuecomment-668713840
	p.pool.Close()
	p.logger.Println("PGX storage down")
	return nil
}
