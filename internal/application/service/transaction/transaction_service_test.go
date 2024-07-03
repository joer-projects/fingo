package transaction_service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joer-projects/fingo/internal/domain/entity/transaction"
	"github.com/joer-projects/fingo/internal/infra/db/accounting_db"
	"github.com/joer-projects/fingo/internal/infra/db/connection/postgres"
	"github.com/joer-projects/fingo/internal/infra/repo"
)

func TestTransactionService_Add(test *testing.T) {
	test.Run("should add a new transaction", func(test *testing.T) {
		ctx := context.Background()
		pgConn, err := postgres.NewConnection(ctx)
		if err != nil {
			test.Errorf("error occurred: %v", err)
		}
		defer pgConn.Close(ctx)
		queries := accounting_db.New(pgConn)
		transactionPostgresRepo, err := repo.NewTransactionPostgresRepo(ctx, queries)
		if err != nil {
			test.Errorf("error occurred: %v", err)
		}
		transactionService := NewTransactionService(transactionPostgresRepo)

		Memo := "Test Memo"
		ProjectId := "Test Project ID"

		input := transaction.TransactionNewProps{
			TypeId:      1,
			Memo:        &Memo,
			ProjectId:   &ProjectId,
			PostingDate: time.Now(),
			CreatedBy:   "test",
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
		}

		txn, err := transactionService.Add(&input)

		if err != nil {
			test.Errorf("error occurred: %v", err)
		}

		getTransaction, err := transactionPostgresRepo.Get(txn.Id)
		getTransactionRaw := getTransaction.ToRaw()

		fmt.Println(txn.Memo)
		fmt.Println(getTransactionRaw.ProjectId)

		if err != nil {
			test.Errorf("error occurred: %v", err)
		}

		if txn.Id != getTransactionRaw.Id {
			test.Errorf("ids are not equal")
		}

		if *txn.Memo != *getTransactionRaw.Memo {
			test.Errorf("memos are not equal")
		}

		// if *txn.ProjectId != *getTransactionRaw.ProjectId {
		// 	test.Errorf("project ids are not equal")
		// }

	})
}
