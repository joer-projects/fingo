package transaction

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	t.Run("should create a new transaction", func(t *testing.T) {

		id, err := uuid.NewV7()

		if err != nil {
			t.Error(err)
		}

		AccountNumber := "1234567890"
		Memo := "test"
		ProjectId := "123"
		PostingDate := time.Now()
		CreatedBy := "test"

		txn, err := NewTransaction(TransactionNewProps{
			TypeId:        1,
			AccountNumber: &AccountNumber,
			Memo:          &Memo,
			ProjectId:     &ProjectId,
			PostingDate:   PostingDate,
			CreatedBy:     CreatedBy,
			TransactionLines: []TransactionLineNewProps{
				{
					LedgerAccountId:   "1234567890",
					BusinessPartnerId: "1234567890",
					Credit:            100.0,
					Debit:             0.0,
					ProjectId:         "123",
					Description:       "test",
				},
				{
					LedgerAccountId:   "1234567890",
					BusinessPartnerId: "1234567890",
					Credit:            0.0,
					Debit:             100.0,
					ProjectId:         "123",
					Description:       "test",
				},
			},
		}, id.String())

		assert.NoError(t, err)
		assert.Equal(t, id.String(), txn.GetID())
		assert.Equal(t, 1, txn.props.TransactionType.Id)
		assert.Equal(t, AccountNumber, *txn.props.AccountCode)
		assert.Equal(t, Memo, *txn.props.Memo)
		assert.Equal(t, ProjectId, *txn.props.ProjectId)
		assert.Equal(t, PostingDate, txn.props.PostingDate)
		assert.Equal(t, CreatedBy, txn.Entity.GetCreatedBy())
		assert.Equal(t, 2, len(txn.props.TransactionLines))
		assert.Equal(t, "1234567890", txn.props.TransactionLines[0].LedgerAccountId)
		assert.Equal(t, "1234567890", txn.props.TransactionLines[0].BusinessPartnerId)
		assert.Equal(t, 100.0, txn.props.TransactionLines[0].Credit)
		assert.Equal(t, 0.0, txn.props.TransactionLines[0].Debit)
	})
}
