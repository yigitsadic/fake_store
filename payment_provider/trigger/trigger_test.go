package trigger

import (
	"github.com/go-chi/chi/v5"
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

	r, err := buildRequest("http://google.com/lorems_ipsums", payload)
	if err != nil {
		t.Errorf("unexpected to get and error while building request but got=%s", err)
	}

	if r.URL.String() != "http://google.com/lorems_ipsums" {
		t.Errorf("does not satisfy expected url. got=%s", r.URL.String())
	}

	if r.Method != http.MethodPost {
		t.Errorf("expected to create POST request but got=%s", r.Method)
	}

	if r.Header.Get("Content-Type") != "application/json" {
		t.Errorf("expected to see application/json as content type header")
	}

	if r.Header.Get("Accept") != "application/json" {
		t.Errorf("expected to see application/json as accept header")
	}
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
		if err != nil {
			t.Fatalf("unexpected to get an error at this step but got=%s", err)
		}

		if err = makeRequest(req); err == nil {
			t.Error("expected to get an error but got nothing")
		}
	})

	t.Run("it should return nil if server responds with status ok", func(t *testing.T) {
		r := chi.NewRouter()
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := buildRequest(ts.URL, payload)
		if err != nil {
			t.Fatalf("unexpected to get an error at this step but got=%s", err)
		}

		if err = makeRequest(req); err != nil {
			t.Errorf("unexpected to get an error but got=%s", err)
		}
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

		if err := SendHookRequest(ts.URL, payload); err == nil {
			t.Error("expected to get an error but got nothing")
		}
	})

	t.Run("it should return nil if it can react to hookURL", func(t *testing.T) {
		r := chi.NewRouter()
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		ts := httptest.NewServer(r)
		defer ts.Close()

		if err := SendHookRequest(ts.URL, payload); err != nil {
			t.Errorf("unexpected to get an error but got=%s", err)
		}
	})
}
