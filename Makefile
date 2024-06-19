.PHONY: build run test 

test: 
	go test ./...

docker.up:
	docker-compose up -d

docker.down:
	docker-compose down

seed.accounting:
	go run ./internal/infra/seed/accounting/accounting_seed.go

sqlc.gen:
	docker run --rm -v $(shell pwd):/src -w /src sqlc/sqlc generate
