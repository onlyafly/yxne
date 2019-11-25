all: run

run: install
	yxne

fmt:
	go fmt ./...

test: fmt
	go test ./...

install: fmt
	go install ./...