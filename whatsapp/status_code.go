package whatsapp

import (
	"errors"
	"net/http"
)

func validateStatus(res *http.Response, b []byte) error {
	if res.StatusCode >= http.StatusBadRequest {
		return errors.New(string(b))
	}
	return nil
}
