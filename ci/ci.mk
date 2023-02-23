ci: build

.PHONY: terraform/backend
terraform/backend:
ifndef TAG_OR_UNIQUE_NAME
	$(error TAG_OR_UNIQUE_NAME is undefined)
endif
	sed "s/TAG_OR_UNIQUE_NAME/${TAG_OR_UNIQUE_NAME}/g" "$(CURDIR)/ci/terraform/terraform.backend.tf.dist" > "$(CURDIR)/ci/terraform/terraform.backend.tf"
	sed "s/TAG_OR_UNIQUE_NAME/${TAG_OR_UNIQUE_NAME}/g" "$(CURDIR)/ci/terraform/caos.auto.tfvars.dist" > "$(CURDIR)/ci/terraform/caos.auto.tfvars"


.PHONY: terraform/clean
terraform/clean:
	rm "$(CURDIR)/ci/terraform/terraform.backend.tf"

.PHONY: terraform/provision-ec2
terraform/provision-ec2: terraform/backend
	touch "/srv/runner/inventory/test-file.ec2"
	cd "$(CURDIR)/ci/terraform" && \
	terraform init -reconfigure && \
	terraform apply -auto-approve

.PHONY: terraform/destroy-ec2
terraform/destroy-ec2: terraform/backend
	cd "$(CURDIR)/ci/terraform" && \
    terraform init -reconfigure && \
    terraform destroy -auto-approve