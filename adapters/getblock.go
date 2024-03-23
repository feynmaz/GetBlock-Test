package adapters

import (
	"encoding/json"
	"fmt"

	"github.com/feynmaz/GetBlock-Test/tools/http"
)

type getBlockAdapter struct {
	url string
}

func NewGetBlockAdapter(url string) *getBlockAdapter {
	return &getBlockAdapter{url: url}
}

func (a *getBlockAdapter) GetLastBlockNumber() (string, error) {
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

func (a *getBlockAdapter) GetBlockByNumber(blockNumber string) (Block, error) {
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

func (a *getBlockAdapter) GetBlockNumberFromHash(blockHash string) (string, error) {
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
