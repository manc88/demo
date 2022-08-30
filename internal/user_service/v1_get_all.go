package userservice

import (
	"context"

	"github.com/manc88/demo/internal/models"
)

func (u *UserService) GetAllV1(ctx context.Context) ([]*models.User, error) {
	users, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
