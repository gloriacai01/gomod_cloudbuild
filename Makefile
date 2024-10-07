
SOURCE_OS ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
SOURCE_ARCH ?= $(shell uname -m)
TARGET_OS ?= $(SOURCE_OS)
TARGET_ARCH ?= $(SOURCE_ARCH)
BIN_OUTPUT_PATH = bin/$(TARGET_OS)-$(TARGET_ARCH)
TOOL_BIN = bin/gotools/$(shell uname -s)-$(shell uname -m)
UNAME_S ?= $(shell uname -s)
GOPATH = $(HOME)/go/bin
export PATH := ${GOPATH}:$(PATH) 

ifeq ($(TARGET_OS),linux)
	CGO_ENABLED = 1
	CGO_LDFLAGS := -l:libjpeg.a
endif

build: 
	rm -f $(BIN_OUTPUT_PATH)/gomod_cloudbuild
	go build $(LDFLAGS) -o $(BIN_OUTPUT_PATH)/gomod_cloudbuild main.go

module.tar.gz: format update-rdk build
	rm -f $(BIN_OUTPUT_PATH)/module.tar.gz
	tar czf $(BIN_OUTPUT_PATH)/module.tar.gz $(BIN_OUTPUT_PATH)/gomod_cloudbuild

setup: 
	if [ "$(UNAME_S)" = "Linux" ]; then \
			sudo apt install -y libnlopt-dev libjpeg-dev pkg-config; \
	fi

clean:
	rm -rf $(BIN_OUTPUT_PATH)/gomod_cloudbuild $(BIN_OUTPUT_PATH)/module.tar.gz gomod_cloudbuild

format:
	gofmt -w -s .
	go get golang.org/x/tools/cmd/goimports@latest
	which goimports
	find . -name '*.go' -exec goimports -w {} +

update-rdk:
	go get go.viam.com/rdk@latest
	go mod tidy