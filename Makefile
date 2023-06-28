.PHONY: test build install

BINDIR=bin

build:
	go build -o $(BINDIR)/manga .

test:
	go test ./...

install:
	go install
