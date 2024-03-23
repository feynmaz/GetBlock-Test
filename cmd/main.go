package main

import (
	"fmt"
	"log"

	"github.com/feynmaz/GetBlock-Test/adapters"
	"github.com/feynmaz/GetBlock-Test/app"
	"github.com/feynmaz/GetBlock-Test/config"
	"github.com/feynmaz/GetBlock-Test/tools/logger"
)

type getBlockNumberResult struct {
	ID      string `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

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
	blockGetter := adapters.NewGetBlockAdapter(url)
	app := app.NewApp(blockGetter)

	fmt.Println(app)
	return nil
}
