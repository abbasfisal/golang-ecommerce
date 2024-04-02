.DEFAULT_GOAL=run
build:
	@go build -o bin/ecomm cmd/main.go

test:
	@go test -v ./..

run:
	go run cmd/main.go