backend = ./backend


.PHONY: build
build:
	go build -C $(backend) -v ./cmd/apiserver


.PHONY: run
run:
	go run -C $(backend) -v ./cmd/apiserver

.PHONY: clean
clean:
	cd $(backend) && rm ./apiserver

.DEFAULT_GOAL := build 