GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

.PHONY: default
default: test

.PHONY: gen-docs
gen-docs: 
	@-rm -rf docs/public
	@cd docs && hugo --minify  --baseURL "https://alimy.me/yesql/" && cd -

.PHONY: run-docs
run-docs: 
	@cd docs && hugo serve --minify && cd -

.PHONY: test
test:
	go test -v ./...

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
