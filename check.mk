FIND_MOD_ARGS=-type f -name "go.mod"
TO_MOD_DIR=dirname {} \; | sort | grep -E '^./'
DISTRIBUTIONS_MODS := $(shell find ./distributions/**/* $(FIND_MOD_ARGS) -exec $(TO_MOD_DIR) )

## Retrive all used dependencies in the same output
RESULTS := $(foreach dir,$(DISTRIBUTIONS_MODS),$(shell cd $(dir) && go list -mod=mod -m -json all))


.PHONY: third-party-notices
third-party-notices: OUT ?= THIRD_PARTY_NOTICES.md
third-party-notices: internal/tools
	## subst replace every occurrence of " with \"
	@echo "$(subst ",\",$(RESULTS))" | go-licence-detector \
		-rules $(CURDIR)/internal/assets/license/rules.json \
		-noticeTemplate $(CURDIR)/internal/assets/license/THIRD_PARTY_NOTICES.md.tmpl \
		-noticeOut $(OUT)


.PHONY: third-party-notices-check
third-party-notices-check: internal/tools
	@OUT=THIRD_PARTY_NOTICES_GENERATED.md $(MAKE) third-party-notices
	@diff THIRD_PARTY_NOTICES_GENERATED.md THIRD_PARTY_NOTICES.md 2>&1 > /dev/null || (echo "THIRD_PARTY_NOTICES.md should be generated"; exit 1)
