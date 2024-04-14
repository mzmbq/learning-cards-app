include .env

DATABASE_URL=postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable
backend = ./backend
frontend = ./frontend


build:
	@go build -C $(backend) -v ./cmd/apiserver

run:
	@set -a; . ./.env; set +a; go run -C $(backend) -v ./cmd/apiserver

run_all:
	@go run -C $(backend) -v ./cmd/apiserver &
	@cd $(frontend) && npm start

clean:
	@cd $(backend) && rm ./apiserver

migrate_up:
	@cd $(backend) && migrate -path=./migrations -database=$(DATABASE_URL) -verbose up

migrate_down:
	@cd $(backend) && migrate -path=./migrations -database=$(DATABASE_URL) -verbose down

db_up:
	@echo "Starting database container..."
	@docker compose up -d

db_down:
	@echo "Stopping database container..."
	@docker compose down


.PHONY: build run run_all clean migrate_up migrate_down db_up db_down

.DEFAULT_GOAL := build