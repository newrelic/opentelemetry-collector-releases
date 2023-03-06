.PHONY: terraform/backend
terraform/backend:
ifndef TAG_OR_UNIQUE_NAME
	$(error TAG_OR_UNIQUE_NAME is undefined)
endif
	sed "s/TAG_OR_UNIQUE_NAME/${TAG_OR_UNIQUE_NAME}/g" "$(CURDIR)/test/provision/terraform/terraform.backend.tf.dist" > "$(CURDIR)/test/provision/terraform/terraform.backend.tf"

.PHONY: terraform/prepare_ec2_list-linux
terraform/prepare_ec2_list-linux:
ifndef TAG_OR_UNIQUE_NAME
	$(error TAG_OR_UNIQUE_NAME is undefined)
endif
	sed "s/TAG_OR_UNIQUE_NAME/${TAG_OR_UNIQUE_NAME}/g" "$(CURDIR)/test/provision/terraform/caos-linux.auto.tfvars.dist" > "$(CURDIR)/test/provision/terraform/caos.auto.tfvars"


.PHONY: ansible/requirements
ansible/requirements:
	ansible-galaxy install -r "$(CURDIR)/test/packaging/ansible/requirements.yaml"

.PHONY: terraform/clean
terraform/clean:
	rm "$(CURDIR)/test/provision/terraform/terraform.backend.tf" "$(CURDIR)/test/provision/terraform/caos.auto.tfvars"

.PHONY: test/provision-linux
test/provision-linux: terraform/backend terraform/prepare_ec2_list-linux ansible/requirements
	cd "$(CURDIR)/test/provision/terraform" && \
	terraform init -reconfigure && \
	terraform apply -auto-approve

.PHONY: test/provision-destroy
test/provision-destroy: terraform/backend
	cd "$(CURDIR)/test/provision/terraform" && \
    terraform init -reconfigure && \
    terraform destroy -auto-approve
