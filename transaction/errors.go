package transaction

import "errors"

var (
	ErrNoBlocksRead = errors.New("must read at least 1 block")
)
