package ledger_account

import "github.com/joer-projects/fingo/core"

type LedgerAccountName struct {
	Value string
}

type LedgerAccountProps struct {
	Name           LedgerAccountName
	Classification LedgerAccountClassification
	AccountCode    *string
	TaxCode        *string
}

type LedgerAccount struct {
	core.Entity
	props LedgerAccountProps
}
