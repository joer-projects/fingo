package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/joer-projects/fingo/cmd/convert"
	"github.com/joer-projects/fingo/internal/domain/entity/transaction"
	"github.com/joer-projects/fingo/internal/infra/db/accounting_db"
	"github.com/joer-projects/fingo/internal/infra/db/connection/postgres"
)

type TransactionPostgresRepo struct {
	context  context.Context
	postgres *pgx.Conn
	queries  *accounting_db.Queries
}

func NewTransactionPostgresRepo(context context.Context) (*TransactionPostgresRepo, error) {
	postgres, err := postgres.NewConnection(context)
	queries := accounting_db.New(postgres)

	if err != nil {
		return nil, err
	}

	return &TransactionPostgresRepo{
		context,
		postgres,
		queries,
	}, nil
}

func (r *TransactionPostgresRepo) Add(t *transaction.Transaction) (transaction.TransactionRaw, error) {
	raw := t.ToRaw()
	r.queries.AddTransaction(r.context, accounting_db.AddTransactionParams{
		ID:                          raw.Id,
		AccountingTransactionTypeID: int32(raw.TransactionTypeId),
		ProjectID:                   convert.StringToPGText(raw.ProjectId),
		Memo:                        convert.StringToPGText(raw.Memo),
		PostingDate:                 convert.TimeToPGTimestamp(&raw.PostingDate),
		CreatedBy:                   raw.CreatedBy,
		CreatedAt:                   convert.TimeToPGTimestamp(&raw.CreatedAt),
	})
	return raw, nil
}

func (r *TransactionPostgresRepo) Save(t *transaction.Transaction) (transaction.TransactionRaw, error) {
	raw := t.ToRaw()

	r.queries.
}
