ci: build

.PHONY: terraform/backend
terraform/backend:
ifndef TERRAFORM_BACKEND
	$(error TERRAFORM_BACKEND is undefined)
endif
	sed "s/REPLACE_ME/${TERRAFORM_BACKEND}/g" ./terraform/terraform.backend.tf.dist > ./terraform/terraform.backend.tf

.PHONY: terraform/clean
terraform/clean:
	rm ./terraform/terraform.backend.tf

.PHONY: terraform/provision-ec2
terraform/provision-ec2: terraform/backend
	cd ./terraform && \
	terraform init -reconfigure && \
	terraform apply -auto-approve

.PHONY: terraform/destroy-ec2
terraform/destroy-ec2: terraform/backend
	cd ./terraform && \
    terraform init -reconfigure && \
    terraform destroy -auto-approve