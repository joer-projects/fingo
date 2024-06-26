package transaction

import (
	"time"

	"github.com/joer-projects/fingo/core"
)

type transactionProps struct {
	TransactionType  transactionType
	ProjectId        *string
	AccountCode      *string
	Memo             *string
	TransactionLines []transactionLine
	PostingDate      time.Time
}

type Transaction struct {
	core.Entity
	props transactionProps
}
