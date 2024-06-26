package transaction

import (
	"encoding/json"
	"time"
)

type TransactionNewProps struct {
	TypeId           int
	AccountNumber    *string
	Memo             *string
	TransactionLines []TransactionLineCreateProps
	ProjectId        *string
	PostingDate      time.Time
	CreatedBy        string
}

type transactionProps struct {
	TransactionType  transactionType
	ProjectId        *string
	AccountNumber    *string
	Memo             *string
	TransactionLines []transactionLine
	PostingDate      time.Time
	UpdatedBy        *string
	UpdatedAt        *time.Time
	CreatedBy        string
	CreatedAt        time.Time
}

type TransactionRaw struct {
	Id                string            `json:"id"`
	TransactionTypeId int               `json:"transaction_type_id"`
	ProjectId         *string           `json:"project_id"`
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

func New(props TransactionNewProps, id string) (Transaction, error) {
	transactionType, err := createTransactionType(props.TypeId)
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
			ProjectId:        props.ProjectId,
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

func Restore(raw TransactionRaw) Transaction {
	transactionType := restoreTransactionType(transactionType{
		Id:   raw.TransactionTypeId,
		Name: transactionTypeNameMap[raw.TransactionTypeId],
	})
	return Transaction{
		id: raw.Id,
		props: transactionProps{
			TransactionType:  transactionType,
			ProjectId:        raw.ProjectId,
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

func (t Transaction) ToJson() (string, error) {
	data, err := json.MarshalIndent(TransactionRaw{
		Id:                t.id,
		TransactionTypeId: t.props.TransactionType.Id,
		ProjectId:         t.props.ProjectId,
		AccountNumber:     t.props.AccountNumber,
		Memo:              t.props.Memo,
		PostingDate:       t.props.PostingDate,
		TransactionLines:  t.props.TransactionLines,
		UpdatedBy:         t.props.UpdatedBy,
		UpdatedAt:         t.props.UpdatedAt,
		CreatedBy:         t.props.CreatedBy,
		CreatedAt:         t.props.CreatedAt,
	}, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
