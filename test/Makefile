# Define the list of subdirectories that contain a Makefile
SUBDIRS := $(patsubst %/,%,$(wildcard */))
TARGETS := $(SUBDIRS)

.PHONY: all $(TARGETS) clean $(addsuffix -clean,$(TARGETS))

$(TARGETS):
	$(MAKE) -C $@


clean: $(addsuffix -clean,$(SUBDIRS))

$(addsuffix -clean,$(TARGETS)):
	$(MAKE) -C $(patsubst %-clean,%,$@) clean
