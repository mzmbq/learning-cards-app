backend = ./backend


.PHONY: build
build:
	go build -C $(backend) -v ./cmd/apiserver


.PHONY: run
run:
	go run -C $(backend) -v ./cmd/apiserver


.DEFAULT_GOAL := build