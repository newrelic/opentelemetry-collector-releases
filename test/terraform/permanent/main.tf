locals {
  distros = toset(distinct(flatten([
    for _, v in fileset(path.module, "../../../distributions/**") :
    regex("../../../distributions/([^/]*).*", dirname(v))
  ])))
  releases_bucket_name                            = "nr-releases"
  required_permissions_boundary_arn_for_new_roles = "arn:aws:iam::${var.aws_account_id}:policy/resource-provisioner-boundary"
}

data "aws_vpc" "this" {
  id = "vpc-015d2f927c8b5dea7"
}


module "ci_e2e_cluster" {
  source = "../modules/eks_cluster"

  name       = "aws-ci-e2etest"
  permission_boundary = local.required_permissions_boundary_arn_for_new_roles
}

data "aws_caller_identity" "current" {}

data "aws_iam_session_context" "current" {
  arn = data.aws_caller_identity.current.arn
}

module "ecr" {
  depends_on = [module.ci_e2e_cluster]

  for_each = local.distros

  source = "terraform-aws-modules/ecr/aws"

  repository_name = each.value

  repository_image_tag_mutability = "MUTABLE"

  repository_read_write_access_arns = [data.aws_iam_session_context.current.issuer_arn]
  repository_read_access_arns       = [module.ci_e2e_cluster.cluster_iam_role_arn]

  repository_lifecycle_policy = jsonencode({
    rules = [
      {
        rulePriority = 1,
        description  = "Keep last 10 nightly images",
        selection = {
          tagStatus     = "tagged",
          tagPrefixList = ["nightly"],
          countType     = "imageCountMoreThan",
          countNumber   = 10
        },
        action = {
          type = "expire"
        }
      }
    ]
  })
}

module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"

  bucket = local.releases_bucket_name
  acl    = "private"

  control_object_ownership = true
  object_ownership         = "ObjectWriter"

  versioning = {
    enabled = true
  }
}

data "aws_eks_cluster_auth" "this" {
  name = module.ci_e2e_cluster.cluster_name
}

resource "random_string" "deploy_id" {
  length  = 6
  special = false
}

resource "helm_release" "ci_e2e_nightly" {
  depends_on = [module.ci_e2e_cluster, module.ecr]

  name  = "ci-e2etest-nightly"
  chart = "../../charts/nr_backend"

  create_namespace = true
  namespace        = "nightly"

  set {
    name  = "image.repository"
    value = module.ecr["nr-otel-collector"].repository_url
  }

  set {
    name  = "image.tag"
    value = "nightly"
  }

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
    value = "${var.test_environment}-${random_string.deploy_id.result}-k8s_node"
  }
}

module "ci_e2e_ec2" {
  source               = "../modules/ec2"
  releases_bucket_name = local.releases_bucket_name
  nr_ingest_key        = var.nr_ingest_key
  # reuse vpc to avoid having to pay for second NAT gateway for this simple use case
  vpc_id              = module.ci_e2e_cluster.eks_vpc_id
  deploy_id           = random_string.deploy_id.result
  # TODO: using sticky version until we have nightly-labeled binaries setup
  collector_version   = "0.8.7"
  permission_boundary = local.required_permissions_boundary_arn_for_new_roles
}
