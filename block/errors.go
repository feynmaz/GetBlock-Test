package block

import "errors"

var (
	ErrNoBlocksRequested = errors.New("must request at least 1 block")
)
