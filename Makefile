# Directory where all .proto project directories are located
PROTO_BASE_DIR = proto

# Go version (default to v1.20 if not specified)
GO_VERSION ?= v1.20

# Project name (default to an empty string if not specified)
PROJECT_NAME ?=

# Proto Directory for a specified project
PROTO_DIR = $(PROTO_BASE_DIR)/$(PROJECT_NAME)

# Output directory for generated Go files based on Go version and project name
OUT_DIR = gen/go/$(GO_VERSION)/$(PROJECT_NAME)

# Find all .proto files in the PROTO_DIR and its subdirectories
PROTO_FILES := $(shell find $(PROTO_DIR) -name '*.proto')

# Define the command to run protoc
define PROTOC_GEN
	protoc -I $(PROTO_DIR) $(1) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative
endef

# Ensure output directory exists before generating Protobuf files
proto-generate: $(OUT_DIR)
	$(foreach PROTO_FILE, $(PROTO_FILES), $(call PROTOC_GEN,$(PROTO_FILE));)

# Create the output directory
$(OUT_DIR):
	mkdir -p $(OUT_DIR)

# Clean generated files for a specific Go version and project, excluding go.mod files
clean:
	find gen/go/$(GO_VERSION)/$(PROJECT_NAME) -mindepth 1 -not -name 'go.mod' -exec rm -rf {} +