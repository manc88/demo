package grpctransport

import (
	"github.com/manc88/demo/internal/models"
	"github.com/manc88/demo/internal/user_service/grpc_transport/proto"
)

func UserFromProto(p *proto.User) models.User {
	return models.User{
		UID:     p.GetUid(),
		Name:    p.GetName(),
		Email:   p.GetEmail(),
		Age:     p.GetAge(),
		Deleted: p.GetDeleted(),
	}
}

func UserFromProtoP(p *proto.User) *models.User {
	tmp := UserFromProto(p)
	return &tmp
}

func ProtoFromUser(u *models.User) proto.User {
	return proto.User{
		Uid:     u.UID,
		Name:    u.Name,
		Email:   u.Email,
		Age:     u.Age,
		Deleted: u.Deleted,
	}
}

func ProtoFromUserP(u *models.User) *proto.User {
	tmp := ProtoFromUser(u)
	return &tmp
}
