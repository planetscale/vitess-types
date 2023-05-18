gomod := github.com/planetscale/vitess-types

PROTO_OUT := gen
PROTO_SRC := src
VITESS_PROTO_ROOT := $(PROTO_SRC)/vitess

VERSIONS := $(addprefix $(VITESS_PROTO_ROOT)/,$(shell jq -r '.versions | to_entries[].key' manifest.json))

BIN := bin

clean: clean-proto clean-bin

clean-proto:
	rm -rf $(PROTO_OUT)

clean-bin:
	rm -rf $(BIN)

$(BIN):
	mkdir -p $(BIN)

$(PROTO_OUT):
	mkdir -p $(PROTO_OUT)

TOOL_INSTALL := env GOBIN=$(PWD)/$(BIN) go install

$(BIN)/protoc-gen-go: Makefile | $(BIN)
	$(TOOL_INSTALL) google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1

$(BIN)/protoc-gen-go-vtproto: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.4.0

$(BIN)/protoc-gen-connect-go: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@v1.7.0

$(BIN)/gofumpt: Makefile | $(BIN)
	$(TOOL_INSTALL) mvdan.cc/gofumpt@v0.4.0

$(BIN)/buf: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/bufbuild/buf/cmd/buf@v1.19.0

$(BIN)/yq: Makefile | $(BIN)
	$(TOOL_INSTALL) github.com/mikefarah/yq/v4@v4.30.8

PROTO_TOOLS := $(BIN)/protoc-gen-go $(BIN)/protoc-gen-connect-go $(BIN)/protoc-gen-go-vtproto $(BIN)/buf
tools: $(PROTO_TOOLS) $(BIN)/gofumpt $(BIN)/staticcheck $(BIN)/govulncheck $(BIN)/yq

proto: $(PROTO_TOOLS)
	$(BIN)/buf generate -v

download:
	go run download.go
	$(MAKE) clean-proto proto

fmt: fmt-go fmt-proto

fmt-go: $(BIN)/gofumpt
	$(BIN)/gofumpt -l -w .

fmt-yaml: $(BIN)/yq
ifeq (, $(shell command -v fd 2>/dev/null))
	@echo "!! Maybe install 'fd', it's a lot faster (https://github.com/sharkdp/fd)"
	find . -type f \( -name '*.yaml' -o -name '*.yml' \) -exec $(BIN)/yq -iP eval-all . {} \;
else
	fd . -t f -e yaml -e yml -x $(BIN)/yq -iP eval-all . {} \;
endif

update:
	go get -v -u ./...
	go mod tidy
	$(MAKE) clean proto

.PHONY: proto tools update \
		clean clean-proto clean-bin \
		fmt fmt-go fmt-yaml \
		download $(VERSIONS)
