.PHONY: build run test 

test: 
	go test ./...

env.up:
	docker-compose up -d

env.down:
	docker-compose down
	docker volume rm fingo_postgres_data

seed:
	go run ./internal/infra/seed/accounting/accounting_seed.go

sqlc.gen:
	docker run --rm -v $(shell pwd):/src -w /src sqlc/sqlc generate
