package transaction

const (
	journal_entry = iota + 1
	vendor_payment
	incoming_payment
	deposit
	ap_invoice
	ap_down_payment
	ap_credit_memo
	ar_invoice
	ar_down_payment
	ar_credit_memo
)

var transactionTypeNameMap = map[int32]string{
	journal_entry:    "Journal Entry",
	vendor_payment:   "Vendor Payment",
	incoming_payment: "Incoming Payment",
	deposit:          "Deposit",
	ap_invoice:       "AP Invoice",
	ap_down_payment:  "AP Down Payment",
	ap_credit_memo:   "AP Credit Memo",
	ar_invoice:       "AR Invoice",
	ar_down_payment:  "AR Down Payment",
	ar_credit_memo:   "AR Credit Memo",
}

type transactionType struct {
	Id   int32
	Name string
}

func createTransactionType(id int32) (transactionType, error) {
	typeName := transactionTypeNameMap[id]

	if typeName == "" {
		return transactionType{}, ErrInvalidTransactionType
	}

	return transactionType{
		Id:   id,
		Name: typeName,
	}, nil
}

func restoreTransactionType(raw transactionType) transactionType {
	return transactionType{
		Id:   raw.Id,
		Name: raw.Name,
	}
}
