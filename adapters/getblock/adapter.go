package getblock

import (
	"encoding/json"
	"fmt"

	"github.com/feynmaz/GetBlock-Test/tools/hex"
	"github.com/feynmaz/GetBlock-Test/tools/http"
	"github.com/feynmaz/GetBlock-Test/transaction"
)

type GetBlockAdapter struct {
	url string
}

func NewGetBlockAdapter(url string) *GetBlockAdapter {
	return &GetBlockAdapter{url: url}
}

// Implements TransactionsGetter interface
func (a *GetBlockAdapter) GetTransactions(numberOfBlocks int) ([]transaction.Transaction, error) {
	lastBlockNumber, err := a.getLastBlockNumber()
	if err != nil {
		return nil, fmt.Errorf("failed to get last block number: %w", err)
	}

	transactions := make([]transaction.Transaction, 0, numberOfBlocks)
	blockNumber := lastBlockNumber

	for i := 0; i < numberOfBlocks; i++ {
		block, err := a.getBlockByNumber(blockNumber)
		if err != nil {
			return nil, fmt.Errorf("failed to get block by number: %w", err)
		}

		for _, tr := range block.Result.Transactions {

			value, err := hex.HexToBigInt(tr.Value)
			if err != nil {
				return nil, fmt.Errorf("failed to convert hex to big int: %w", err)
			}

			transaction := transaction.Transaction{
				From:  tr.From,
				To:    tr.To,
				Value: value,
			}
			transactions = append(transactions, transaction)
		}

		blockNumber, err = a.getBlockNumberFromHash(block.Result.ParentHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get block number from hash: %w", err)
		}
	}

	return transactions, nil
}

func (a *GetBlockAdapter) getLastBlockNumber() (string, error) {
	getBlockNumber := `{
		"jsonrpc": "2.0",
		"method": "eth_blockNumber",
		"params": [],
		"id": "getblock.io"
	}`

	responseBody, err := http.DoPostRequest(a.url, getBlockNumber)
	if err != nil {
		return "", fmt.Errorf("failed to do post request: %w", err)
	}

	var blockNumber BlockNumber
	if err := json.Unmarshal(responseBody, &blockNumber); err != nil {
		return "", fmt.Errorf("failed to unmarshall result: %w", err)
	}

	return blockNumber.Result, nil
}

func (a *GetBlockAdapter) getBlockByNumber(blockNumber string) (Block, error) {
	getBlockByNumber := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByNumber",
		"params": [
			"%s",
			true
		],
		"id": "getblock.io"
	}`, blockNumber)

	responseBody, err := http.DoPostRequest(a.url, getBlockByNumber)
	if err != nil {
		return Block{}, fmt.Errorf("failed to do post request: %w", err)
	}

	var block Block
	if err := json.Unmarshal(responseBody, &block); err != nil {
		return Block{}, fmt.Errorf("failed to unmarshall result: %w", err)
	}

	if block.Result.Hash == "" {
		return Block{}, fmt.Errorf("no block with number %s", blockNumber)
	}

	return block, nil
}

func (a *GetBlockAdapter) getBlockNumberFromHash(blockHash string) (string, error) {
	getBlockByNumber := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByHash",
		"params": [
			"%s",
			false
		],
		"id": "getblock.io"
	}`, blockHash)

	responseBody, err := http.DoPostRequest(a.url, getBlockByNumber)
	if err != nil {
		return "", fmt.Errorf("failed to do post request: %w", err)
	}

	var res struct {
		Result struct {
			Number string `json:"number"`
		} `json:"result"`
	}
	if err := json.Unmarshal(responseBody, &res); err != nil {
		return "", fmt.Errorf("failed to unmarshall result: %w", err)
	}
	return res.Result.Number, nil
}
