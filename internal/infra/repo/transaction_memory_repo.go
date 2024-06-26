package repo

import Transaction "github.com/joer-projects/fingo/internal/domain/entity/transaction"

type TransactionMemoryRepo struct {
	transactions map[string]Transaction.Transaction
}

func (r *TransactionMemoryRepo) Get(id string) (Transaction.Transaction, error) {
	return r.transactions[id], nil
}

func (r *TransactionMemoryRepo) Add(t Transaction.Transaction) (Transaction.TransactionRaw, error) {
	r.transactions[t.GetID()] = t
	return Transaction.TransactionRaw{}, nil
}

func (r *TransactionMemoryRepo) Save(t Transaction.Transaction) (Transaction.TransactionRaw, error) {
	r.transactions[t.GetID()] = t
	return Transaction.TransactionRaw{}, nil
}

func NewTransactionMemoryRepo() Transaction.TransactionRepo {
	return &TransactionMemoryRepo{
		transactions: make(map[string]Transaction.Transaction),
	}
}
