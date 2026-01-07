run:
    GOEXPERIMENT=greenteagc go run ./cmd/happynum/main.go

build:
    GOEXPERIMENT=greenteagc go build -o ./build/happynum-go ./cmd/happynum/main.go

test:
    go test ./...
