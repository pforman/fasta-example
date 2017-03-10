all: build test coverage live-test

build: fasta-example

fasta-example: cmd/fasta-example/main.go *.go
	go build -o fasta-example ./cmd/fasta-example

install:
	go install ./...

test:
	go test

coverage:
	go test -coverprofile=c.out
	go tool cover -func=c.out

live-test:
	./scripts/test-fasta.sh

annotate: coverage
	gocov convert c.out | gocov annotate -
