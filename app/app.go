package app

import (
	"errors"
	"fmt"

	"github.com/feynmaz/GetBlock-Test/balance"
	"github.com/feynmaz/GetBlock-Test/block"
)

var (
	ErrGetBlocks = errors.New("failed to get latest blocks")
)

type app struct {
	blockGetter    block.BlockGetter
	balanceService balance.Service
}

func NewApp(blockGetter block.BlockGetter) *app {
	return &app{
		blockGetter: blockGetter,
	}
}

func (a *app) GetBiggestBalanceChange(blockCount uint) (string, error) {
	blocks, err := a.blockGetter.GetLatestBlocks(blockCount)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGetBlocks, err)
	}

	balances := a.balanceService.GetBalances(blocks)
	address, change := a.balanceService.GetBiggestBalanceChange(balances)
	
	fmt.Printf("address %s changed for %s Gwei in last %d blocks", address, change, blockCount)

	return string(address), nil
}
