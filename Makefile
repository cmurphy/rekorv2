#
# Copyright 2025 The Sigstore Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: all test clean clean-gen lint gosec

all: rekor-server

include Makefile.swagger

OPENAPIDEPS = openapi.yaml $(shell find pkg/types -iname "*.json")
SRCS = $(shell find cmd -iname "*.go") $(shell find pkg -iname "*.go"|grep -v pkg/generated) pkg/generated/restapi/configure_rekor_server.go $(SWAGGER_GEN)
TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(abspath $(TOOLS_DIR)/bin)

GIT_VERSION ?= $(shell git describe --tags --always --dirty)
GIT_HASH ?= $(shell git rev-parse HEAD)
DATE_FMT = +%Y-%m-%dT%H:%M:%SZ
SOURCE_DATE_EPOCH ?= $(shell git log -1 --pretty=%ct)
ifdef SOURCE_DATE_EPOCH
    BUILD_DATE ?= $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u -r "$(SOURCE_DATE_EPOCH)" "$(DATE_FMT)" 2>/dev/null || date -u "$(DATE_FMT)")
else
    BUILD_DATE ?= $(shell date "$(DATE_FMT)")
endif
GIT_TREESTATE = "clean"
DIFF = $(shell git diff --quiet >/dev/null 2>&1; if [ $$? -eq 1 ]; then echo "1"; fi)
ifeq ($(DIFF), 1)
    GIT_TREESTATE = "dirty"
endif

REKOR_LDFLAGS=-X sigs.k8s.io/release-utils/version.gitVersion=$(GIT_VERSION) \
              -X sigs.k8s.io/release-utils/version.gitCommit=$(GIT_HASH) \
              -X sigs.k8s.io/release-utils/version.gitTreeState=$(GIT_TREESTATE) \
              -X sigs.k8s.io/release-utils/version.buildDate=$(BUILD_DATE)
SERVER_LDFLAGS=$(REKOR_LDFLAGS)

GOBIN ?= $(shell go env GOPATH)/bin

# Binaries
SWAGGER := $(TOOLS_BIN_DIR)/swagger

Makefile.swagger: $(SWAGGER) $(OPENAPIDEPS) ## Generate Swagger code and Makefile
	$(SWAGGER) validate openapi.yaml
	$(SWAGGER) generate client -f openapi.yaml -q -r COPYRIGHT.txt -t pkg/generated --additional-initialism=TUF --additional-initialism=DSSE
	$(SWAGGER) generate server -f openapi.yaml -q -r COPYRIGHT.txt -t pkg/generated --exclude-main -A rekor_server --flag-strategy=pflag --default-produces application/json --additional-initialism=TUF --additional-initialism=DSSE
	@echo "# This file is generated after swagger runs as part of the build; do not edit!" > Makefile.swagger
	@echo "SWAGGER_GEN=`find pkg/generated/client pkg/generated/models pkg/generated/restapi -iname '*.go' | grep -v 'configure_rekor_server' | sort -d | tr '\n' ' ' | sed 's/ $$//'`" >> Makefile.swagger;

lint: ## Run golangci-lint checks
	$(GOBIN)/golangci-lint run -v ./...

gosec: ## Run gosec security scanner
	$(GOBIN)/gosec ./...

rekor-server: $(SRCS)
	CGO_ENABLED=0 go build -trimpath -ldflags "$(SERVER_LDFLAGS)" -o rekor-server ./cmd/rekor-server

test: ## Run all tests
	go test ./...

clean: ## Remove built binaries and artifacts
	rm -rf dist
	rm -rf hack/tools/bin
	rm -rf rekor-server

clean-gen: clean ## Clean generated files and swagger code
	rm -rf $(SWAGGER_GEN)

## --------------------------------------
## Tooling Binaries
## --------------------------------------

$(SWAGGER): $(TOOLS_DIR)/go.mod ## Build swagger tool
	cd $(TOOLS_DIR); go build -trimpath -tags=tools -o $(TOOLS_BIN_DIR)/swagger github.com/go-swagger/go-swagger/cmd/swagger

##################
# help
##################

help: ## Display this help message
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
