GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOSRC=.
BINARY_DIR=bin

all: clean build install

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_DIR)/stonk $(GOSRC)/main.go

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)

.PHONY: test
test:
	$(GOTEST) -v ./test

.PHONY: install
install:
	$(GOINSTALL)
