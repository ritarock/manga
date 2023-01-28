.PHONY: test build install

BINDIR=bin

build:
	go build -o $(BINDIR)/manga .

test:
	go test $(shell go list ./... | grep internal)

install:
	go install
