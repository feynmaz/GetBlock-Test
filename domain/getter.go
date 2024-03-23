package domain

type BlockDataGetter interface {
	GetLastBlockNumber() (string, error)
	GetBlockByNumber(blockNumber string) (Block, error)
	GetBlockNumberFromHash(blockHash string) (string, error)
}
