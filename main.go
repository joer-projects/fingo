package main

import (
	tx "fingo/internal/domain/entity/transaction"
	"fmt"
	"time"

	uuid "github.com/google/uuid"
)

func main() {
	AccountNumber := "123456789"
	Memo := "Memo"

	transaction, err := tx.New(tx.NewProps{
		TypeID:        1,
		AccountNumber: &AccountNumber,
		Memo:          &Memo,
		PostingDate:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		CreatedBy:     "test",
		TransactionLines: []tx.TransactionLineCreateProps{
			{
				LedgerAccountID:   "123456789",
				BusinessPartnerID: "123456789",
				ProjectID:         "123456789",
				Credit:            100.0,
				Debit:             0,
				Description:       "test",
			},
			{
				LedgerAccountID:   "123456789",
				BusinessPartnerID: "123456789",
				ProjectID:         "123456789",
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
