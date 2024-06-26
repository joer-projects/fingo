package main

import (
	"fmt"
	"time"

	Transaction "github.com/JoerProjects/fingo/internal/domain/entity/transaction"

	"github.com/google/uuid"
)

func main() {
	AccountNumber := "123456789"
	Memo := "Memo"

	transaction, err := Transaction.New(Transaction.TransactionNewProps{
		TypeId:        1,
		AccountNumber: &AccountNumber,
		Memo:          &Memo,
		PostingDate:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		CreatedBy:     "test",
		TransactionLines: []Transaction.TransactionLineCreateProps{
			{
				LedgerAccountId:   "123456789",
				BusinessPartnerId: "123456789",
				ProjectId:         "123456789",
				Credit:            100.0,
				Debit:             0,
				Description:       "test",
			},
			{
				LedgerAccountId:   "123456789",
				BusinessPartnerId: "123456789",
				ProjectId:         "123456789",
				Credit:            0,
				Debit:             100,
				Description:       "test",
			},
		},
	}, uuid.New().String())

	if err != nil {
		panic(err)
	}

	json, err := transaction.ToJson()

	if err != nil {
		panic(err)
	}

	fmt.Print(string(json))
}
