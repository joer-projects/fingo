package transaction_service

import (
	"fmt"
	"testing"
	"time"

	"github.com/joer-projects/fingo/internal/domain/entity/transaction"
	"github.com/joer-projects/fingo/internal/infra/repo"
)

func TestTransactionService_Add(test *testing.T) {
	test.Run("should add a new transaction", func(test *testing.T) {
		transactionService := NewTransactionService(repo.NewTransactionMemoryRepo())

		AccountNumber := "1234567890"
		Memo := "Test Memo"

		txn, err := transactionService.Add(transaction.TransactionNewProps{
			TypeId:        1,
			AccountNumber: &AccountNumber,
			Memo:          &Memo,
			PostingDate:   time.Now(),
			CreatedBy:     "test",
			TransactionLines: []transaction.TransactionLineNewProps{
				{
					LedgerAccountId:   "1",
					BusinessPartnerId: "1",
					Credit:            100.0,
					Description:       "test",
				},
				{
					LedgerAccountId:   "2",
					BusinessPartnerId: "2",
					Debit:             100.0,
					Description:       "test",
				},
			},
		})

		if err != nil {
			test.Errorf("error occurred: %v", err)
		}

		fmt.Println(txn)
	})
}
