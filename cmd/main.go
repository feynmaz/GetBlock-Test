package main

import (
	"fmt"
	"log"

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

	
	return nil
}
