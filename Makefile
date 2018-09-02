PKG_PREFIX = github.com/ibraimgm/enigma
NAME = enigma
VERSION = 0.1.0
BUILD = `git rev-parse HEAD`

GOTOOLS = \
	github.com/golang/dep

LDFLAGS = -ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

tools: ## Install the tools needed for development and dependency management
	go get -u -v $(GOTOOLS)
	# 'dep' need extra steps:
	cd $(GOPATH)/src/github.com/golang/dep && git describe --abbrev=0 --tags | xargs git checkout && go install ./cmd/dep && git checkout master

clean: ## Remove the generated binary
	-@rm -f $(NAME)

deps: ## Ensure package dependencies are available
	@dep ensure

build: deps $(NAME) ## Builds the application for the current platform
	@true

check: deps ## Run tests (WIP)
	@-rm -f profile.cov
	@-rm -f cover.html
	@go test -covermode=count -coverprofile=profile.cov  `go list ./... | grep -v cmd`

cover: check
	@-rm -f cover.html
	@go tool cover -html=profile.cov -o cover.html

$(NAME):
	@go build $(LDFLAGS) $(PKG_PREFIX)/cmd/enigma

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL = help

.PHONY: tools clean deps build help
