data "aws_vpc" "this" {
  id = "vpc-015d2f927c8b5dea7"
}


module "ci_e2e_cluster" {
  source = "../modules/eks_cluster"

  name       = "aws-ci-e2etest"
  account_id = var.aws_account_id
}
