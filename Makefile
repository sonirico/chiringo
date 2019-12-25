.PHONY: run test clean

run:
	go run *.go

test:
	go test ./...

format:
	go fmt ./...

clean:
	go clean -modcache
	go clean -testcache

