package main

import (
	"context"
	to "fingo/cmd/to"
	transaction "fingo/internal/domain/entity/transaction"
	db "fingo/internal/infra/db"
	"time"

	pgx "github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

var ctx context.Context

func SeedAccounting(client *db.Queries) {
	AccountNumber := "123456789"
	Memo := "Memo"

	transaction, err := transaction.New(transaction.NewProps{
		TypeID:        1,
		AccountNumber: &AccountNumber,
		Memo:          &Memo,
		PostingDate:   time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		CreatedBy:     "test",
		TransactionLines: []transaction.TransactionLineCreateProps{
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

	raw := transaction.ToRaw()
	err = client.AddTransaction(ctx, db.AddTransactionParams{
		ID:                raw.ID,
		ProjectID:         to.StringToPGText(raw.ProjectID),
		TransactionTypeID: raw.TransactionTypeID,
		AccountNumber:     to.StringToPGText(raw.AccountNumber),
		Memo:              to.StringToPGText(raw.Memo),
		PostingDate:       to.TimeToPGTimestamp(&raw.PostingDate),
		CreatedBy:         raw.CreatedBy,
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	ctx = context.Background()
	pgConnection, err := pgx.Connect(ctx, "postgres://fingo:fingo@localhost:5432/fingo_accounting?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer pgConnection.Close(ctx)

	client := db.New(pgConnection)

	SeedAccounting(client)
}
