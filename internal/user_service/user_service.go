package userservice

import (
	"context"
	"log"
)

type UserService struct {
	ctx           context.Context
	repo          IRepository
	logger        *log.Logger
	messageBroker IMessageBroker
	config        *Config
}

func New(ctx context.Context, c *Config) *UserService {
	return &UserService{
		ctx:    ctx,
		config: c,
	}
}

func (u *UserService) SetRepository(r IRepository) {
	u.repo = r
}

func (u *UserService) SetMessageBroker(m IMessageBroker) {
	u.messageBroker = m
}

func (u *UserService) SetLogger(l *log.Logger) {
	u.logger = l
}

func (u *UserService) Shutdown() {
	u.messageBroker.Close()
}
