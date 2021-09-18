package trigger

import (
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/fake_store/payment_provider/database"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_buildRequest(t *testing.T) {
	payload := database.PaymentHookMessage{
		ID:          "111",
		Amount:      33,
		ReferenceID: "4343",
		Status:      2,
		CreatedAt:   time.Now().UTC(),
	}

	r, err := buildRequest("https://google.com/lorems_ipsums", payload)

	assert.Nilf(t, err, "unexpected to get and error while building request but got=%s", err)
	assert.Equal(t, "https://google.com/lorems_ipsums", r.URL.String())
	assert.Equal(t, http.MethodPost, r.Method)
	assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
	assert.Equal(t, "application/json", r.Header.Get("Accept"))
}

func Test_makeRequest(t *testing.T) {
	payload := database.PaymentHookMessage{
		ID:          "111",
		Amount:      33,
		ReferenceID: "4343",
		Status:      2,
		CreatedAt:   time.Now().UTC(),
	}

	t.Run("it should return error unless server responds with status ok", func(t *testing.T) {
		r := chi.NewRouter()
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnprocessableEntity)
		})
		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := buildRequest(ts.URL, payload)

		require.Nil(t, err)
		assert.NotNil(t, makeRequest(req))
	})

	t.Run("it should return nil if server responds with status ok", func(t *testing.T) {
		r := chi.NewRouter()
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := buildRequest(ts.URL, payload)

		require.Nil(t, err)
		assert.Nil(t, makeRequest(req))
	})
}

func TestSendHookRequest(t *testing.T) {
	payload := database.PaymentHookMessage{
		ID:          "111",
		Amount:      33,
		ReferenceID: "4343",
		Status:      2,
		CreatedAt:   time.Now().UTC(),
	}

	t.Run("it should return error if cannot reach to hookURL", func(t *testing.T) {
		r := chi.NewRouter()
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})
		ts := httptest.NewServer(r)
		defer ts.Close()

		assert.NotNil(t, SendHookRequest(ts.URL, payload))
	})

	t.Run("it should return nil if it can react to hookURL", func(t *testing.T) {
		r := chi.NewRouter()
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		ts := httptest.NewServer(r)
		defer ts.Close()

		assert.Nil(t, SendHookRequest(ts.URL, payload))
	})
}
