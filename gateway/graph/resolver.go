//go:generate go run github.com/99designs/gqlgen

package graph

import "github.com/yigitsadic/fake_store/auth/client/client"

type Resolver struct {
	AuthClient client.AuthServiceClient
}
