package transaction

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	id := uuid.New()
	t.Run("should create a new transaction", func(t *testing.T) {

		AccountNumber := "1234567890"
		Memo := "test"
		ProjectId := "123"
		PostingDate := time.Now()
		CreatedBy := "test"

		tx, err := New(NewProps{
			TypeID:        1,
			AccountNumber: &AccountNumber,
			Memo:          &Memo,
			ProjectID:     &ProjectId,
			PostingDate:   PostingDate,
			CreatedBy:     CreatedBy,
			TransactionLines: []TransactionLineCreateProps{
				{
					LedgerAccountID:   "1234567890",
					BusinessPartnerID: "1234567890",
					Credit:            100.0,
					Debit:             0.0,
					ProjectID:         "123",
					Description:       "test",
				},
				{
					LedgerAccountID:   "1234567890",
					BusinessPartnerID: "1234567890",
					Credit:            0.0,
					Debit:             100.0,
					ProjectID:         "123",
					Description:       "test",
				},
			},
		}, id.String())

		assert.NoError(t, err)
		assert.Equal(t, id.String(), tx.Id())
		assert.Equal(t, 1, tx.props.TransactionType.Id)
		assert.Equal(t, AccountNumber, *tx.props.AccountNumber)
		assert.Equal(t, Memo, *tx.props.Memo)
		assert.Equal(t, ProjectId, *tx.props.ProjectID)
		assert.Equal(t, PostingDate, tx.props.PostingDate)
		assert.Equal(t, CreatedBy, tx.props.CreatedBy)
		assert.Equal(t, 2, len(tx.props.TransactionLines))
		assert.Equal(t, "1234567890", tx.props.TransactionLines[0].LedgerAccountID)
		assert.Equal(t, "1234567890", tx.props.TransactionLines[0].BusinessPartnerID)
		assert.Equal(t, 100.0, tx.props.TransactionLines[0].Credit)
		assert.Equal(t, 0.0, tx.props.TransactionLines[0].Debit)
	})
}
