
BIN_OUTPUT_PATH = bin
TOOL_BIN = bin/gotools/$(shell uname -s)-$(shell uname -m)
UNAME_S ?= $(shell uname -s)
GOPATH = $(HOME)/go/bin
export PATH := ${GOPATH}:$(PATH) 
SHELL := /bin/sh 

build:
	rm -f $(BIN_OUTPUT_PATH)/gomod_cloudbuild
	go build $(LDFLAGS) -o $(BIN_OUTPUT_PATH)/gomod_cloudbuild main.go

module.tar.gz: build
	rm -f $(BIN_OUTPUT_PATH)/module.tar.gz
	tar czf $(BIN_OUTPUT_PATH)/module.tar.gz $(BIN_OUTPUT_PATH)/gomod_cloudbuild

setup: clean format update-rdk

clean:
	rm -rf $(BIN_OUTPUT_PATH)/gomod_cloudbuild $(BIN_OUTPUT_PATH)/module.tar.gz gomod_cloudbuild

format:
	gofmt -w -s .
	go get golang.org/x/tools/cmd/goimports@latest
	find . -name '*.go' -exec goimports -w {} +

update-rdk:
	go get go.viam.com/rdk@latest
	go mod tidy