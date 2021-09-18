package handlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	auth_grpc.RegisterAuthServiceServer(s, &Server{JWTTokenSecret: "122"})
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
		require.Nil(t, err)

		defer conn.Close()

		c := auth_grpc.NewAuthServiceClient(conn)

		res, err := c.LoginUser(ctx, &auth_grpc.AuthRequest{})

		require.Nil(t, err)
		assert.NotEqual(t, "", res.GetId())
		assert.NotEqual(t, "", res.GetAvatar())
		assert.NotEqual(t, "", res.GetFullName())
		assert.NotEqual(t, "", res.GetJwtToken())
	})
}
