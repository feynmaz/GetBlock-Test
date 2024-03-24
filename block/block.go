package block

import "math/big"

type Block struct {
	Transactions []Transaction
}

type Transaction struct {
	From  string
	To    string
	Value *big.Int
}
