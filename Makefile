MODULE   = $(shell env GO111MODULE=on $(GO) list -m)
PKGS     = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
BIN      = $(CURDIR)/bin

GO      = go
M = $(shell printf "\033[34;1m▶\033[0m")

binext=""
ifeq ($(GOOS),windows)
  binext=".exe"
endif

export GO111MODULE=on

$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) installing package $(PACKAGE)...)
	env GO111MODULE=on GOBIN=$(BIN) $(GO) install $(PACKAGE)

GOLINT = $(BIN)/golint
$(BIN)/golint: PACKAGE=golang.org/x/lint/golint@latest

.PHONY: all
all:  clean dep lint vet build test coverage coverhtml

.PHONY: lint
lint: | $(GOLINT) ; $(info $(M) running golint...) @ ## Run golint
		$(GOLINT) -set_exit_status $(PKGS)

.PHONY: vet
vet: ; $(info $(M) running go vet...) @ ## Run go vet on all source files
	$(GO) vet $(PKGS)

.PHONY: test
test: ; ## Run unittests
	$(GO) test -short $(PKGS)

.PHONY: race
race: dep ; ## Run data race detector
	$(GO) test -race -short $(PKGS)

.PHONY: msan
msan: dep ; ## Run memory sanitizer
	$(GO) test -msan -short $(PKGS)

.PHONY: coverage
coverage: ; ## Generate global code coverage report
	./tools/coverage.sh;

.PHONY: coverhtml
coverhtml: ; ## Generate global code coverage report in HTML
	./tools/coverage.sh html;

.PHONY: dep
dep: ; $(info $(M) getting dependencies...) @ ## Get the dependencies
	$(GO) mod download

.PHONY: build
build: $(BIN) dep ; $(info $(M) building executable...) @ ## Build program binary
	CGO_ENABLED=0 $(GO) build \
 		-tags release \
 		-v \
 		-o $(BIN)/$(notdir $(basename $(MODULE)))$(binext) main.go

.PHONY: clean
clean: ; $(info $(M) cleaning...)	@ ## Cleanup everything
	$(GO) clean
	rm -rf $(BIN)

.PHONY: help
help: ## Display this help screen
	@grep -hE '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-17s\033[0m %s\n", $$1, $$2}'
