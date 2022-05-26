.PHONY: all build clean install uninstall fmt simplify check run test

build:
	@go build ./...

test:
	@go test ./test/...