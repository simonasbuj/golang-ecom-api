
run:
	go run ./cmd

# pre-commit
lint:
	golangci-lint run --verbose --max-issues-per-linter=0 --max-same-issues=0

lint-fix:
	golangci-lint run --verbose --fix

.PHONY: test
test:
	go test -v ./...


# database
up:
	goose up

down:
	goose down

sqlc-generate:
	sqlc generate
