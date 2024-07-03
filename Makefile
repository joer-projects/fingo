# GNU Make. https://www.gnu.org/software/make/


# Build & Setup

.PHONY: setup
setup:
	@make deps.get \
	&& make build \
	&& make db.setup

.PHONY: deps.get
deps.get:
	@go mod download

.PHONY: build
build:
	@go build -o bin/fingo src/main.go

.PHONY: run
run:
	@go run src/main.go


# Code quality

.PHONY: fmt
fmt:
	@go fmt ./**/*.go

.PHONY: test
test:
	@go test -race -failfast ./...

test.watch:
	@find src -regex '.*\.go' | entr -cr make test


# Codegen

# Usage: `make gen.migration name=<migration name>`
.PHONY: gen.migration
gen.migration:
	@migrate create \
		-ext sql \
		-dir src/pg/migrations \
		-seq \
		"$(name)"


# Postgres

export PG_HOST    ?= localhost
export PG_PORT    ?= 5432
export PG_USER    ?= fingo
export PG_PASS    ?= fingo
export PG_NAME    ?= fingo
export PG_SSLMODE ?= disable

.PHONY: db.setup
db.setup:
	@make db.create \
	&& make db.migrate \
	&& make db.seed

.PHONY: db.create
db.create:
	@createdb -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) $(PG_NAME)

.PHONY: db.migrate
db.migrate:
	@migrate \
		-source file://src/pg/migrations \
		-database 'postgres://$(PG_USER):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_NAME)?sslmode=${PG_SSLMODE}' \
		up

.PHONY: db.seed
db.seed:
	@psql -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) -d $(PG_NAME) -f src/pg/seed.sql

.PHONY: db.drop
db.drop:
	@dropdb -h $(PG_HOST) -p $(PG_PORT) -U $(PG_USER) $(PG_NAME)

.PHONY: db.reset
db.reset:
	@make db.drop \
	&& make db.setup
