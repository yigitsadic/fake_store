package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/yigitsadic/fake_store/payment_provider/database"
	"html/template"
	"net/http"
)

const (
	contentType = "application/json"
)

type Server struct {
	BaseURL                 string
	ShowTemplate            *template.Template
	PaymentIntentRepository database.Repository
	SendHookRequest         func(hookURL string, message database.PaymentHookMessage) error
}

type createPaymentIntentRequest struct {
	Amount      float64 `json:"amount"`
	ReferenceID string  `json:"reference_id"`

	HookURL    string `json:"hook_url"`
	SuccessURL string `json:"success_url"`
	FailureURL string `json:"failure_url"`
}

type createPaymentIntentResponse struct {
	ID         string `json:"id"`
	PaymentURL string `json:"payment_url"`
}

func (s *Server) HandleCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b := createPaymentIntentRequest{}
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)

			return
		}

		intent, err := s.PaymentIntentRepository.Create(
			b.ReferenceID,
			b.HookURL,
			b.SuccessURL,
			b.FailureURL,
			b.Amount,
		)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		w.Header().Set("Content-Type", contentType)

		response := createPaymentIntentResponse{
			ID:         intent.ID,
			PaymentURL: fmt.Sprintf("%s/payments/%s", s.BaseURL, intent.ID),
		}

		json.NewEncoder(w).Encode(response)
	}
}

func (s *Server) HandleShow() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		paymentIntentID := chi.URLParam(r, "paymentIntentID")

		intent, err := s.PaymentIntentRepository.FindOne(paymentIntentID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		if err = s.ShowTemplate.Execute(w, intent); err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		return
	}
}

func (s *Server) HandleComplete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		paymentIntentID := chi.URLParam(r, "paymentIntentID")

		intent, err := s.PaymentIntentRepository.FindOne(paymentIntentID)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if intent.Status == database.PaymentCompleted {
			http.Redirect(w, r, intent.SuccessURL, http.StatusFound)
			return
		}

		redirectURL := intent.FailureURL

		cond1 := s.SendHookRequest(intent.HookURL, intent.CreateHookMessage()) == nil
		cond2 := s.PaymentIntentRepository.MarkAsCompleted(paymentIntentID) == nil

		if cond1 && cond2 {
			redirectURL = intent.SuccessURL
		}

		http.Redirect(w, r, redirectURL, http.StatusFound)
	}
}
