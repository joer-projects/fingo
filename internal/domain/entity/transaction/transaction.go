package transaction

import (
	"time"
)

type NewProps struct {
	ProjectID        *string
	TypeID           int32
	AccountNumber    *string
	Memo             *string
	TransactionLines []TransactionLineCreateProps
	PostingDate      time.Time
	CreatedBy        string
}

type transactionProps struct {
	TransactionType  transactionType
	ProjectID        *string
	AccountNumber    *string
	Memo             *string
	TransactionLines []transactionLine
	PostingDate      time.Time
	UpdatedBy        *string
	UpdatedAt        *time.Time
	CreatedBy        string
	CreatedAt        time.Time
}

type transactionRaw struct {
	ID                string            `json:"id"`
	TransactionTypeID int32             `json:"transaction_type_id"`
	ProjectID         *string           `json:"project_id"`
	AccountNumber     *string           `json:"account_number"`
	Memo              *string           `json:"memo"`
	TransactionLines  []transactionLine `json:"transaction_lines"`
	PostingDate       time.Time         `json:"posting_date"`
	UpdatedBy         *string           `json:"updated_by"`
	UpdatedAt         *time.Time        `json:"updated_at"`
	CreatedBy         string            `json:"created_by"`
	CreatedAt         time.Time         `json:"created_at"`
}

type Transaction struct {
	id    string
	props transactionProps
}

func New(props NewProps, id string) (Transaction, error) {
	transactionType, err := createTransactionType(props.TypeID)
	if err != nil {
		return Transaction{}, err
	}

	if props.PostingDate.IsZero() {
		return Transaction{}, ErrInvalidPostingDate
	}

	if props.CreatedBy == "" {
		return Transaction{}, ErrCreatedByIsRequired
	}

	transactionLines := make([]transactionLine, len(props.TransactionLines))
	for index, lineProps := range props.TransactionLines {
		line, err := NewLine(lineProps)
		if err != nil {
			return Transaction{}, err
		}
		transactionLines[index] = line
	}
	new := Transaction{
		id: id,
		props: transactionProps{
			TransactionType:  transactionType,
			ProjectID:        props.ProjectID,
			AccountNumber:    props.AccountNumber,
			Memo:             props.Memo,
			TransactionLines: transactionLines,
			PostingDate:      props.PostingDate,
			CreatedBy:        props.CreatedBy,
			CreatedAt:        time.Now(),
		},
	}
	return new, nil
}

func Restore(raw transactionRaw) Transaction {
	transactionType := restoreTransactionType(transactionType{
		Id:   raw.TransactionTypeID,
		Name: transactionTypeNameMap[raw.TransactionTypeID],
	})
	return Transaction{
		id: raw.ID,
		props: transactionProps{
			TransactionType:  transactionType,
			ProjectID:        raw.ProjectID,
			AccountNumber:    raw.AccountNumber,
			Memo:             raw.Memo,
			PostingDate:      raw.PostingDate,
			TransactionLines: raw.TransactionLines,
			UpdatedBy:        raw.UpdatedBy,
			UpdatedAt:        raw.UpdatedAt,
			CreatedBy:        raw.CreatedBy,
			CreatedAt:        raw.CreatedAt,
		},
	}
}

func (t Transaction) Id() string {
	return t.id
}
