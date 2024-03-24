package transaction

type TransactionsGetter interface {
	GetTransactions(numberOfBlocks int) ([]Transaction, error)
}
