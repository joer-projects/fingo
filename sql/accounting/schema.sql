CREATE TABLE "accounting_transaction_type" (
  "id" INTEGER PRIMARY KEY,
  "name" VARCHAR(56) NOT NULL
);

CREATE TABLE "accounting_transaction" (
  "id" VARCHAR(36) PRIMARY KEY,
  "project_id" VARCHAR(36),
  "transaction_type_id" INTEGER NOT NULL,
  "account_number" VARCHAR(100),
  "memo" VARCHAR(255),
  "posting_date" TIMESTAMPTZ(6) NOT NULL,
  "updated_by" VARCHAR(36),
  "updated_at" TIMESTAMP(6),
  "created_by" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP(6) NOT NULL,

  CONSTRAINT "accounting_transaction_to_accounting_transaction_type_fkey" 
  FOREIGN KEY ("transaction_type_id") 
  REFERENCES "accounting_transaction_type" ("id") 
  ON DELETE CASCADE ON UPDATE CASCADE
);
