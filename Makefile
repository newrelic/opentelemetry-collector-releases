# Exclude current and hidden directories
FIND_PATH = . -not -path '*/\.*' -mindepth 2
# Define the list of subdirectories that contain a Makefile
SUBDIRS := $(patsubst ./%/Makefile,%,$(shell find $(FIND_PATH) -name Makefile))
TARGETS := $(SUBDIRS)

.PHONY: all $(TARGETS) clean $(addsuffix -clean,$(TARGETS)) help

$(TARGETS):
	$(MAKE) -C $@

clean: $(addsuffix -clean,$(SUBDIRS))

$(addsuffix -clean,$(TARGETS)):
	$(MAKE) -C $(patsubst %-clean,%,$@) clean

help:
	@echo "## Available targets:"
	@echo $(TARGETS)
	@echo "## Available clean targets:"
	@echo $(addsuffix -clean,$(TARGETS))
