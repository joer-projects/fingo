package transaction

type transactionLine struct {
	LedgerAccountId   string  `json:"ledger_account_id"`
	BusinessPartnerId string  `json:"business_partner_id"`
	ProjectId         string  `json:"project_id"`
	Credit            float64 `json:"credit"`
	Debit             float64 `json:"debit"`
	Amount            float64 `json:"amount"`
	Description       string  `json:"description"`
}

type TransactionLineNewProps struct {
	LedgerAccountId   string
	BusinessPartnerId string
	ProjectId         string
	Credit            float64
	Debit             float64
	Description       string
}

func NewTransactionLine(props TransactionLineNewProps) (transactionLine, error) {
	return transactionLine{
		LedgerAccountId:   props.LedgerAccountId,
		BusinessPartnerId: props.BusinessPartnerId,
		ProjectId:         props.ProjectId,
		Credit:            props.Credit,
		Debit:             props.Debit,
		Amount:            props.Credit - props.Debit,
		Description:       props.Description,
	}, nil
}
