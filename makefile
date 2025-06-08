GO=go
BUILD=build
GOFILES=*.go
TARGET=./bin/htmlint

all: build

build:
	$(GO) $(BUILD) -o $(TARGET) $(GOFILES)
