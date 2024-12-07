.PHONY: ci
ci: build

.PHONY: build
build:
	@$(MAKE) for-all-target TARGET="build-distro"

.PHONY: generate-sources
generate-sources:
	@$(MAKE) for-all-target TARGET="generate-sources-for-distro"

.PHONY: clean-build-dir
clean-build-dir:
	@$(MAKE) for-all-target TARGET="clean-build-dir-distro"
