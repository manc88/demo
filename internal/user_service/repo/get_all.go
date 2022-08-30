package repo

import (
	"context"

	"github.com/manc88/demo/internal/models"
)

func (u *UserRepository) GetAll(ctx context.Context) ([]*models.User, error) {

	ctx, cancel := context.WithTimeout(ctx, u.defaultTimeOut)
	defer cancel()

	if users, err := u.cache.Load(ctx); err == nil {
		return users, nil
	}

	u.Lock()
	defer u.Unlock()

	users, err := u.storage.GetAll(ctx)
	if err != nil {
		return users, err
	}

	err = u.cache.Store(ctx, users)
	if err != nil {
		u.logger.Println(err)
	}

	return users, nil

}
