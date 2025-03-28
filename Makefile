export CGO_ENABLED=1

.PHONY: build run clean

build:
	go build -o bin/copy-comment-remover cmd/app/main.go

run:
	go run cmd/app/main.go

clean:
	rm -f copy-comment-remover
