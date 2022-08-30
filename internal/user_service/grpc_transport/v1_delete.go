package grpctransport

import (
	"context"

	"github.com/manc88/demo/internal/user_service/grpc_transport/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *Transport) DeleteV1(ctx context.Context, r *proto.DeleteV1Request) (*proto.DeleteV1Response, error) {
	ctx, cancel := context.WithTimeout(ctx, t.defaultTimeOut)
	defer cancel()
	uid := r.GetUid()
	err := t.service.DeleteV1(ctx, uid)
	if err != nil {
		return &proto.DeleteV1Response{}, status.Error(codes.Internal, err.Error())
	}

	return &proto.DeleteV1Response{}, nil
}
