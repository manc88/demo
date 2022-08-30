package grpctransport

import (
	"context"

	"github.com/manc88/demo/internal/user_service/grpc_transport/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *Transport) CreateV1(ctx context.Context, r *proto.CreateV1Request) (*proto.CreateV1Response, error) {
	ctx, cancel := context.WithTimeout(ctx, u.defaultTimeOut)
	defer cancel()
	user := UserFromProto(r.GetUser())
	uid, err := u.service.CreateV1(ctx, &user)
	if err != nil {
		return &proto.CreateV1Response{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreateV1Response{Uid: uid}, nil
}
