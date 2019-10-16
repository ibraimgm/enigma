PKG_PREFIX = github.com/ibraimgm/enigma
NAME = enigma
VERSION = 0.1.0
BUILD = `git rev-parse HEAD`

LDFLAGS = -ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

clean: ## Remove the generated binary
	-@rm -f $(NAME)

download: ## download dependencies
	@ go get ./...

build: ## Builds the application for the current platform
	@ go build $(LDFLAGS) $(PKG_PREFIX)/cmd/enigma

check: ## Run tests (WIP)
	@ go test -covermode=count -coverprofile=coverage.txt  `go list ./... | grep -v cmd`

cover: check
	@-rm -f cover.html
	@go tool cover -html=coverage.txt -o cover.html

$(NAME): build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL = help

.PHONY: tools clean deps build help
