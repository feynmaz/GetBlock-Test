package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func DoPostRequest(url, requestBodyRaw string) ([]byte, error) {
	requestBody := bytes.NewBuffer([]byte(requestBodyRaw))
	request, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return responseBody, nil
}
