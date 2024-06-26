.PHONY: build run test 

test: 
	go test ./...

env.up:
	docker-compose up -d

env.down:
	docker-compose down
	docker volume rm fingo_postgres_data
