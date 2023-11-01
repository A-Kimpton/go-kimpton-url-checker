package tasks

import (
	"encoding/json"
	"log"
	"net/http"

	"go.kimpton.io/url-checker/internal/domain"
	"go.kimpton.io/url-checker/internal/domain/messages"
)

var _ domain.Task = &HTTPChecker{}

type HTTPChecker struct{}

func (h *HTTPChecker) Execute(data []byte) (domain.ResultCode, error) {

	var req messages.HTTPSRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		return domain.ResultCodeInternalError, err
	}

	resp, err := http.Get(req.Url)
	if err != nil {
		return domain.ResultCodeInternalError, err
	}
	defer resp.Body.Close()

	// TODO: Handle the response
	log.Printf("Response: %+v\n", resp.Status)

	// Task successfully handled, reguardless of the outcome of the request
	return domain.ResultCodeSuccess, nil
}
