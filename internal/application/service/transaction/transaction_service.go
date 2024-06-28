package transaction_service

import (
	"github.com/google/uuid"
	"github.com/joer-projects/fingo/internal/domain/entity/transaction"
)

type TransactionService struct {
	repo transaction.TransactionRepo
}

func NewTransactionService(repo transaction.TransactionRepo) TransactionService {
	return TransactionService{
		repo,
	}
}

func (t *TransactionService) Add(input transaction.TransactionNewProps) (transaction.TransactionRaw, error) {
	txn, err := transaction.NewTransaction(transaction.TransactionNewProps{
		TypeId:           input.TypeId,
		AccountNumber:    input.AccountNumber,
		Memo:             input.Memo,
		PostingDate:      input.PostingDate,
		CreatedBy:        input.CreatedBy,
		TransactionLines: input.TransactionLines,
	}, uuid.New().String())

	if err != nil {
		return transaction.TransactionRaw{}, err
	}

	raw, err := t.repo.Add(&txn)

	if err != nil {
		return transaction.TransactionRaw{}, err
	}

	return raw, nil
}
