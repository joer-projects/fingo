-- pgFormatter-ignore

CREATE TABLE "accounting_transaction_type"(
  "accounting_transaction_type_id" integer PRIMARY KEY,
  "name" varchar(56) NOT NULL
);

CREATE TABLE "accounting_transaction"(
  "id" varchar(36) PRIMARY KEY,
  "project_id" varchar(36),
  "accounting_transaction_type_id" integer NOT NULL,
  "memo" varchar(255),
  "posting_date" timestamptz(6) NOT NULL,
  "updated_by" varchar(36),
  "updated_at" timestamptz(6),
  "created_by" varchar(255) NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  CONSTRAINT "accounting_transaction_to_accounting_transaction_type_fkey" FOREIGN KEY ("accounting_transaction_type_id") REFERENCES "accounting_transaction_type"("accounting_transaction_type_id")
);

CREATE TABLE "ledger_account_class"(
  "ledger_account_class_id" integer PRIMARY KEY,
  "name" varchar(100) NOT NULL UNIQUE
);

CREATE TABLE "ledger_account_type"(
  "ledger_account_type_id" integer PRIMARY KEY,
  "name" varchar(100) NOT NULL UNIQUE
);

CREATE TABLE "ledger_account_sub_type"(
  "ledger_account_sub_type_id" integer PRIMARY KEY,
  "name" varchar(100) NOT NULL UNIQUE
);

CREATE TABLE "ledger_account"(
  "ledger_account_id" varchar(36) PRIMARY KEY,
  "description" varchar(500),
  "account_code" varchar(56),
  "tax_code" varchar(56),
  "name" varchar(100) NOT NULL,
  "ledger_account_class_id" integer NOT NULL,
  "ledger_account_type_id" integer NOT NULL,
  "ledger_account_sub_type_id" integer NOT NULL,
  "ledger_account_level" integer NOT NULL,
  "ledger_account_parent_id" varchar(36),
  "updated_by" varchar(36),
  "updated_at" timestamptz(6),
  "created_by" varchar(36) NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  CONSTRAINT "ledger_account_to_ledger_account_class_fkey" FOREIGN KEY ("ledger_account_class_id") REFERENCES "ledger_account_class"("ledger_account_class_id") ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT "ledger_account_to_ledger_account_type_fkey" FOREIGN KEY ("ledger_account_type_id") REFERENCES "ledger_account_type"("ledger_account_type_id") ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT "ledger_account_to_ledger_account_sub_type_fkey" FOREIGN KEY ("ledger_account_sub_type_id") REFERENCES "ledger_account_sub_type"("ledger_account_sub_type_id") ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT "ledger_account_to_ledger_account_parent_fkey" FOREIGN KEY ("ledger_account_parent_id") REFERENCES "ledger_account"("ledger_account_id") ON DELETE CASCADE ON UPDATE CASCADE
);
