package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx"
)

func (p *PgStorage) Delete(ctx context.Context, uid int64) error {
	err := p.pool.QueryRow(ctx, deleteQuery, uid).Scan(&uid)
	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("not found")
		}
		return err
	}
	return nil
}
