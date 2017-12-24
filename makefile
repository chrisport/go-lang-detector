all: test

deps:
	godep save ./...

test:
	godep go test ./...
