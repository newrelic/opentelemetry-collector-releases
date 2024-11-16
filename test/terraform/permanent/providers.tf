terraform {
  required_version = "1.9.8"
  required_providers {
    aws = {
      version = "5.76.0"
    }
  }
}

terraform {
  backend "s3" {
    encrypt        = true
    dynamodb_table = "terraform-states-lock"
    region         = "us-east-1"
    key            = "newrelic/opentelemetry-collector-releases/permanent/terraform.tfstate"
    # 'bucket' and 'role_arn' provided via '-backend-config'
  }
}

provider "aws" {
  region = var.aws_region
  allowed_account_ids = [var.aws_account_id]
  # expect AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY as env vars

  assume_role {
    role_arn = "arn:aws:iam::${var.aws_account_id}:role/resource-provisioner"
  }
}
