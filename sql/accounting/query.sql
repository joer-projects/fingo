-- name: AddTransaction :exec
INSERT INTO accounting_transaction (
  id,
  project_id,
  transaction_type_id,
  account_number,
  memo,
  posting_date,
  created_by,
  created_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetTransaction :one
SELECT * FROM accounting_transaction WHERE id = $1;

-- name: GetTransactions :many
SELECT * FROM accounting_transaction;
