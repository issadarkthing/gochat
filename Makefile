.PHONY: build gochat server

BIN_DIR := $(CURDIR)/bin

default: build

$(BIN_DIR):
	mkdir -p $@

gochat:
	@cd ./client; \
	go build -v -o $(BIN_DIR)/gochat

server:
	@cd ./server; \
	go build -v -o $(BIN_DIR)/server

clean:
	rm -rf bin

build: gochat server
