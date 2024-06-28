.PHONY: build run test

test:
	go test ./...

env.up:
	docker-compose up -d
	sleep 1
	sqlc generate
	sleep 1
	psql postgresql://admin:admin@localhost:5432/accounting -f sql/accounting/schema.sql
	sleep 1
	go run seed/main.go

env.down:
	docker-compose down
	docker volume rm fingo_postgres_data