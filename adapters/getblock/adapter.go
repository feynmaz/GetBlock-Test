package getblock

import (
	"encoding/json"
	"fmt"

	"github.com/feynmaz/GetBlock-Test/tools/http"
	"github.com/feynmaz/GetBlock-Test/transaction"
)

type GetBlockAdapter struct {
	url string
}

func NewGetBlockAdapter(url string) *GetBlockAdapter {
	return &GetBlockAdapter{url: url}
}

// Implements BlockGetter interface
func (a *GetBlockAdapter) GetTransactions(numberOfBlocks int) ([]transaction.Transaction, error) {
	// lastBlockNumber, err := a.getLastBlockNumber()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get last block number: %w", err)
	// }

	// blocks := make([]block.Block, n)
	// blockNumber := lastBlockNumber

	// for i := 0; i < int(n); i++ {
	// 	block, err := a.GetBlockByNumber(blockNumber)
	// 	if err != nil {
	// 		return nil, fmt.Errorf()
	// 	}
	// }

	return nil, nil
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

func (a *GetBlockAdapter) GetBlockByNumber(blockNumber string) (Block, error) {
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
