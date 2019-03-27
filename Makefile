# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BUILD_DIR=./bin
BINARY_NAME=happynum
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_DARWIN=$(BINARY_NAME)_darwin
BINARY_WINDOWS=$(BINARY_NAME)_windows
GIT_TAG=`git describe --tags`
PACKAGE_PATH=github.com/jjmark15/happynum_go

all: test build
build:
	$(GOBUILD) -ldflags "-X $(PACKAGE_PATH)/internal/pkgcli.tagVersion=$(GIT_TAG)" -o $(BUILD_DIR)/$(BINARY_NAME)_$(GIT_TAG) -v ./cmd/$(BINARY_NAME)
# test:
# 		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BUILD_DIR)/$(BINARY_NAME)*
# deps:
# 		$(GOGET) github.com/jjmark15/


# Cross compilation
build-linux-amd64:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-X $(PACKAGE_PATH)/internal/pkgcli.tagVersion=$(GIT_TAG)" -o $(BUILD_DIR)/$(BINARY_LINUX)_amd64_$(GIT_TAG) -v ./cmd/$(BINARY_NAME)
build-linux-arm:
		CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -ldflags "-X $(PACKAGE_PATH)/internal/pkgcli.tagVersion=$(GIT_TAG)" -o $(BUILD_DIR)/$(BINARY_LINUX)_arm_$(GIT_TAG) -v ./cmd/$(BINARY_NAME)
build-darwin-amd64:
		CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -ldflags "-X $(PACKAGE_PATH)/internal/pkgcli.tagVersion=$(GIT_TAG)" -o $(BUILD_DIR)/$(BINARY_DARWIN)_amd64_$(GIT_TAG) -v ./cmd/$(BINARY_NAME)
build-windows-amd64:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -ldflags "-X $(PACKAGE_PATH)/internal/pkgcli.tagVersion=$(GIT_TAG)" -o $(BUILD_DIR)/$(BINARY_WINDOWS)_amd64_$(GIT_TAG).exe -v ./cmd/$(BINARY_NAME)

build-all: build-linux-amd64 build-linux-arm build-darwin-amd64 build-windows-amd64