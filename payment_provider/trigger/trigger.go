package trigger

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/cenkalti/backoff/v4"
	"github.com/yigitsadic/fake_store/payment_provider/database"
	"net/http"
	"time"
)

func buildRequest(hookURL string, payload database.PaymentHookMessage) (*http.Request, error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, hookURL, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

func makeRequest(req *http.Request) error {
	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		return nil
	}

	return errors.New("server not returned with successful status code")
}

// SendHookRequest makes several requests to given hookURL address. After 5 try it fails.
func SendHookRequest(hookURL string, payload database.PaymentHookMessage) error {
	req, err := buildRequest(hookURL, payload)
	if err != nil {
		return err
	}

	operation := func() error {
		return makeRequest(req)
	}

	return backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))
}
