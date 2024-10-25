NAME := mcinstaller
VERSION := 2.1

ifeq ($(OS),Windows_NT)
	BINARY := $(NAME)-$(VERSION)-win.exe
else
	BINARY := $(NAME)-$(VERSION)-linux
endif

build: 
	@echo "Building $(BINARY)..."
	@mkdir -p bin
	go build -o bin/$(BINARY)

clean:
	@echo "Cleaning up..."
	@rm -f bin/$(BINARY)

.PHONY: clean build
