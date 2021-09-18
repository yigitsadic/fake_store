package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	client := http.Client{}
	var readContent string

	t.Run("it should write to context if it's present", func(t *testing.T) {
		readContent = ""
		expectedID := "123213-3322-1223-3445"

		c := &jwt.StandardClaims{
			Id:        expectedID,
			ExpiresAt: time.Now().Add(time.Hour * 2).UTC().Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
		ss, err := token.SignedString(secret)

		assert.Nil(t, err)

		ts := buildTestServer(t, &readContent)
		defer ts.Close()

		req := buildRequestWithAuthorization(t, ts.URL, ss)

		_, err = client.Do(req)

		assert.Nil(t, err)
		assert.Equal(t, expectedID, readContent)
	})

	t.Run("it should write empty string to context with non-related token", func(t *testing.T) {
		readContent = ""
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.z44tlyeOKLLrGMdctidcC7kZ6i8jQ4LWv1UogjXSnlI"

		ts := buildTestServer(t, &readContent)
		defer ts.Close()

		req := buildRequestWithAuthorization(t, ts.URL, token)

		_, err := client.Do(req)

		assert.Nil(t, err)
		assert.Equal(t, "", readContent)
	})

	t.Run("it should write empty string if not present", func(t *testing.T) {
		readContent = ""

		ts := buildTestServer(t, &readContent)
		defer ts.Close()

		req := buildRequestWithAuthorization(t, ts.URL, "")

		_, err := client.Do(req)

		assert.Nil(t, err)
		assert.Equal(t, "", readContent)
	})
}

func buildTestServer(t *testing.T, userID *string) *httptest.Server {
	t.Helper()

	r := chi.NewRouter()
	r.Use(Auth)
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		*userID = request.Context().Value(userIDCtxKey).(string)
	})

	return httptest.NewServer(r)
}

func buildRequestWithAuthorization(t *testing.T, url, token string) *http.Request {
	t.Helper()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	require.Nil(t, err)

	req.Header.Set("Authorization", "Bearer "+token)

	return req
}
