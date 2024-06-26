package schema

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TransactionTypeId int    `json:"transaction_type_id"`
	ProjectId         string `json:"project_id"`
	AccountNumber     string `json:"account_number"`
	Memo              string `json:"memo"`
	TransactionLines  []TransactionLine
	PostingDate       string `json:"posting_date"`
	UpdatedBy         string `json:"updated_by"`
	UpdatedAt         string `json:"updated_at"`
	CreatedBy         string `json:"created_by"`
	CreatedAt         string `json:"created_at"`
}

type TransactionLine struct {
	gorm.Model
	LedgerAccountId   string  `json:"ledger_account_id"`
	BusinessPartnerId string  `json:"business_partner_id"`
	ProjectId         string  `json:"project_id"`
	Credit            float64 `json:"credit"`
	Debit             float64 `json:"debit"`
	Amount            float64 `json:"amount"`
	Description       string  `json:"description"`
}
