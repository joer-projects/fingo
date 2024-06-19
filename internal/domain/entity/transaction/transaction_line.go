package transaction

type transactionLine struct {
	LedgerAccountID   string  `json:"ledger_account_id"`
	BusinessPartnerID string  `json:"business_partner_id"`
	ProjectID         string  `json:"project_id"`
	Credit            float64 `json:"credit"`
	Debit             float64 `json:"debit"`
	Amount            float64 `json:"amount"`
	Description       string  `json:"description"`
}

type TransactionLineCreateProps struct {
	LedgerAccountID   string
	BusinessPartnerID string
	ProjectID         string
	Credit            float64
	Debit             float64
	Description       string
}

func NewLine(props TransactionLineCreateProps) (transactionLine, error) {
	return transactionLine{
		LedgerAccountID:   props.LedgerAccountID,
		BusinessPartnerID: props.BusinessPartnerID,
		ProjectID:         props.ProjectID,
		Credit:            props.Credit,
		Debit:             props.Debit,
		Amount:            props.Credit - props.Debit,
		Description:       props.Description,
	}, nil
}
