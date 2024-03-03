backend = ./backend
frontend = ./frontend


.PHONY: build
build:
	go build -C $(backend) -v ./cmd/apiserver


.PHONY: run
run:
	go run -C $(backend) -v ./cmd/apiserver &
	cd $(frontend) && npm start


.PHONY: clean
clean:
	cd $(backend) && rm ./apiserver

.DEFAULT_GOAL := build 