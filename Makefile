# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BUILD_DIR=./bin
BINARY_NAME=happynum
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_DARWIN=$(BINARY_NAME)_darwin
BINARY_WINDOWS=$(BINARY_NAME)_windows
MODULE=github.com/jjmark15/happynum_go
GIT_TAG=`git describe --tags`
LDFLAGS="-X main.tagVersion=$(GIT_TAG)"

all: test build
install:
	$(GOINSTALL) -v -ldflags $(LDFLAGS) $(MODULE)/cmd/$(BINARY_NAME)
build:
	$(GOBUILD) -ldflags $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_$(GIT_TAG) -v $(MODULE)/cmd/$(BINARY_NAME)
# test:
# 		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BUILD_DIR)/$(BINARY_NAME)*

# Cross compilation
build-linux-amd64:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_LINUX)_amd64_$(GIT_TAG) -v $(MODULE)/cmd/$(BINARY_NAME)
build-linux-arm:
		CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -ldflags $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_LINUX)_arm_$(GIT_TAG) -v $(MODULE)/cmd/$(BINARY_NAME)
build-darwin-amd64:
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_DARWIN)_amd64_$(GIT_TAG) -v $(MODULE)/cmd/$(BINARY_NAME)
build-windows-amd64:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_WINDOWS)_amd64_$(GIT_TAG).exe -v $(MODULE)/cmd/$(BINARY_NAME)

build-all: build-linux-amd64 build-linux-arm build-darwin-amd64 build-windows-amd64