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
ifndef PREVIOUS_IMAGE
$(error PREVIOUS_IMAGE is not set)
endif

ifndef CURRENT_IMAGE
$(error CURRENT_IMAGE is not set)
endif
	ansible-playbook -i $(ANSIBLE_INVENTORY) deploy_canaries.yml -e "current_image=$(CURRENT_IMAGE) previous_image=$(PREVIOUS_IMAGE)"
