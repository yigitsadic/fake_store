package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yigitsadic/fake_store/payment_provider/database"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type mockBadRepository struct {
}

func (m mockBadRepository) Create(_, _, _, _ string, _ float64) (*database.PaymentIntent, error) {
	return nil, errors.New("I am giving an error because I want")
}

func (m mockBadRepository) FindOne(string) (*database.PaymentIntent, error) {
	panic("implement me")
}

func (m mockBadRepository) MarkAsCompleted(string) error {
	return errors.New("I am giving an error because I want")
}

type mockGoodRepository struct {
	Storage map[string]*database.PaymentIntent
}

func initGoodRepositoryWith(records []database.PaymentIntent) *mockGoodRepository {
	storage := make(map[string]*database.PaymentIntent)

	for _, v := range records {
		storage[v.ID] = &v
	}

	return &mockGoodRepository{Storage: storage}
}

func (m *mockGoodRepository) Create(_, _, _, _ string, _ float64) (*database.PaymentIntent, error) {
	return &database.PaymentIntent{ID: "MOCK_ID"}, nil
}

func (m *mockGoodRepository) FindOne(ID string) (*database.PaymentIntent, error) {
	res, ok := m.Storage[ID]
	if ok {
		return res, nil
	}

	return nil, errors.New("not found on db")
}

func (m *mockGoodRepository) MarkAsCompleted(ID string) error {
	record, ok := m.Storage[ID]
	if ok {
		record.Status = database.PaymentCompleted

		return nil
	}

	return errors.New("record not found")
}

func TestServer_HandleShow(t *testing.T) {
	tmp, err := template.New("index.html").Parse(`{{ .AmountDisplay }} EUR - {{ .ID }}`)
	require.Nil(t, err)

	repo := initGoodRepositoryWith([]database.PaymentIntent{
		{
			ID:          "abcdef",
			Amount:      53,
			ReferenceID: "132132",
			Status:      1,
			CreatedAt:   time.Now().UTC(),
			SuccessURL:  "/lorem",
			FailureURL:  "/ipsum",
			HookURL:     "/nice",
		},
	})

	t.Run("it should return 404 if cannot find in database", func(t *testing.T) {
		r := chi.NewRouter()

		s := Server{
			ShowTemplate:            tmp,
			PaymentIntentRepository: repo,
		}

		r.Get("/payments/{paymentIntentID}", s.HandleShow())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodGet, ts.URL+"/payments/abc", nil)
		require.Nil(t, err)

		client := http.Client{}
		res, err := client.Do(req)

		require.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("it should return 200 if it founds in database", func(t *testing.T) {
		r := chi.NewRouter()

		s := Server{
			ShowTemplate:            tmp,
			PaymentIntentRepository: repo,
			SendHookRequest:         nil,
		}

		r.Get("/payments/{paymentIntentID}", s.HandleShow())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodGet, ts.URL+"/payments/abcdef", nil)
		require.Nil(t, err)

		client := http.Client{}
		res, err := client.Do(req)

		require.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}

func TestServer_HandleCreate(t *testing.T) {
	repo := initGoodRepositoryWith([]database.PaymentIntent{})
	payload := createPaymentIntentRequest{
		Amount:      34,
		ReferenceID: "12312321",
		HookURL:     "/hook",
		SuccessURL:  "/success",
		FailureURL:  "/failure",
	}

	t.Run("it should return 422 for empty body", func(t *testing.T) {
		r := chi.NewRouter()

		s := Server{
			PaymentIntentRepository: repo,
		}

		r.Post("/", s.HandleCreate())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
		require.Nil(t, err)

		client := http.Client{}
		res, err := client.Do(req)

		require.Nil(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	})

	t.Run("it should return 422 if it cannot save into database", func(t *testing.T) {
		r := chi.NewRouter()
		badRepo := mockBadRepository{}
		s := Server{PaymentIntentRepository: badRepo}

		r.Post("/", s.HandleCreate())

		ts := httptest.NewServer(r)
		defer ts.Close()

		b, err := json.Marshal(payload)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPost, ts.URL, bytes.NewBuffer(b))
		require.Nil(t, err)

		client := http.Client{}
		res, err := client.Do(req)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, res.StatusCode)
	})

	t.Run("it should return 200 if it is successful", func(t *testing.T) {
		r := chi.NewRouter()
		s := Server{PaymentIntentRepository: repo, BaseURL: "https://google.com"}

		r.Post("/", s.HandleCreate())

		ts := httptest.NewServer(r)
		defer ts.Close()

		b, err := json.Marshal(payload)
		require.Nil(t, err)

		req, err := http.NewRequest(http.MethodPost, ts.URL, bytes.NewBuffer(b))
		require.Nil(t, err)

		client := http.Client{}
		res, err := client.Do(req)

		require.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		data, err := io.ReadAll(res.Body)
		defer res.Body.Close()

		require.Nil(t, err)

		var response createPaymentIntentResponse

		require.Nil(t, json.Unmarshal(data, &response))
		assert.Equal(t, "MOCK_ID", response.ID)
		assert.Equal(t, "https://google.com/payments/"+response.ID, response.PaymentURL)
	})
}

func TestServer_HandleComplete(t *testing.T) {
	repo := initGoodRepositoryWith([]database.PaymentIntent{
		{
			ID:          "abcdef",
			Amount:      53,
			ReferenceID: "132132",
			Status:      database.PaymentInitialized,
			CreatedAt:   time.Now().UTC(),
			SuccessURL:  "/success",
			FailureURL:  "/failure",
			HookURL:     "/",
		},
		{
			ID:          "completed_example",
			Amount:      53,
			ReferenceID: "132132",
			Status:      database.PaymentCompleted,
			CreatedAt:   time.Now().UTC(),
			SuccessURL:  "/success",
			FailureURL:  "/failure",
			HookURL:     "/",
		},
	})

	t.Run("it should respond with not found if it's absent", func(t *testing.T) {
		r := chi.NewRouter()
		s := Server{PaymentIntentRepository: repo}
		r.Post("/{paymentIntentID}", s.HandleComplete())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodPost, ts.URL+"/abc", nil)
		require.Nil(t, err)

		client := http.Client{}
		res, err := client.Do(req)

		require.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})

	t.Run("it should redirect to success even it's completed before", func(t *testing.T) {
		var callCounter int

		r := chi.NewRouter()
		s := Server{
			PaymentIntentRepository: repo,
			SendHookRequest: func(_ string, _ database.PaymentHookMessage) error {
				callCounter++
				return nil
			},
		}
		r.Post("/{paymentIntentID}", s.HandleComplete())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodPost, ts.URL+"/completed_example", nil)
		require.Nil(t, err)

		req.Header.Set("Accept", "")

		client := http.Client{}
		res, err := client.Do(req)

		require.Nil(t, err)
		assert.Equal(t, 0, callCounter)
		assert.Equal(t, "/success", res.Request.URL.Path)
	})

	t.Run("it should redirect to failure page if it cannot push hook", func(t *testing.T) {
		var callCounter int

		rr := &mockGoodRepository{Storage: map[string]*database.PaymentIntent{
			"abcdef": {
				ID:         "abcdef",
				Status:     database.PaymentInitialized,
				HookURL:    "/hook",
				FailureURL: "/failure",
				SuccessURL: "/success",
			},
		}}

		r := chi.NewRouter()
		s := Server{
			PaymentIntentRepository: rr,
			SendHookRequest: func(_ string, mes database.PaymentHookMessage) error {
				callCounter++

				return errors.New("something happened")
			},
		}
		r.Post("/{paymentIntentID}", s.HandleComplete())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodPost, ts.URL+"/abcdef", nil)
		require.Nil(t, err)

		req.Header.Set("Accept", "")

		client := http.Client{}
		res, err := client.Do(req)

		require.Nil(t, err)
		assert.True(t, callCounter > 0)
		assert.Equal(t, "/failure", res.Request.URL.Path)
	})

	t.Run("it should redirect to success page if everything goes smoothly", func(t *testing.T) {
		var callCounter int
		var targetURL string

		rr := &mockGoodRepository{Storage: map[string]*database.PaymentIntent{
			"abcdef": {
				ID:         "abcdef",
				Status:     database.PaymentInitialized,
				HookURL:    "/hook",
				FailureURL: "/failure",
				SuccessURL: "/success",
			},
		}}

		r := chi.NewRouter()
		s := Server{
			PaymentIntentRepository: rr,
			SendHookRequest: func(target string, mes database.PaymentHookMessage) error {
				targetURL = target
				callCounter++

				return nil
			},
		}
		r.Post("/{paymentIntentID}", s.HandleComplete())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodPost, ts.URL+"/abcdef", nil)
		require.Nil(t, err)

		req.Header.Set("Accept", "")

		client := http.Client{}
		res, err := client.Do(req)

		require.Nil(t, err)
		assert.Equal(t, 1, callCounter)
		assert.Equal(t, "/hook", targetURL)
		assert.Equal(t, "/success", res.Request.URL.Path)
	})
}
