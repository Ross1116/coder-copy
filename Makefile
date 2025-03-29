export CGO_ENABLED=1

.PHONY: build run clean

build:
	go build -o bin/coder-copy cmd/app/main.go

run:
	go run cmd/app/main.go

clean:
	rm -f coder-copy
