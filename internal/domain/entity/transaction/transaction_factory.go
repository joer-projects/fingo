package transaction

import (
	"time"

	"github.com/joer-projects/fingo/core"
)

type TransactionNewProps struct {
	TypeId           int
	AccountNumber    *string
	Memo             *string
	TransactionLines []TransactionLineNewProps
	ProjectId        *string
	PostingDate      time.Time
	CreatedBy        string
}

func NewTransaction(props TransactionNewProps, id string) (Transaction, error) {
	transactionType, err := createTransactionType(props.TypeId)
	if err != nil {
		return Transaction{}, err
	}

	transactionLines := make([]transactionLine, len(props.TransactionLines))
	for index, lineProps := range props.TransactionLines {
		line, err := NewTransactionLine(lineProps)
		if err != nil {
			return Transaction{}, err
		}
		transactionLines[index] = line
	}

	txn := Transaction{
		Entity: core.NewEntity(core.NewEntityProps{
			ID:        id,
			CreatedBy: props.CreatedBy,
			CreatedAt: time.Now(),
		}),
		props: transactionProps{
			TransactionType:  transactionType,
			ProjectId:        props.ProjectId,
			AccountCode:      props.AccountNumber,
			Memo:             props.Memo,
			TransactionLines: transactionLines,
			PostingDate:      props.PostingDate,
		},
	}

	return txn, nil
}

func Restore(raw TransactionRaw) Transaction {
	transactionType := restoreTransactionType(transactionType{
		Id:   raw.TransactionTypeId,
		Name: transactionTypeNameMap[raw.TransactionTypeId],
	})
	return Transaction{
		Entity: core.NewEntity(core.NewEntityProps{
			ID:        raw.Id,
			UpdatedBy: raw.UpdatedBy,
			UpdatedAt: raw.UpdatedAt,
			CreatedBy: raw.CreatedBy,
			CreatedAt: raw.CreatedAt,
		}),
		props: transactionProps{
			ProjectId:        raw.ProjectId,
			TransactionType:  transactionType,
			AccountCode:      raw.AccountCode,
			Memo:             raw.Memo,
			PostingDate:      raw.PostingDate,
			TransactionLines: raw.TransactionLines,
		},
	}
}
