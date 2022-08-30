package userservice

import (
	"context"
	"io"

	"github.com/manc88/demo/internal/models"
)

type ICreateCommand interface {
	Create(context.Context, *models.User) (int64, error)
}

type IDeleteCommand interface {
	Delete(context.Context, int64) error
}

type IGetAllCommand interface {
	GetAll(ctx context.Context) ([]*models.User, error)
}

type ICommands interface {
	ICreateCommand
	IDeleteCommand
	IGetAllCommand
}

type IRepository interface {
	io.Closer
	ICommands
}

type IStorage interface {
	io.Closer
	ICommands
}

type ICache interface {
	Store(context.Context, []*models.User) error
	Load(context.Context) ([]*models.User, error)
	Reset(context.Context) error
}

type IMessageBroker interface {
	io.Closer
	Write(dest string, data []byte) error
}

type ITransport interface {
	RegisterUserService(ICommands)
}
