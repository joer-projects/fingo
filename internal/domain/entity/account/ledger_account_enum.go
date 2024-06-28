package ledger_account

import "fmt"

type ledgerAccountClass struct {
	ASSET     int
	EQUITY    int
	EXPENSE   int
	LIABILITY int
	REVENUE   int
}

func (l *ledgerAccountClass) Name(i int) (string, error) {
	if i < 0 || i > 4 {
		return "", fmt.Errorf("LedgerAccountClassEnum - invalid index")
	}
	return [...]string{"Asset", "Equity", "Expense", "Liability", "Revenue"}[i-1], nil
}

func (l *ledgerAccountClass) ID(i int) (int, error) {
	if i < 0 || i > 4 {
		return 0, fmt.Errorf("LedgerAccountClassEnum - invalid index")
	}
	return [...]int{asset, equity, expense, liability, revenue}[i-1], nil
}

func (l *ledgerAccountClass) Len() int {
	return l.REVENUE - 1
}

const (
	asset = iota + 1
	equity
	expense
	liability
	revenue
)

var LedgerAccountClassEnum = ledgerAccountClass{
	asset,
	equity,
	expense,
	liability,
	revenue,
}
