package main

import (
	"context"
	"github.com/yigitsadic/fake_store/auth/auth_grpc/auth_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	auth_grpc.RegisterAuthServiceServer(s, &server{JWTTokenSecret: "122"})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestServer_LoginUser(t *testing.T) {
	t.Run("it creates random user login credentials", func(t *testing.T) {
		ctx := context.Background()

		conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
		if err != nil {
			t.Fatalf("unexpected to get connection error. Err=%s", err)
		}

		defer conn.Close()

		c := auth_grpc.NewAuthServiceClient(conn)

		res, err := c.LoginUser(ctx, &auth_grpc.AuthRequest{})
		if err != nil {
			t.Fatalf("unexpected to get an error while creating user credentials. Err=%s", err)
		}

		if res.GetId() == "" {
			t.Error("expected random id but got nothing")
		}

		if res.GetAvatar() == "" {
			t.Error("expected random avatar but got nothing")
		}

		if res.GetFullName() == "" {
			t.Error("expected random full name but got nothing")
		}

		if res.GetJwtToken() == "" {
			t.Error("expected random jwt token but got nothing")
		}
	})
}
