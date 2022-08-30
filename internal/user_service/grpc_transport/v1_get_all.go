package grpctransport

import (
	"context"

	"github.com/manc88/demo/internal/user_service/grpc_transport/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *Transport) GetAllV1(ctx context.Context, _ *proto.GetAllV1Request) (*proto.GetAllV1Response, error) {
	ctx, cancel := context.WithTimeout(ctx, t.defaultTimeOut)
	defer cancel()

	users, err := t.service.GetAllV1(ctx)
	if err != nil {
		return &proto.GetAllV1Response{}, status.Error(codes.Internal, err.Error())
	}

	response := &proto.GetAllV1Response{Users: make([]*proto.User, 0, len(users))}
	for _, u := range users {
		response.Users = append(response.Users, ProtoFromUserP(u))
	}

	return response, nil
}
