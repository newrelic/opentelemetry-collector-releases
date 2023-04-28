##########################################
# 		     Dynamic targets 			 #
##########################################
# Exclude current, distributions and hidden directories
FIND_PATH = . -mindepth 2 -not -path '*/\.*' -not -path '*/distributions/*'
# Define the list of subdirectories that contain a Makefile
SUBDIRS := $(patsubst ./%/Makefile,%,$(shell find $(FIND_PATH) -name Makefile))
TARGETS := $(SUBDIRS)

.PHONY: all $(TARGETS) clean $(addsuffix -clean,$(TARGETS)) help

$(TARGETS):
	$(MAKE) -C $@

clean: $(addsuffix -clean,$(SUBDIRS))

$(addsuffix -clean,$(TARGETS)):
	$(MAKE) -C $(patsubst %-clean,%,$@) clean


##########################################
# 		     Static targets 			 #
##########################################
FIND_MOD_ARGS=-type f -name "manifest.yaml"
TO_DIR=dirname {} \; | sort | grep -E '^./'

ALL_DISTRIBUTIONS := $(shell find ./distributions/* $(FIND_MOD_ARGS) -exec $(TO_DIR) )

include build.mk
include check.mk

help:
	@echo "## Available targets:"
	@echo $(TARGETS)
	@echo "## Available clean targets:"
	@echo $(addsuffix -clean,$(TARGETS))
