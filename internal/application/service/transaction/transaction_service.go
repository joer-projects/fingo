package transaction_service

import (
	Transaction "github.com/JoerProjects/fingo/internal/domain/entity/transaction"
	"github.com/google/uuid"
)

type TransactionService struct {
	repo Transaction.TransactionRepo
}

func NewTransactionService(repo Transaction.TransactionRepo) TransactionService {
	return TransactionService{
		repo,
	}
}

func (t *TransactionService) Add(input Transaction.TransactionNewProps) (Transaction.TransactionRaw, error) {
	txn, err := Transaction.New(Transaction.TransactionNewProps{
		TypeId:           input.TypeId,
		AccountNumber:    input.AccountNumber,
		Memo:             input.Memo,
		PostingDate:      input.PostingDate,
		CreatedBy:        input.CreatedBy,
		TransactionLines: input.TransactionLines,
	}, uuid.New().String())

	if err != nil {
		return Transaction.TransactionRaw{}, err
	}

	raw, err := t.repo.Add(txn)

	if err != nil {
		return Transaction.TransactionRaw{}, err
	}

	return raw, nil
}
