package handlers

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/yigitsadic/fake_store/auth/auth_grpc/auth_grpc"
	"github.com/yigitsadic/fake_store/auth/utilities"
)

const diceBearURL = "https://avatars.dicebear.com/api/human/%s.svg"

type Server struct {
	auth_grpc.UnimplementedAuthServiceServer

	JWTTokenSecret string
}

func (s *Server) LoginUser(context.Context, *auth_grpc.AuthRequest) (*auth_grpc.UserResponse, error) {
	resp := auth_grpc.UserResponse{
		Id:       faker.UUIDDigit(),
		Avatar:   fmt.Sprintf(diceBearURL, faker.UUIDDigit()),
		FullName: faker.FirstName() + " " + faker.LastName(),
	}
	resp.JwtToken = utilities.GenerateJWTToken(s.JWTTokenSecret, resp.Id, resp.Avatar, resp.FullName)

	return &resp, nil
}
