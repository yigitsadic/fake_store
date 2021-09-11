package helper

import "errors"

var (
	unauthorizedError = errors.New("unauthorized to access this resource")
)

func Authenticated(userId interface{}) (string, error) {
	v, ok := userId.(string)
	if !ok {
		return "", unauthorizedError
	}

	if v == "" {
		return "", unauthorizedError
	}

	return v, nil
}
