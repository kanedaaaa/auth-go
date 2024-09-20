build:
	@go build -o bin/auth-go cmd/main.go

test:
	@go test -v ./...

run:
	@./bin/auth-go