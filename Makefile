GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

.PHONY: default
default: test

.PHONY: test
test:
	go test -v ./...

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
