package userservice

import (
	"context"
	"encoding/json"
	"time"

	"github.com/manc88/demo/internal/models"
)

func (u *UserService) CreateV1(ctx context.Context, user *models.User) (int64, error) {
	uid, err := u.repo.Create(ctx, user)
	if err != nil {
		return -1, err
	}

	b, err := json.Marshal(
		&UserCreateInfo{
			UserId: uid,
			Name:   user.Name,
			Email:  user.Email,
			Age:    user.Age,
			Time:   time.Now().Unix(),
		},
	)
	if err != nil {
		u.logger.Println(err)
		return uid, nil
	}

	err = u.messageBroker.Write(u.config.UserCreationTopic, b)
	if err != nil {
		u.logger.Println(err)
		return uid, nil
	}
	return uid, nil
}
