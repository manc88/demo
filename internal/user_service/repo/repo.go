package repo

import (
	"log"
	"sync"
	"time"

	userservice "github.com/manc88/demo/internal/user_service"
)

var _ userservice.IRepository = (*UserRepository)(nil)

type UserRepository struct {
	sync.Mutex
	cache          userservice.ICache
	storage        userservice.IStorage
	logger         *log.Logger
	defaultTimeOut time.Duration
}

func NewUserRepository(timeout time.Duration) *UserRepository {
	return &UserRepository{
		defaultTimeOut: timeout,
	}
}

func (u *UserRepository) SetStorage(s userservice.IStorage) {
	u.storage = s
}

func (u *UserRepository) SetCache(c userservice.ICache) {
	u.cache = c
}

func (u *UserRepository) SetLogger(l *log.Logger) {
	u.logger = l
}

func (u *UserRepository) Close() error {
	return u.storage.Close()
}
