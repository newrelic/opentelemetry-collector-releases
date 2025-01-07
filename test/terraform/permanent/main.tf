data "aws_vpc" "this" {
  id = "vpc-015d2f927c8b5dea7"
}


module "ci_e2e_cluster" {
  source = "../modules/eks_cluster"

  name       = "aws-ci-e2etest"
  account_id = var.aws_account_id
}

data "aws_eks_cluster_auth" "this" {
  name = module.ci_e2e_cluster.cluster_name
}

provider "random" {}

resource "random_string" "hostname_suffix" {
  length  = 6
  special = false
}

resource "helm_release" "ci_e2e_nightly" {
  depends_on = [module.ci_e2e_cluster]

  name  = "ci-e2etest-nightly"
  chart = "../../charts/nr_backend"

  create_namespace = true
  namespace        = "nightly"

  set {
    name  = "image.pullPolicy"
    value = "Always"
  }

  set {
    name  = "secrets.nrBackendUrl"
    value = var.nr_backend_url
  }

  set {
    name  = "secrets.nrIngestKey"
    value = var.nr_ingest_key
  }

  set {
    name  = "collector.hostname"
    value = "nr-otel-collector-${var.test_environment}-${random_string.hostname_suffix.result}"
  }
}

data "aws_caller_identity" "current" {}

module "ecr" {
  depends_on = [module.ci_e2e_cluster]

  source = "terraform-aws-modules/ecr/aws"

  repository_name = "nr-otel-collector"

  repository_read_write_access_arns = [data.aws_caller_identity.current.arn]
  repository_read_access_arns = [module.ci_e2e_cluster.cluster_iam_role_arn]

  repository_lifecycle_policy = jsonencode({
    rules = [
      {
        rulePriority = 1,
        description  = "Keep last 10 nightly images",
        selection = {
          tagStatus   = "tagged",
          tagPrefixList = ["nightly"],
          countType   = "imageCountMoreThan",
          countNumber = 10
        },
        action = {
          type = "expire"
        }
      }
    ]
  })
}
