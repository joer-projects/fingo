-- pgFormatter-ignore

-- name: SeedTransactionType :exec
INSERT INTO accounting_transaction_type(accounting_transaction_type_id, name) VALUES ($1, $2);

-- name: SeedLedgerAccountClass :exec
INSERT INTO ledger_account_class(ledger_account_class_id, name) VALUES ($1, $2);

-- name: SeedLedgerAccountType :exec
INSERT INTO ledger_account_type(ledger_account_type_id, name) VALUES ($1, $2);

-- name: SeedLedgerAccountSubType :exec
INSERT INTO ledger_account_sub_type(ledger_account_sub_type_id, name) VALUES ($1, $2);

-- name: SeedLedgerAccount :exec
INSERT INTO ledger_account(
  ledger_account_id, 
  description,
  account_code,
  tax_code,
  name,
  ledger_account_class_id,
  ledger_account_type_id,
  ledger_account_sub_type_id,
  ledger_account_level,
  ledger_account_parent_id,
  updated_by,
  updated_at,
  created_by,
  created_at
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);

-- name: AddTransaction :exec
INSERT INTO accounting_transaction(
  id, 
  project_id, 
  accounting_transaction_type_id,
  memo, 
  posting_date, 
  created_by, 
  created_at
  ) VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: SaveTransaction :exec
INSERT INTO accounting_transaction(
  id,
  project_id,
  accounting_transaction_type_id,
  memo,
  posting_date,
  updated_by,
  updated_at,
  created_by,
  created_at
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9);

-- name: GetTransaction :one
SELECT * FROM accounting_transaction WHERE id = $1;

-- name: GetTransactions :many
SELECT * FROM accounting_transaction;

