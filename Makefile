.DEFAULT_GOAL := all

fmt:
	@go fmt ./...

run:
	@go run main.go

tidy:
	@go mod tidy

all: fmt run

.PHONY: fmt run all
