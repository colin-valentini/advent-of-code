.PHONY: build test fmt vet lint tidy fix

build:
	go build ./...

test:
	go test ./...

fmt:
	gofmt -s -w .

vet:
	go vet ./...

lint:
	golangci-lint run ./...

tidy:
	go mod tidy -v

check: lint vet test

fix: fmt tidy
