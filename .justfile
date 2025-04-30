run:
    go run ./cmd/happynum/main.go

build:
    go build -o ./build/happynum-go ./cmd/happynum/main.go

test:
    go test ./...
