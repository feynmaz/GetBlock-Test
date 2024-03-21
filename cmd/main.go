package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/feynmaz/GetBlock-Test/config"
	"github.com/feynmaz/GetBlock-Test/tools/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.GetDefault()

	if err := logger.InitLogrus(cfg.LogLevel, cfg.LogJson); err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}

	url := fmt.Sprintf("https://go.getblock.io/%s", cfg.AccessToken)
	bodyJson := `{
		"id": "blockNumber",
		"jsonrpc": "2.0",
		"method": "eth_getBlockByNumber",
		"params": [
			"latest",
			false
		]
	}`
	requestBody := bytes.NewBuffer([]byte(bodyJson))
	request, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Print the response body
	fmt.Println("Response:", string(responseBody))

	return nil
}
