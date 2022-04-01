DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
###############################################################################
###                                Protobuf                                 ###
###############################################################################

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen sh ./scripts/protocgen.sh
.PHONY: proto-gen

proto-format:
	@echo "Formatting Protobuf files"
	$(DOCKER) run --rm -v $(CURDIR):/workspace \
	--workdir /workspace tendermintdev/docker-build-proto \
	find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -i {} \;
.PHONY: proto-format

# This generates the SDK's custom wrapper for google.protobuf.Any. It should only be run manually when needed
proto-gen-any:
	$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace tendermintdev/sdk-proto-gen sh ./scripts/protocgen-any.sh
.PHONY: proto-gen-any

proto-swagger-gen:
	@./scripts/protoc-swagger-gen.sh
.PHONY: proto-swagger-gen

proto-lint:
	$(DOCKER_BUF) lint --error-format=json
.PHONY: proto-lint

proto-check-breaking:
	# we should turn this back on after our first release
	# $(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=master
.PHONY: proto-check-breaking