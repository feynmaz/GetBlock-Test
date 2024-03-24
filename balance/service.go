package balance

import (
	"math/big"

	"github.com/feynmaz/GetBlock-Test/transaction"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetBalances(transactions []transaction.Transaction) Balances {
	balances := make(Balances)
	var ok bool
	var balanceFrom *big.Int
	var balanceTo *big.Int

	for _, transaction := range transactions {
		if balanceFrom, ok = balances[address(transaction.From)]; !ok {
			balanceFrom = big.NewInt(0)
		}
		newBalanceFrom := big.NewInt(0)
		newBalanceFrom.Sub(balanceFrom, transaction.Value)
		balances[address(transaction.From)] = newBalanceFrom

		if balanceTo, ok = balances[address(transaction.To)]; !ok {
			balanceTo = big.NewInt(0)
		}
		newBalanceTo := big.NewInt(0)
		newBalanceTo.Add(balanceTo, transaction.Value)
		balances[address(transaction.To)] = newBalanceTo
	}
	return balances
}

func (s *Service) GetBiggestBalanceChange(balances Balances) (address, *big.Int) {
	maxAbsBalance := big.NewInt(0)
	maxAbsBalanceAddress := address("")

	for address, balance := range balances {
		absBalance := new(big.Int).Abs(balance)
		if absBalance.Cmp(maxAbsBalance) == 1 {
			maxAbsBalanceAddress = address
			maxAbsBalance = absBalance
		}
	}
	_ = maxAbsBalance

	return maxAbsBalanceAddress, balances[maxAbsBalanceAddress]
}
