package block

type BlockGetter interface {
	GetLatestBlocks(n uint) ([]Block, error)
}
