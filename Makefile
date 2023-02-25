.PHONY: test build install

BINDIR=bin

gqlgen:
	go run github.com/99designs/gqlgen

build:
	go build -o $(BINDIR)/manga .

test:
	go test $(shell go list ./... | grep internal)

install:
	go install
