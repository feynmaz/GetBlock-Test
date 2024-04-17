package main

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"

	"github.com/feynmaz/GetBlock-Test/adapters/getblock"
	"github.com/feynmaz/GetBlock-Test/app"
	"github.com/feynmaz/GetBlock-Test/config"
	"github.com/feynmaz/GetBlock-Test/tools/logger"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var blockCount int = 10

	cfg := config.GetDefault()

	if err := logger.InitLogrus(cfg.LogLevel, cfg.LogJson); err != nil {
		return fmt.Errorf("failed to init logger: %w", err)
	}
	logrus.Debugf("config: %s", cfg)

	url := fmt.Sprintf("https://go.getblock.io/%s", cfg.AccessToken)
	transactionsGetter := getblock.NewGetBlockAdapter(cfg.ApiTimeout, url)
	app := app.NewApp(transactionsGetter)

	addressWithBiggestChange, err := app.GetBiggestBalanceChange(blockCount)
	if err != nil {
		return fmt.Errorf("failed to get biggest balance change: %w", err)
	}
	logrus.Infof("address with the biggest change: %s", addressWithBiggestChange)

	return nil
}
