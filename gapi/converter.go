package gapi

import (
	db "github.com/PichayuthK/go-simple-bank/db/sqlc"
	"github.com/PichayuthK/go-simple-bank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func converterUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
