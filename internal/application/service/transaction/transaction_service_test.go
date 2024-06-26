package transaction_service

import (
	"fmt"
	"testing"
	"time"

	"github.com/JoerProjects/fingo/internal/domain/entity/transaction"
	"github.com/JoerProjects/fingo/internal/infra/repository"
)

func TestTransactionService_Add(t *testing.T) {
	t.Run("should add a new transaction", func(t *testing.T) {
		transactionService := NewTransactionService(repository.NewTransactionMemoryRepo())

		AccountNumber := "1234567890"
		Memo := "Test Memo"

		transaction, err := transactionService.Add(transaction.TransactionNewProps{
			TypeId:        1,
			AccountNumber: &AccountNumber,
			Memo:          &Memo,
			PostingDate:   time.Now(),
			CreatedBy:     "test",
			TransactionLines: []transaction.TransactionLineCreateProps{
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
			t.Errorf("error occurred: %v", err)
		}

		fmt.Println(transaction)
	})
}
