package main

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/yigitsadic/fake_store/auth/auth_grpc/auth_grpc"
)

const diceBearURL = "https://avatars.dicebear.com/api/human/%s.svg"

type server struct {
	auth_grpc.UnimplementedAuthServiceServer
}

func (s *server) LoginUser(context.Context, *auth_grpc.AuthRequest) (*auth_grpc.UserResponse, error) {
	resp := auth_grpc.UserResponse{
		Id:       faker.UUIDDigit(),
		Avatar:   fmt.Sprintf(diceBearURL, faker.UUIDDigit()),
		FullName: faker.FirstName() + " " + faker.LastName(),
	}
	resp.JwtToken = generateJWTToken(resp.Id, resp.Avatar, resp.FullName)

	return &resp, nil
}
