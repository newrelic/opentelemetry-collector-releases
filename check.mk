.PHONY: licenses
licenses: internal/tools generate-sources
	@$(MAKE) for-all-target TARGET="third-party-notices"

.PHONY: licenses-check
licenses-check: internal/tools generate-sources
	@$(MAKE) for-all-target TARGET="third-party-notices-check"
