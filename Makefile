# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_DIR=./bin
BINARY_NAME=happynum_cli
BINARY_LINUX=$(BINARY_NAME)_linux

all: test build
build:
		$(GOBUILD) -ldflags "-X main.tagVersion=`git describe --tags`" -o $(BUILD_DIR)/$(BINARY_NAME) -v ./happynum_cli
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/markbates/goth
		$(GOGET) github.com/markbates/pop


# Cross compilation
build-linux-amd64:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-X main.tagVersion=`git describe --tags`" -o $(BUILD_DIR)/$(BINARY_LINUX)_amd64 -v ./happynum_cli
build-linux-arm64:
		CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -ldflags "-X main.tagVersion=`git describe --tags`" -o $(BUILD_DIR)/$(BINARY_LINUX)_arm64 -v ./happynum_cli

build-all: build-linux-amd64 build-linux-arm64