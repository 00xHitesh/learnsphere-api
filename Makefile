.PHONY: run build test lint sqlc migrate-up migrate-down docker-up docker-down

run:
	go run ./cmd/api

build:
	go build -o bin/api ./cmd/api

test:
	go test ./...

test-v:
	go test -v ./...

test-pkg:
	go test -v ./$(pkg)/...

lint:
	golangci-lint run ./...

sqlc:
	sqlc generate

docker-up:
	docker compose up -d

docker-down:
	docker compose down

migrate-up:
	migrate -path db/migrations -database "$$DATABASE_URL" up

migrate-down:
	migrate -path db/migrations -database "$$DATABASE_URL" down 1

tidy:
	go mod tidy
