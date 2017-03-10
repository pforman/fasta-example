all: build test coverage

build: fasta-example

fasta-example: cmd/main.go *.go
	go build -o bin/fasta-example ./cmd

install:
	go install ./cmd/...

test:
	go test

coverage:
	go test -coverprofile=c.out
	go tool cover -func=c.out

coverage-report: coverage
	gocov convert c.out | gocov annotate -
