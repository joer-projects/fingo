package main

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

	defer dbConn.Close(ctx)

	query := accounting_db.New(dbConn)

	dbTx, err := dbConn.Begin(ctx)
	if err != nil {
		return err
	}
	defer dbTx.Rollback(ctx)
	qtx := query.WithTx(dbTx)

	// Transaction Types
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

	// Ledger Account Classes
	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 1, Name: "Asset"})
	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 2, Name: "Equity"})
	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 3, Name: "Expense"})
	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 4, Name: "Liability"})
	qtx.SeedLedgerAccountClass(ctx, accounting_db.SeedLedgerAccountClassParams{LedgerAccountClassID: 5, Name: "Revenue"})

	// Asset
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 1, Name: "Bank"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 2, Name: "Other Current Asset"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 3, Name: "Fixed Asset"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 4, Name: "Other Asset"})
	// Equity
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 5, Name: "Equity"})
	// Expense
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 6, Name: "Expense"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 7, Name: "Other Expense"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 8, Name: "Cost of Goods Sold"})
	// Liability
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 9, Name: "Accounts Payable"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 10, Name: "Credit Card"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 11, Name: "Long Term Liability"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 12, Name: "Other Current Liability"})
	// Revenue
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 13, Name: "Income"})
	qtx.SeedLedgerAccountType(ctx, accounting_db.SeedLedgerAccountTypeParams{LedgerAccountTypeID: 14, Name: "Other Income"})

	// // Ledger Accounts Sub-Types

	// Bank
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 1, Name: "Cash on Hand"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 2, Name: "Checking"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 3, Name: "Money Market"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 4, Name: "Rents Held in Trust"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 5, Name: "Savings"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 6, Name: "Trust Accounts"})
	// Other Current Asset
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 7, Name: "Allowance for Bad Debts"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 8, Name: "Development Costs"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 9, Name: "Employee Cash Advances"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 10, Name: "Other Current Assets"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 11, Name: "Inventory"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 12, Name: "Investment Mortgage Real Estate Loans"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 13, Name: "Investment Other"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 14, Name: "Investment Tax Exempt Securities"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 15, Name: "Investment US Government Obligations"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 16, Name: "Loans to Officers"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 17, Name: "Loans to Others"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 18, Name: "Loans to Stockholders"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 19, Name: "Prepaid Expenses"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 20, Name: "Retained"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 21, Name: "Undeposited Funds"})
	// Fixed Asset
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 22, Name: "Accumulated Depletion"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 23, Name: "Accumulated Depreciation"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 24, Name: "Depreciable Assets"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 25, Name: "Fixed Asset Computers"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 26, Name: "Fixed Asset Copper"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 27, Name: "Fixed Asset Furniture"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 28, Name: "Fixed Asset Phone"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 29, Name: "Fixed Asset Photo"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 30, Name: "Fixed Asset Software"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 31, Name: "Fixed Asset Other Tools Equipment"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 32, Name: "Furniture and Fixtures"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 33, Name: "Land"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 34, Name: "Leasehold Improvements"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 35, Name: "Other Fixed Assets"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 36, Name: "Accumulated Amortization"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 37, Name: "Buildings"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 38, Name: "Intangible Assets"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 39, Name: "Machinery and Equipment"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 40, Name: "Vehicles"})
	// Other Asset
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 41, Name: "Lease Buyout"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 42, Name: "Other Long Term Assets"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 43, Name: "Security Deposits"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 44, Name: "Accumulated Amortization of Other Assets"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 45, Name: "Goodwill"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 46, Name: "Licenses"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 47, Name: "Organizational Costs"})
	// Accounts Receivable
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 48, Name: "Accounts Receivable"})
	// Equity
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 49, Name: "Opening Balance Equity"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 50, Name: "Partners Equity"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 51, Name: "Retained Earnings"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 52, Name: "Accumulated Adjustment"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 53, Name: "Owners Equity"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 54, Name: "Paid in Capital or Surplus"}) // DEFAULT
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 55, Name: "Partners Contributions"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 56, Name: "Partners Distributions"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 57, Name: "Preferred Stock"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 58, Name: "Common Stock"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 59, Name: "Treasury Stock"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 60, Name: "Estimated Taxes"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 61, Name: "Health Care"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 62, Name: "Personal Income"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 63, Name: "Personal Expense"})
	// Expense
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 64, Name: "Advertising Promotion"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 65, Name: "Bad Debts"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 66, Name: "Bank Charges"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 67, Name: "Charitable Contributions"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 68, Name: "Communications and Fees"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 69, Name: "Entertainment"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 70, Name: "Entertainment Meals"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 71, Name: "Equipment Rental"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 72, Name: "Finance Costs"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 73, Name: "Global Tax Expense"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 74, Name: "Insurance"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 75, Name: "Interest Paid"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 76, Name: "Legal Professional Fees"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 77, Name: "Office Expenses"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 78, Name: "Office General Administrative Expenses"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 79, Name: "Other Business Expenses"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 80, Name: "Other Miscellaneous Service Cost"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 81, Name: "Professional Meals"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 82, Name: "Rent or Lease of Buildings"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 83, Name: "Repair Maintenance"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 84, Name: "Shipping Freight Delivery"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 85, Name: "Supplies Materials"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 86, Name: "Travel"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 87, Name: "Travel Meals"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 88, Name: "Utility"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 89, Name: "Auto"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 90, Name: "Cost of Labor"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 91, Name: "Dues Subscriptions"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 92, Name: "Payroll Expenses"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 93, Name: "Taxes Paid"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 94, Name: "Unapplied Cash Bill Payment Expense"})
	// Other Expense
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 95, Name: "Depreciation"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 96, Name: "Exchange Gain or Loss"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 97, Name: "Other Miscellaneous Expense"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 98, Name: "Penalties Settlements"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 99, Name: "Amortization"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 100, Name: "Gas and Fuel"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 101, Name: "Home Office"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 102, Name: "Home Owner Rental Insurance"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 103, Name: "Other Home Office Expenses"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 104, Name: "Mortgage Interest"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 105, Name: "Rent and Lease"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 106, Name: "Repairs and Maintenance"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 107, Name: "Parking and Tolls"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 108, Name: "Vehicle"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 109, Name: "Vehicle Insurance"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 110, Name: "Vehicle Lease"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 111, Name: "Vehicle Loan Interest"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 112, Name: "Vehicle Loan"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 113, Name: "Vehicle Registration"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 114, Name: "Vehicle Repairs"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 115, Name: "Other Vehicle Expenses"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 116, Name: "Waste and Road Services"})
	// Cost of Goods Sold
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 117, Name: "Equipment Rental Costs"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 118, Name: "Other Costs of Service Costs"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 119, Name: "Shipping Freight Delivery Costs"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 120, Name: "Supplies Materials Costs"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 121, Name: "Cost of Labor Costs"})
	// Accounts Payable
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 122, Name: "Accounts Payable"})
	// Credit Card
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 123, Name: "Credit Card"})
	// Long Term Liability
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 124, Name: "Notes Payable"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 125, Name: "Other Long Term Liabilities"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 126, Name: "Shareholder Notes Payable"})
	// Other Current Liability
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 127, Name: "Direct Deposit Payable"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 128, Name: "Line of Credit"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 129, Name: "Loan Payable"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 130, Name: "Global Tax Payable"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 131, Name: "Global Tax Suspense"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 132, Name: "Other Current Liabilities"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 133, Name: "Payroll Clearing"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 134, Name: "Payroll Tax Payable"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 135, Name: "Prepaid Expenses Payable"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 136, Name: "Rents in Trust Liability"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 137, Name: "Trust Accounts Liabilities"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 138, Name: "Federal Income Tax Payable"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 139, Name: "Insurance Payable"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 140, Name: "Sales Tax Payable"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 141, Name: "State Local Income Tax Payable"})
	// Income
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 142, Name: "Non Profit Income"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 143, Name: "Other Primary Income"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 144, Name: "Sales of Product Income"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 145, Name: "Service Fee Income"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 146, Name: "Discounts Refunds Given"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 147, Name: "Unapplied Cash Payment Income"})
	// Other Income
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 148, Name: "Dividend Income"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 149, Name: "Interest Earned"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 150, Name: "Other Investment Income"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 151, Name: "Other Miscellaneous Income"})
	qtx.SeedLedgerAccountSubType(ctx, accounting_db.SeedLedgerAccountSubTypeParams{LedgerAccountSubTypeID: 152, Name: "Tax Exempt Interest"})

	err = dbTx.Commit(ctx)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Seed()

	if err != nil {
		panic(err)
	}
}
