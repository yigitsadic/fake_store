package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
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
	if err != nil {
		t.Fatalf("unexpected to get an while parsing tempalte but got=%s", err)
	}

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

		s := server{
			ShowTemplate:            tmp,
			PaymentIntentRepository: repo,
		}

		r.Get("/payments/{paymentIntentID}", s.HandleShow())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodGet, ts.URL+"/payments/abc", nil)
		if err != nil {
			t.Fatalf("unexpected to get an error while building request but got=%s", err)
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected to get an error while sending GET request but got=%s", err)
		}

		if res.StatusCode != http.StatusNotFound {
			t.Errorf("expected to get status not found but got=%d", res.StatusCode)
		}
	})

	t.Run("it should return 200 if it founds in database", func(t *testing.T) {
		r := chi.NewRouter()

		s := server{
			ShowTemplate:            tmp,
			PaymentIntentRepository: repo,
			SendHookRequest:         nil,
		}

		r.Get("/payments/{paymentIntentID}", s.HandleShow())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodGet, ts.URL+"/payments/abcdef", nil)
		if err != nil {
			t.Fatalf("unexpected to get an error while building request but got=%s", err)
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected to get an error while sending GET request but got=%s", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected to get status ok response but got=%d", res.StatusCode)
		}
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

		s := server{
			PaymentIntentRepository: repo,
		}

		r.Post("/", s.HandleCreate())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
		if err != nil {
			t.Fatalf("unexpected to get an error while building request but got=%s", err)
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected to get an error while sending create intent request but got=%s", err)
		}

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected status code was=%d but got=%d", http.StatusUnprocessableEntity, res.StatusCode)
		}
	})

	t.Run("it should return 422 if it cannot save into database", func(t *testing.T) {
		r := chi.NewRouter()
		badRepo := mockBadRepository{}
		s := server{PaymentIntentRepository: badRepo}

		r.Post("/", s.HandleCreate())

		ts := httptest.NewServer(r)
		defer ts.Close()

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal("unexpected to get an error while marshaling struct to json")
		}

		req, err := http.NewRequest(http.MethodPost, ts.URL, bytes.NewBuffer(b))
		if err != nil {
			t.Fatalf("unexpected to get an error while building request but got=%s", err)
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected to get an error while sending create intent request but got=%s", err)
		}

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Errorf("expected status code was=%d but got=%d", http.StatusUnprocessableEntity, res.StatusCode)
		}
	})

	t.Run("it should return 200 if it is successful", func(t *testing.T) {
		r := chi.NewRouter()
		s := server{PaymentIntentRepository: repo, BaseURL: "https://google.com"}

		r.Post("/", s.HandleCreate())

		ts := httptest.NewServer(r)
		defer ts.Close()

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal("unexpected to get an error while marshaling struct to json")
		}

		req, err := http.NewRequest(http.MethodPost, ts.URL, bytes.NewBuffer(b))
		if err != nil {
			t.Fatalf("unexpected to get an error while building request but got=%s", err)
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected to get an error while sending create intent request but got=%s", err)
		}

		if res.StatusCode != http.StatusOK {
			t.Errorf("expected status code was=%d but got=%d", http.StatusUnprocessableEntity, res.StatusCode)
		}

		data, err := io.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			t.Fatalf("error occurred during reading request body. err=%s", err)
		}

		var response createPaymentIntentResponse

		if err = json.Unmarshal(data, &response); err != nil {
			t.Fatalf("error occurred during unmarshaling json into struct. err=%s", err)
		}

		if response.ID != "MOCK_ID" {
			t.Errorf("expected to see MOCK_ID id at response but got=%s", response.ID)
		}

		expectedPaymentURL := "https://google.com/payments/" + response.ID

		if response.PaymentURL != expectedPaymentURL {
			t.Errorf("expected to get correct payment URL=%q but got=%q", expectedPaymentURL, response.PaymentURL)
		}
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
		s := server{PaymentIntentRepository: repo}
		r.Post("/{paymentIntentID}", s.HandleComplete())

		ts := httptest.NewServer(r)
		defer ts.Close()

		req, err := http.NewRequest(http.MethodPost, ts.URL+"/abc", nil)
		if err != nil {
			t.Fatalf("unexpected to get an error while building request but got=%s", err)
		}

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected to get an error while completing payment but got=%s", err)
		}

		if res.StatusCode != http.StatusNotFound {
			t.Errorf("expected to get 404 response but got=%d", res.StatusCode)
		}
	})

	t.Run("it should redirect to success even it's completed before", func(t *testing.T) {
		var callCounter int

		r := chi.NewRouter()
		s := server{
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
		if err != nil {
			t.Fatalf("unexpected to get an error while building request but got=%s", err)
		}

		req.Header.Set("Accept", "")

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected to get an error while completing payment but got=%s", err)
		}

		if callCounter != 0 {
			t.Errorf("unexpected to see call counter incremented")
		}

		if res.Request.URL.Path != "/success" {
			t.Errorf("expected to redirect %q but got=%q", "/success", res.Request.URL.Path)
		}
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
		s := server{
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
		if err != nil {
			t.Fatalf("unexpected to get an error while building request but got=%s", err)
		}

		req.Header.Set("Accept", "")

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected to get an error while completing payment but got=%s", err)
		}

		if callCounter < 1 {
			t.Errorf("expected to see call counter incremented")
		}

		if res.Request.URL.Path != "/failure" {
			t.Errorf("expected to redirect %q but got=%q", "/failure", res.Request.URL.Path)
		}
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
		s := server{
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
		if err != nil {
			t.Fatalf("unexpected to get an error while building request but got=%s", err)
		}

		req.Header.Set("Accept", "")

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("unexpected to get an error while completing payment but got=%s", err)
		}

		if callCounter != 1 {
			t.Errorf("expected to see call counter incremented")
		}

		if targetURL != "/hook" {
			t.Errorf("expected hook url not satisfied. expected=%q got=%q", "/hook", targetURL)
		}

		if res.Request.URL.Path != "/success" {
			t.Errorf("expected to redirect %q but got=%q", "/success", res.Request.URL.Path)
		}
	})
}
