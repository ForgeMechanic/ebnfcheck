# Makefile for building ebnfcheck variants

BINARY=ebnfcheck

.PHONY: all build build-go build-iso build-w3c clean

all: build

build:
	go build -o $(BINARY) ./...

# Build a binary with the default dialect overridden at link time.
build-go:
	go build -ldflags "-X main.dialectFlag=go" -o $(BINARY)-go ./...

build-iso:
	go build -ldflags "-X main.dialectFlag=iso" -o $(BINARY)-iso ./...

build-w3c:
	go build -ldflags "-X main.dialectFlag=w3c" -o $(BINARY)-w3c ./...

build-all: build-go build-iso build-w3c

clean:
	rm -f $(BINARY) $(BINARY)-go $(BINARY)-iso $(BINARY)-w3c
