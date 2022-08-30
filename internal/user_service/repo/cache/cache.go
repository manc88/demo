package cache

import (
	"context"
	"encoding/json"

	"github.com/manc88/demo/internal/models"
	userservice "github.com/manc88/demo/internal/user_service"
	"github.com/manc88/demo/pkg/redis"
)

var _ userservice.ICache = (*UserCache)(nil)

type usersList []*models.User

func (u usersList) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *usersList) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

type UserCache struct {
	redis       *redis.Redis
	userListKey string
}

func NewUserCache(r *redis.Redis, key string) *UserCache {
	return &UserCache{
		redis:       r,
		userListKey: key,
	}
}

func (uc *UserCache) Store(ctx context.Context, users []*models.User) error {
	err := uc.redis.Client.Set(ctx, uc.userListKey, (usersList)(users), uc.redis.TTL()).Err()
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserCache) Load(ctx context.Context) ([]*models.User, error) {
	users := make(usersList, 0)
	res := uc.redis.Client.Get(ctx, uc.userListKey)
	if res.Err() != nil {
		return nil, res.Err()
	}

	err := res.Scan(&users)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserCache) Reset(ctx context.Context) error {
	uc.redis.Client.Del(ctx, uc.userListKey)
	return nil
}
