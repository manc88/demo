package userservice

import (
	"context"
)

func (u *UserService) DeleteV1(ctx context.Context, uid int64) error {
	err := u.repo.Delete(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}
