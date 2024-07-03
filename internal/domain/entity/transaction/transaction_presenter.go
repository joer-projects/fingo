package transaction

import (
	"encoding/json"
	"time"
)

func (t Transaction) ToJson() ([]byte, error) {
	raw := t.ToRaw()
	data, err := json.Marshal(raw)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t Transaction) ToRaw() *TransactionRaw {
	raw := TransactionRaw{
		Id:                t.GetID(),
		TransactionTypeId: t.props.TransactionType.Id,
		ProjectId:         t.props.ProjectId,
		Memo:              t.props.Memo,
		PostingDate:       t.props.PostingDate,
		TransactionLines:  t.props.TransactionLines,
		UpdatedBy:         t.Entity.GetUpdatedBy(),
		UpdatedAt:         t.Entity.GetUpdatedAt(),
		CreatedBy:         t.Entity.GetCreatedBy(),
		CreatedAt:         t.Entity.GetCreatedAt(),
	}
	return &raw
}

type TransactionRaw struct {
	Id                string            `json:"id"`
	TransactionTypeId int               `json:"transaction_type_id"`
	ProjectId         *string           `json:"project_id"`
	Memo              *string           `json:"memo"`
	TransactionLines  []transactionLine `json:"transaction_lines"`
	PostingDate       time.Time         `json:"posting_date"`
	UpdatedBy         *string           `json:"updated_by"`
	UpdatedAt         *time.Time        `json:"updated_at"`
	CreatedBy         string            `json:"created_by"`
	CreatedAt         time.Time         `json:"created_at"`
}
