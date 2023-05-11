
GROUP ?= all
FOR_GROUP_TARGET=for-$(GROUP)-target

# Define a delegation target for each distribution
.PHONY: $(ALL_DISTRIBUTIONS)
$(ALL_DISTRIBUTIONS):
	$(MAKE) -C $@ $(TARGET)

# Trigger each module's delegation target
.PHONY: for-all-target
for-all-target: $(ALL_DISTRIBUTIONS)

.PHONY: licenses
licenses: internal/tools build
	@$(MAKE) $(FOR_GROUP_TARGET) TARGET="third-party-notices"

.PHONY: licenses-check
licenses-check: internal/tools build
	@$(MAKE) $(FOR_GROUP_TARGET) TARGET="third-party-notices-check"
