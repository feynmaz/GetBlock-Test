package transaction

import "math/big"

type Transaction struct {
	From  string
	To    string
	Value *big.Int
	Gas   *big.Int
}
