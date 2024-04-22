package getblock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	httptools "github.com/feynmaz/GetBlock-Test/tools/http"
	"github.com/feynmaz/GetBlock-Test/transaction"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type GetBlockAdapter struct {
	httpClient *http.Client
	url        string
}

func NewGetBlockAdapter(timeout time.Duration, url string) *GetBlockAdapter {
	httpClient := &http.Client{
		Timeout: timeout,
	}
	return &GetBlockAdapter{
		httpClient: httpClient,
		url:        url,
	}
}

// Implements TransactionsGetter interface
func (a *GetBlockAdapter) GetTransactions(numberOfBlocks int) ([]*transaction.Transaction, error) {
	lastBlockHash, err := a.getLastBlockHash()
	if err != nil {
		return nil, fmt.Errorf("failed to get last block number: %w", err)
	}

	transactionsAll := make([]*transaction.Transaction, 0, numberOfBlocks)
	blockHash := lastBlockHash

	transactionsCh := make(chan []*transaction.Transaction, 1)
	doneCh := make(chan struct{})
	g := new(errgroup.Group)

	go func() {
		defer close(doneCh)
		for tr := range transactionsCh {
			transactionsAll = append(transactionsAll, tr...)
		}
	}()

	for i := 0; i < numberOfBlocks; i++ {
		block, err := a.getBlockByHash(blockHash)
		if err != nil {
			return nil, fmt.Errorf("failed to get block by number: %w", err)
		}

		g.Go(func() error {

			transactionsNew, err := GetTransactionsFromBlock(block)
			if err != nil {
				return fmt.Errorf("failed to get transactions from block: %w", err)
			}
			transactionsCh <- transactionsNew

			return nil
		})

		blockHash = block.Result.ParentHash
		if i%10 == 0 {
			logrus.Infof("read %d blocks, %d left", i+1, numberOfBlocks-i-1)
		}
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}
	close(transactionsCh)
	<-doneCh

	return transactionsAll, nil
}

func (a *GetBlockAdapter) getLastBlockHash() (string, error) {
	getBlockNumber := `{
		"jsonrpc": "2.0",
		"method": "eth_blockNumber",
		"params": [],
		"id": "getblock.io"
	}`

	responseBody, err := httptools.DoPostRequest(a.httpClient, a.url, getBlockNumber)
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

	responseBody, err = httptools.DoPostRequest(a.httpClient, a.url, getBlockByNumber)
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

func (a *GetBlockAdapter) getBlockByHash(blockHash string) (*Block, error) {
	getBlockByNumber := fmt.Sprintf(`{
		"jsonrpc": "2.0",
		"method": "eth_getBlockByHash",
		"params": [
			"%s",
			true
		],
		"id": "getblock.io"
	}`, blockHash)

	responseBody, err := httptools.DoPostRequest(a.httpClient, a.url, getBlockByNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to do post request: %w", err)
	}

	var block Block
	if err := json.Unmarshal(responseBody, &block); err != nil {
		return nil, fmt.Errorf("failed to unmarshall result: %w", err)
	}

	if block.Result.Hash == "" {
		return nil, fmt.Errorf("no block with hash %s", blockHash)
	}

	return &block, nil
}
