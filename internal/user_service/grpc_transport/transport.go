package grpctransport

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	userservice "github.com/manc88/demo/internal/user_service"
	"github.com/manc88/demo/internal/user_service/grpc_transport/proto"
	"google.golang.org/grpc"
)

type Transport struct {
	ctx context.Context
	proto.UnimplementedUserServiceServer
	service        *userservice.UserService
	host           string
	port           string
	defaultTimeOut time.Duration
}

func New(ctx context.Context, c *Config) *Transport {
	return &Transport{
		ctx:            ctx,
		host:           c.Host,
		port:           c.Port,
		defaultTimeOut: 3 * time.Second,
	}
}

func (u *Transport) SetService(s *userservice.UserService) {
	u.service = s
}

func (u *Transport) Serve() error {
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", u.host, u.port))
		if err != nil {
			return
		}
		s := grpc.NewServer()
		proto.RegisterUserServiceServer(s, u)

		if err := s.Serve(lis); err != nil {
			return
		}
	}()
	<-u.ctx.Done()
	u.service.Shutdown()
	return errors.New("user service shutdown")
}
