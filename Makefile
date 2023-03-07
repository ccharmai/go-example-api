.DEFAULT_GOAL := all

fmt:
	@go fmt ./...

run:
	@go run main.go

all: fmt run

.PHONY: fmt run all
