package main

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/yigitsadic/fake_store/auth/client/client"
	"google.golang.org/grpc"
	"log"
	"net"
)

const DiceBearUrl = "https://avatars.dicebear.com/api/human/%s.svg"

type Server struct {
	client.UnimplementedAuthServiceServer
}

func (s *Server) LoginUser(context.Context, *client.AuthRequest) (*client.UserResponse, error) {
	resp := client.UserResponse{
		Id:       faker.UUIDDigit(),
		Avatar:   fmt.Sprintf(DiceBearUrl, faker.UUIDDigit()),
		FullName: faker.FirstName() + " " + faker.LastName(),
	}
	resp.JwtToken = GenerateJWTToken(resp.Id, resp.Avatar, resp.FullName)

	return &resp, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	s := Server{}

	client.RegisterAuthServiceServer(grpcServer, &s)

	log.Println("Started to serve auth grpc")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve due to %s\n", err)
	}
}
