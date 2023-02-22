ci: build

.PHONY: terraform/backend
terraform/backend:
ifndef TERRAFORM_BACKEND
	$(error TERRAFORM_BACKEND is undefined)
endif
	sed "s/REPLACE_ME/${TERRAFORM_BACKEND}/g" "$(CURDIR)/ci/terraform/terraform.backend.tf.dist" > "$(CURDIR)/ci/terraform/terraform.backend.tf"

.PHONY: terraform/clean
terraform/clean:
	rm "$(CURDIR)/ci/terraform/terraform.backend.tf"

.PHONY: terraform/provision-ec2
terraform/provision-ec2: terraform/backend
	cd "$(CURDIR)/ci/terraform" && \
	terraform init -reconfigure && \
	terraform apply -auto-approve

.PHONY: terraform/destroy-ec2
terraform/destroy-ec2: terraform/backend
	cd "$(CURDIR)/ci/terraform" && \
    terraform init -reconfigure && \
    terraform destroy -auto-approve