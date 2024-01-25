GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)

.PHONY: test
test:
	go test ./...
