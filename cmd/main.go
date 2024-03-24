package main

import (
	"fmt"
	"log"

	"github.com/feynmaz/GetBlock-Test/app"
	"github.com/feynmaz/GetBlock-Test/block"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// cfg := config.GetDefault()

	// if err := logger.InitLogrus(cfg.LogLevel, cfg.LogJson); err != nil {
	// 	return fmt.Errorf("failed to init logger: %w", err)
	// }

	// url := fmt.Sprintf("https://go.getblock.io/%s", cfg.AccessToken)
	// fmt.Println(url)

	var blockCount uint = 100

	blockGetter := block.NewMockBlockGetter()
	app := app.NewApp(blockGetter)

	addressWithBiggestChange, err := app.GetBiggestBalanceChange(blockCount)
	if err != nil {
		return fmt.Errorf("failed to get biggest balance change: %w", err)
	}
	fmt.Printf("address with biggest change: %s", addressWithBiggestChange)

	return nil
}
