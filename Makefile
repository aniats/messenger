include configs/.env

build:
	go build -o bin/messenger ./cmd/messenger

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

test:
	go test ./...

generate:
	go generate ./...
