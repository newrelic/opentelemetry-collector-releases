.DEFAULT_GOAL := canaries

ROLES_PATH:=$(CURDIR)/roles
COLLECTIONS_PATH:=$(CURDIR)/collections

ANSIBLE_INVENTORY ?= $(CURDIR)/inventory.yaml


$(ROLES_PATH) $(COLLECTIONS_PATH):
	@mkdir -p $@

.PHONY: dependencies
dependencies: $(ROLES_PATH) $(COLLECTIONS_PATH)
	ansible-galaxy role install -r requirements.yml -p $(ROLES_PATH)
	ansible-galaxy collection install -r requirements.yml -p $(COLLECTIONS_PATH)


.PHONY: canaries
canaries: dependencies
ifndef PREVIOUS_VERSION
$(error PREVIOUS_VERSION is not set)
endif

ifndef CURRENT_VERSION
$(error CURRENT_VERSION is not set)
endif
	ansible-playbook -i $(ANSIBLE_INVENTORY) deploy_canaries.yml -e "current_version=$(CURRENT_VERSION) previous_version=$(PREVIOUS_VERSION)"
