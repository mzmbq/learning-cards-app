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


.PHONY: migrate_up
migrate_up:
	cd $(backend) && migrate -path=./migrations -database=$(DATABASE_URL) -verbose up
	

.PHONY: migrate_down
migrate_down:
	cd $(backend) && migrate -path=./migrations -database=$(DATABASE_URL) -verbose down


.DEFAULT_GOAL := build 