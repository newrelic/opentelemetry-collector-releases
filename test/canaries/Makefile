# Include Ansible dependencies common installation strategy
include ../Ansible.common

.DEFAULT_GOAL := canaries

ANSIBLE_INVENTORY ?= $(CURDIR)/inventory.yml

.PHONY: canaries
canaries: ansible/dependencies
ifndef PREVIOUS_IMAGE
	$(error PREVIOUS_IMAGE is not set)
endif

ifndef CURRENT_IMAGE
	$(error CURRENT_IMAGE is not set)
endif

ifndef NR_LICENSE_KEY_CANARIES
	$(error NR_LICENSE_KEY_CANARIES is not set)
endif
	ansible-playbook -i $(ANSIBLE_INVENTORY) deploy_canaries.yml -e "current_image=$(CURRENT_IMAGE) previous_image=$(PREVIOUS_IMAGE) nr_license_key_canaries=$(NR_LICENSE_KEY_CANARIES) docker_username=$(DOCKER_USERNAME) docker_password=$(DOCKER_PASSWORD)"


.PHONY: clean
clean: ansible/clean
