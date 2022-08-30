package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client
	ttl    time.Duration
}

func NewRedis(c *Config) *Redis {
	return &Redis{
		Client: redis.NewClient(
			&redis.Options{
				Addr:     fmt.Sprintf("%s:%s", c.Host, c.Port),
				Password: c.Password,
				DB:       c.DB,
			},
		),
		ttl: time.Duration(c.TTL) * time.Second,
	}
}

func (r *Redis) TTL() time.Duration {
	return r.ttl
}
