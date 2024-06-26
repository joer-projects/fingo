package seed

import (
	"context"

	"github.com/joer-projects/fingo/internal/infra/db/accounting_db"
	"github.com/joer-projects/fingo/internal/infra/db/connection/postgres"
)

func Seed() error {
	ctx := context.Background()
	dbConn, err := postgres.NewConnection(ctx)
	if err != nil {
		return err
	}

	query := accounting_db.New(dbConn)

	dbTx, err := dbConn.Begin(ctx)
	if err != nil {
		return err
	}
	defer dbTx.Rollback(ctx)
	qtx := query.WithTx(dbTx)

	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 1, Name: "Journal Entry"})
	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 2, Name: "Vendor Payment"})
	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 3, Name: "Incoming Payment"})
	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 4, Name: "Deposit"})
	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 5, Name: "A/P Invoice"})
	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 6, Name: "A/P Down Payment"})
	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 7, Name: "A/P Credit Memo"})
	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 8, Name: "A/R Invoice"})
	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 9, Name: "A/R Down Payment"})
	qtx.SeedTransactionType(ctx, accounting_db.SeedTransactionTypeParams{AccountingTransactionTypeID: 10, Name: "A/R Credit Memo"})

	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 1, Name: "Asset"})
	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 2, Name: "Equity"})
	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 3, Name: "Expense"})
	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 4, Name: "Liability"})
	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 5, Name: "Revenue"})

	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 1, Name: "Bank"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 2, Name: "OTHER_CURRENT_ASSET"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 3, Name: "FIXED_ASSET"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 4, Name: "OTHER_ASSET"})

	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 5, Name: "EQUITY"})

	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 6, Name: "EXPENSE"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 7, Name: "OTHER_EXPENSE"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 8, Name: "COST_OF_GOODS_SOLD"})

	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 9, Name: "ACCOUNTS_PAYABLE"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 10, Name: "CREDIT_CARD"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 11, Name: "LONG_TERM_LIABILITY"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 12, Name: "OTHER_CURRENT_LIABILITY"})

	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 13, Name: "INCOME"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 14, Name: "OTHER_INCOME"})

	err = dbTx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}
