.PHONY: terraform/backend
terraform/backend:
ifndef TAG_OR_UNIQUE_NAME
	$(error TAG_OR_UNIQUE_NAME is undefined)
endif
	sed "s/TAG_OR_UNIQUE_NAME/${TAG_OR_UNIQUE_NAME}/g" "$(CURDIR)/test/provision/terraform/terraform.backend.tf.dist" > "$(CURDIR)/test/provision/terraform/terraform.backend.tf"
	sed "s/TAG_OR_UNIQUE_NAME/${TAG_OR_UNIQUE_NAME}/g" "$(CURDIR)/test/provision/terraform/caos.auto.tfvars.dist" > "$(CURDIR)/test/provision/terraform/caos.auto.tfvars"


.PHONY: terraform/clean
terraform/clean:
	rm "$(CURDIR)/test/provision/terraform/terraform.backend.tf"

.PHONY: test/provision
test/provision: terraform/backend
	cd "$(CURDIR)/test/provision/terraform" && \
	terraform init -reconfigure && \
	terraform apply -auto-approve

.PHONY: test/provision-destroy
test/provision-destroy: terraform/backend
	cd "$(CURDIR)/test/provision/terraform" && \
    terraform init -reconfigure && \
    terraform destroy -auto-approve
