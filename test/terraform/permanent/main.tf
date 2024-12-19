data "aws_vpc" "this" {
  id = "vpc-015d2f927c8b5dea7"
}

data "aws_subnets" "this" {
  filter {
    name = "vpc-id"
    values = [data.aws_vpc.this.id]
  }

  filter {
    name = "availability-zone"
    values = ["us-east-1a", "us-east-1b"]
  }
}

module "ci_e2e_cluster" {
  source = "../modules/eks_cluster"

  name       = "aws-ci-e2etest"
  account_id = var.aws_account_id

  vpc_id     = data.aws_vpc.this.id
  subnet_ids = data.aws_subnets.this.ids
}
