LIMIT ?= "testing_hosts_linux"
ANSIBLE_FORKS ?= 5

ANSIBLE_INVENTORY_FOLDER ?= $(CURDIR)/test/packaging/ansible

ifeq ($(origin ANSIBLE_INVENTORY_FILE), undefined)
  ANSIBLE_INVENTORY = $(CURDIR)/test/packaging/ansible/inventory.ec2
else
  ANSIBLE_INVENTORY = $(ANSIBLE_INVENTORY_FOLDER)/$(ANSIBLE_INVENTORY_FILE)
endif

.PHONY: test/packaging/requirements
test/packaging/requirements:
	ansible-galaxy install -r test/packaging/ansible/requirements.yaml

.PHONY: test/packaging
test/packaging: test/packaging/requirements
ifndef NR_LICENSE_KEY
	@echo "NR_LICENSE_KEY variable must be provided for test/packaging"
	exit 1
endif

	@ANSIBLE_DISPLAY_SKIPPED_HOSTS=NO ANSIBLE_DISPLAY_OK_HOSTS=NO ansible-playbook -f $(ANSIBLE_FORKS)  -i $(ANSIBLE_INVENTORY) --limit=$(LIMIT) -e collector_nr_license_key=$(NR_LICENSE_KEY) $(CURDIR)/test/packaging/ansible/test.yaml
