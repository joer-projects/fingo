package transaction

import "errors"

var (
	ErrInvalidPostingDate     = errors.New("invalid posting date")
	ErrInvalidTransactionType = errors.New("invalid transaction type")
	ErrCreatedByIsRequired    = errors.New("created by is required")
)
