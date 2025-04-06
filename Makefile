# Makefile for envr

BINARY_NAME=envr

.PHONY: all build run clean fmt vet

all: build

build:
	go build -o $(BINARY_NAME) main.go

run: build
	./$(BINARY_NAME) check

clean:
	rm -f $(BINARY_NAME)

fmt:
	go fmt ./...

vet:
	go vet ./...



