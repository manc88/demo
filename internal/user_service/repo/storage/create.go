package storage

import (
	"context"

	"github.com/manc88/demo/internal/models"
)

func (p *PgStorage) Create(ctx context.Context, newUser *models.User) (int64, error) {
	err := p.pool.QueryRow(ctx, createQuery, newUser.Name, newUser.Email, newUser.Age).Scan(&newUser.UID)
	if err != nil {
		return 0, err
	}
	return newUser.UID, nil
}
