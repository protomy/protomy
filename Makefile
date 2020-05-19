export GOOS=$(shell go env GOOS)
export GO_BUILD=env GO11MODULE=on go build -ldflags="-s -w"
export GO_INSTALL=env GO11MODULE=on go install
export GO_TEST=env GOTRACEBACK=all GO11MODULE=on go test -race -coverprofile=coverage.txt -covermode=atomic
export GO_VET=env GO11MODULE=on go vet
export GO_RUN=env GO11MODULE=on go run
export PATH := $(PWD)/bin/$(GOOS):$(PATH)

VERSION := $(shell cat ./VERSION)

SOURCES := $(shell find . -name '*.go' -not -name '*_test.go') go.mod go.sum
SOURCES_NO_VENDOR := $(shell find . -path ./vendor -prune -o -name "*.go" -not -name '*_test.go' -print)

all: install

ci: vet test build

install: clean vet test build

bench:
	$(GO_TEST) -bench=. -run=^$$ ./...

build: $(SOURCES)
	$(GO_BUILD) -o bin/protomy cmd/protomy/main.go

changelog:
	bundle exec github_changelog_generator -u protomy -p protomy \
		--header-label "# Protomy Changelog" \
		--no-issues-wo-labels \
		--no-pr-wo-labels \
		--release-branch master \
		--enhancement-label Enhancements \
		--bugs-label Fixes \
        --include-labels enhancement,bug,security,breaking,deprecated,removed

release-notes:
	bundle exec github_changelog_generator -u protomy -p protomy \
		--header-label "" \
		--no-unreleased \
		--no-issues-wo-labels \
		--no-pr-wo-labels \
		--release-branch master \
		--enhancement-label Enhancements \
		--bugs-label Fixes \
		--include-labels enhancement,bug,security,breaking,deprecated,removed

clean:
	$(RM) -r bin
	$(RM) -r dist

fmt: $(SOURCES_NO_VENDOR)
	gofmt -w -s $^

lint:
	golangci-lint run ./...

test:
	$(GO_TEST) ./...

tidy:
	GO11MODULE=on go mod tidy

vet:
	$(GO_VET) -v ./...

.PHONY: all ci install bench clean fmt lint test tidy vet