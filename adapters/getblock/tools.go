package getblock

import (
	"fmt"

	"github.com/feynmaz/GetBlock-Test/tools/hex"
	"github.com/feynmaz/GetBlock-Test/transaction"
)

func GetTransactionsFromBlock(block Block) ([]*transaction.Transaction, error) {
	transactions := make([]*transaction.Transaction, 0)

	for _, tr := range block.Result.Transactions {
		if tr.BlockNumber == "" {
			// If transaction is not successfull
			continue
		}

		value, err := hex.HexToBigInt(tr.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to convert value to int: %w", err)
		}
		gas, err := hex.HexToBigInt(tr.Gas)
		if err != nil {
			return nil, fmt.Errorf("failed to convert gas to int: %w", err)
		}

		transaction := &transaction.Transaction{
			From:  tr.From,
			To:    tr.To,
			Value: value,
			Gas:   gas,
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
