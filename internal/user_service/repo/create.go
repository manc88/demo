package repo

import (
	"context"

	"github.com/manc88/demo/internal/models"
)

func (u *UserRepository) Create(ctx context.Context, newUser *models.User) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, u.defaultTimeOut)
	defer cancel()

	newUID, err := u.storage.Create(ctx, newUser)
	if err != nil {
		return 0, err
	}

	err = u.cache.Reset(ctx)
	if err != nil {
		u.logger.Println(err)
	}

	return newUID, nil
}
