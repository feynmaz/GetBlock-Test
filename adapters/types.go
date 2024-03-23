package adapters

type BlockNumber struct {
	ID      string `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

type Block struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Difficulty       string        `json:"difficulty"`
		ExtraData        string        `json:"extraData"`
		GasLimit         string        `json:"gasLimit"`
		GasUsed          string        `json:"gasUsed"`
		Hash             string        `json:"hash"`
		LogsBloom        string        `json:"logsBloom"`
		Miner            string        `json:"miner"`
		MixHash          string        `json:"mixHash"`
		Nonce            string        `json:"nonce"`
		Number           string        `json:"number"`
		ParentHash       string        `json:"parentHash"`
		ReceiptsRoot     string        `json:"receiptsRoot"`
		Sha3Uncles       string        `json:"sha3Uncles"`
		Size             string        `json:"size"`
		StateRoot        string        `json:"stateRoot"`
		Timestamp        string        `json:"timestamp"`
		TotalDifficulty  string        `json:"totalDifficulty"`
		Transactions     []Transaction `json:"transactions"`
		TransactionsRoot string        `json:"transactionsRoot"`
		Uncles           []any         `json:"uncles"`
	} `json:"result"`
}

type Transaction struct {
	GasPrice             string        `json:"gasPrice"`
	MaxFeePerGas         string        `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string        `json:"maxPriorityFeePerGas"`
	Nonce                string        `json:"nonce"`
	To                   string        `json:"to"`
	V                    string        `json:"v"`
	YParity              string        `json:"yParity"`
	BlockHash            string        `json:"blockHash"`
	From                 string        `json:"from"`
	Type                 string        `json:"type"`
	AccessList           []interface{} `json:"accessList"`
	R                    string        `json:"r"`
	S                    string        `json:"s"`
	BlockNumber          string        `json:"blockNumber"`
	TransactionIndex     string        `json:"transactionIndex"`
	Gas                  string        `json:"gas"`
	Hash                 string        `json:"hash"`
	Input                string        `json:"input"`
	Value                string        `json:"value"`
	ChainId              string        `json:"chainId"`
}
