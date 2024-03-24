package balance

import (
	"fmt"
	"math/big"

	"github.com/feynmaz/GetBlock-Test/block"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetBalances(blocks []block.Block) Balances {
	balances := make(Balances)
	var ok bool
	var balanceFrom *big.Int
	var balanceTo *big.Int

	for _, b := range blocks {
		for _, transaction := range b.Transactions {
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
	}
	return balances
}

func (s *Service) GetBiggestBalanceChange(balances Balances) (address, *big.Int) {
	maxAbsBalance := big.NewInt(0)
	maxAbsBalanceAddress := address("")

	for address, balance := range balances {
		absBalance := new(big.Int).Abs(balance)
		if absBalance.Cmp(balance) == 1 {
			maxAbsBalanceAddress = address
			maxAbsBalance = absBalance
		}
	}
	fmt.Println(maxAbsBalance)
	return maxAbsBalanceAddress, maxAbsBalance
}
