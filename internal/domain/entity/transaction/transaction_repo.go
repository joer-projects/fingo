package transaction

type TransactionRepo interface {
	Get(id string) (*Transaction, error)
	Add(t *Transaction) (*TransactionRaw, error)
	Save(t *Transaction) (*TransactionRaw, error)
}
