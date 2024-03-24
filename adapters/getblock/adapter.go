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
	lastBlockHash, err := a.getLastBlockHash()
	if err != nil {
		return nil, fmt.Errorf("failed to get last block number: %w", err)
	}

	transactions := make([]transaction.Transaction, 0, numberOfBlocks)
	blockHash := lastBlockHash

	for i := 0; i < numberOfBlocks; i++ {
		block, err := a.getBlockByHash(blockHash)
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

		blockHash = block.Result.ParentHash
	}

	return transactions, nil
}

func (a *GetBlockAdapter) getLastBlockHash() (string, error) {
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

	var res struct {
		Result string `json:"result"`
	}
	if err := json.Unmarshal(responseBody, &res); err != nil {
		return "", fmt.Errorf("failed to unmarshall result: %w", err)
	}

	blockNumber := res.Result
	getBlockByNumber := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByNumber",
		"params": [
			"%s",
			false
		],
		"id": "getblock.io"
	}`, blockNumber)

	responseBody, err = http.DoPostRequest(a.url, getBlockByNumber)
	if err != nil {
		return "", fmt.Errorf("failed to do post request: %w", err)
	}

	var block struct {
		Result struct {
			Hash string `json:"hash"`
		} `json:"result"`
	}
	if err := json.Unmarshal(responseBody, &block); err != nil {
		return "", fmt.Errorf("failed to unmarshall result: %w", err)
	}

	return block.Result.Hash, nil
}

func (a *GetBlockAdapter) getBlockByHash(blockHash string) (Block, error) {
	getBlockByNumber := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByHash",
		"params": [
			"%s",
			true
		],
		"id": "getblock.io"
	}`, blockHash)

	responseBody, err := http.DoPostRequest(a.url, getBlockByNumber)
	if err != nil {
		return Block{}, fmt.Errorf("failed to do post request: %w", err)
	}

	var block Block
	if err := json.Unmarshal(responseBody, &block); err != nil {
		return Block{}, fmt.Errorf("failed to unmarshall result: %w", err)
	}

	if block.Result.Hash == "" {
		return Block{}, fmt.Errorf("no block with hash %s", blockHash)
	}

	return block, nil
}
