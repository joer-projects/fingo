package repo

import (
	"context"

	"github.com/joer-projects/fingo/cmd/convert"
	"github.com/joer-projects/fingo/internal/domain/entity/transaction"
	"github.com/joer-projects/fingo/internal/infra/db/accounting_db"
)

type TransactionPostgresRepo struct {
	context context.Context
	queries *accounting_db.Queries
}

func NewTransactionPostgresRepo(context context.Context, queries *accounting_db.Queries) (*TransactionPostgresRepo, error) {
	return &TransactionPostgresRepo{
		context,
		queries,
	}, nil
}

func (r *TransactionPostgresRepo) Add(input *transaction.Transaction) (*transaction.TransactionRaw, error) {
	raw := input.ToRaw()

	err := r.queries.AddTransaction(r.context, accounting_db.AddTransactionParams{
		ID:                          raw.Id,
		AccountingTransactionTypeID: int32(raw.TransactionTypeId),
		ProjectID:                   convert.StringToPGText(raw.ProjectId),
		Memo:                        convert.StringToPGText(raw.Memo),
		PostingDate:                 convert.TimeToPGTimestamp(&raw.PostingDate),
		CreatedBy:                   raw.CreatedBy,
		CreatedAt:                   convert.TimeToPGTimestamp(&raw.CreatedAt),
	})

	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (r *TransactionPostgresRepo) Save(input *transaction.Transaction) (*transaction.TransactionRaw, error) {
	raw := input.ToRaw()

	err := r.queries.SaveTransaction(r.context, accounting_db.SaveTransactionParams{
		ID:                          raw.Id,
		AccountingTransactionTypeID: int32(raw.TransactionTypeId),
		ProjectID:                   convert.StringToPGText(raw.ProjectId),
		Memo:                        convert.StringToPGText(raw.Memo),
		PostingDate:                 convert.TimeToPGTimestamp(&raw.PostingDate),
		UpdatedBy:                   convert.StringToPGText(raw.UpdatedBy),
		UpdatedAt:                   convert.TimeToPGTimestamp(raw.UpdatedAt),
		CreatedBy:                   raw.CreatedBy,
		CreatedAt:                   convert.TimeToPGTimestamp(&raw.CreatedAt),
	})

	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (r *TransactionPostgresRepo) Get(id string) (*transaction.Transaction, error) {
	raw, err := r.queries.GetTransaction(r.context, id)

	if err != nil {
		return nil, err
	}

	txn := transaction.Restore(transaction.TransactionRaw{
		Id:                raw.ID,
		TransactionTypeId: int(raw.AccountingTransactionTypeID),
		ProjectId:         &raw.ProjectID.String,
		Memo:              &raw.Memo.String,
		PostingDate:       raw.PostingDate.Time,
		UpdatedBy:         &raw.UpdatedBy.String,
		UpdatedAt:         &raw.UpdatedAt.Time,
		CreatedBy:         raw.CreatedBy,
		CreatedAt:         raw.CreatedAt.Time,
	})

	return &txn, nil
}
