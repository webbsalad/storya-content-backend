PROTO_SRC_DIR := ./api/content
PROTO_OUT_DIR := ./internal/pb
VENDOR_DIR := ./vendor.protogen

PROTOC_FLAGS := \
	-I . \
	-I $(VENDOR_DIR) \
	--validate_out="lang=go:$(PROTO_OUT_DIR)" \
	--go_out=$(PROTO_OUT_DIR) \
	--go-grpc_out=$(PROTO_OUT_DIR)

GRPC_GATEWAY_FLAGS := \
	--grpc-gateway_out $(PROTO_OUT_DIR) \
	--proto_path=$(VENDOR_DIR) \
	--grpc-gateway_opt generate_unbound_methods=true

.PHONY: all proto proto-deps clean clean-vendor

all: proto

proto: proto-deps
	protoc $(PROTOC_FLAGS) $(PROTO_SRC_DIR)/*.proto
	protoc -I . $(GRPC_GATEWAY_FLAGS) $(PROTO_SRC_DIR)/*.proto

proto-deps:
	rm -rf $(VENDOR_DIR)
	mkdir -p $(VENDOR_DIR)

	git clone --depth=1 https://github.com/googleapis/googleapis.git $(VENDOR_DIR)/googleapis
	mv $(VENDOR_DIR)/googleapis/google/ $(VENDOR_DIR)
	rm -rf $(VENDOR_DIR)/googleapis

	git clone --depth=1 https://github.com/bufbuild/protoc-gen-validate.git $(VENDOR_DIR)/protoc-gen-validate
	mv $(VENDOR_DIR)/protoc-gen-validate/validate/ $(VENDOR_DIR)
	rm -rf $(VENDOR_DIR)/protoc-gen-validate

quality:
	go test -cover -coverprofile=coverage.out ./... 
	golangci-lint run 