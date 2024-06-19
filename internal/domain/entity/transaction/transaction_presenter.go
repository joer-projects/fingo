package transaction

import (
	"encoding/json"
)

func (t Transaction) ToRaw() transactionRaw {
	return transactionRaw{
		ID:                t.id,
		ProjectID:         t.props.ProjectID,
		TransactionTypeID: t.props.TransactionType.Id,
		AccountNumber:     t.props.AccountNumber,
		Memo:              t.props.Memo,
		PostingDate:       t.props.PostingDate,
		UpdatedBy:         t.props.UpdatedBy,
		UpdatedAt:         t.props.UpdatedAt,
		CreatedBy:         t.props.CreatedBy,
		CreatedAt:         t.props.CreatedAt,
	}
}

func (t Transaction) ToJson() (string, error) {
	data, err := json.Marshal(transactionRaw{
		ID:                t.id,
		TransactionTypeID: t.props.TransactionType.Id,
		ProjectID:         t.props.ProjectID,
		AccountNumber:     t.props.AccountNumber,
		Memo:              t.props.Memo,
		PostingDate:       t.props.PostingDate,
		TransactionLines:  t.props.TransactionLines,
		UpdatedBy:         t.props.UpdatedBy,
		UpdatedAt:         t.props.UpdatedAt,
		CreatedBy:         t.props.CreatedBy,
		CreatedAt:         t.props.CreatedAt,
	})
	if err != nil {
		return "", err
	}
	return string(data), nil
}
