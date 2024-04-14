backend = ./backend
frontend = ./frontend


build:
	@go build -C $(backend) -v ./cmd/apiserver

run:
	@go run -C $(backend) -v ./cmd/apiserver &
	@cd $(frontend) && npm start

clean:
	@cd $(backend) && rm ./apiserver

migrate_up:
	@cd $(backend) && migrate -path=./migrations -database=$(DATABASE_URL) -verbose up

migrate_down:
	@cd $(backend) && migrate -path=./migrations -database=$(DATABASE_URL) -verbose down

.PHONY: build run clean migrate_up migrate_down
.DEFAULT_GOAL := build