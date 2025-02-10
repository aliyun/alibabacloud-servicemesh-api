BIN_DIR = ./bin
CLIENT_GEN = $(BIN_DIR)/client-gen
LISTERS_GEN = $(BIN_DIR)/lister-gen
INFORMERS_GEN = $(BIN_DIR)/informer-gen
OUTPUT_BASE ?= asm-generated-clients

CLIENT_GEN_VERSION = v0.27.0
LISTERS_GEN_VERSION = v0.27.0
INFORMERS_GEN_VERSION = v0.27.0

.PHONY: gen-asm-clients client-gen lister-gen informer-gen client-gen-tools

client-gen-tools: client-gen lister-gen informer-gen 

gen-asm-clients: client-gen-tools
	OUTPUT_BASE=$(OUTPUT_BASE) bash scripts/gen-asm-clients.sh

client-gen:
	$(call go-get-tool,$(BIN_DIR)/client-gen,k8s.io/code-generator/cmd/client-gen@$(CLIENT_GEN_VERSION))

lister-gen:
	$(call go-get-tool,$(BIN_DIR)/lister-gen,k8s.io/code-generator/cmd/lister-gen@$(LISTERS_GEN_VERSION))

informer-gen:
	$(call go-get-tool,$(BIN_DIR)/informer-gen,k8s.io/code-generator/cmd/informer-gen@$(INFORMERS_GEN_VERSION))

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2) to $(PROJECT_DIR)/bin" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef