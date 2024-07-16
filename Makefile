include .env

export DATABASE_URL=postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable
export SESSION_KEY

.PHONY: build
build:
	@go build -o ./bin/apiserver -v ./cmd/apiserver

.PHONY: run
run: build
	@./bin/apiserver

.PHONY: run_frontend
run_frontend:
	@cd frontend && npm start

.PHONY: clean
clean:
	@rm ./apiserver

.PHONY: migrate_up
migrate_up:
	@migrate -path=./migrations -database=$(DATABASE_URL) -verbose up

.PHONY: migrate_down
migrate_down:
	@migrate -path=./migrations -database=$(DATABASE_URL) -verbose down

.PHONY: db_up
db_up:
	@echo "Starting database container..."
	@docker compose up -d

.PHONY: db_down
db_down:
	@echo "Stopping database container..."
	@docker compose down

.PHONY: test
test:
	@TESTDB_URL=$(TESTDB_URL) go test ./...



.DEFAULT_GOAL := build