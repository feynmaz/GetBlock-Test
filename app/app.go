package app

import (
	"errors"
	"fmt"

	"github.com/feynmaz/GetBlock-Test/balance"
	"github.com/feynmaz/GetBlock-Test/transaction"
	"github.com/sirupsen/logrus"
)

var (
	ErrGetTransactions = errors.New("failed to get transactions")
)

type app struct {
	transactionsGetter transaction.TransactionsGetter
	balanceService     balance.Service
}

func NewApp(transactionsGetter transaction.TransactionsGetter) *app {
	return &app{
		transactionsGetter: transactionsGetter,
	}
}

func (a *app) GetBiggestBalanceChange(blockCount int) (string, error) {
	if blockCount <= 0 {
		return "", transaction.ErrNoBlocksRead
	}

	logrus.Info("getting transactions")
	transactions, err := a.transactionsGetter.GetTransactions(blockCount)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGetTransactions, err)
	}

	logrus.Info("calculating balances")
	balances := a.balanceService.GetBalances(transactions)

	logrus.Info("calculating biggest balance change")
	address, change := a.balanceService.GetBiggestBalanceChange(balances)

	logrus.Infof("address %s changed for %s Gwei in the last %d blocks", address, change, blockCount)

	return string(address), nil
}
