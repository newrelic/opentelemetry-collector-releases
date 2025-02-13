locals {
  test_specs = {
    for distro in var.distros :
    distro => yamldecode(file("${path.module}/../../../distributions/${distro}/test-spec.yaml"))
  }
  releases_bucket_name                            = "nr-releases"
  required_permissions_boundary_arn_for_new_roles = "arn:aws:iam::${var.aws_account_id}:policy/resource-provisioner-boundary"
}

resource "random_string" "deploy_id" {
  length  = 6
  special = false
}


data "aws_ecr_repository" "ecr_repo" {
    for_each = var.distros
    name = each.value
}

resource "helm_release" "ci_e2e_nightly" {
  for_each   = var.distros

  name  = "ci-e2etest-nightly"
  chart = "../../charts/nr_backend"

  create_namespace = true
  namespace        = "nightly-${each.key}"

  set {
    name  = "image.repository"
    value = data.aws_ecr_repository.ecr_repo[each.key].repository_url
  }

  set {
    name  = "image.tag"
    value = "nightly@${var.nightly_docker_manifest_sha}"
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
    value = "${var.test_environment}-${random_string.deploy_id.result}-${each.key}-k8s_node"
  }
}

module "ci_e2e_ec2" {
  for_each             = toset([for distro in var.distros : distro if local.test_specs[distro].nightly.ec2.enabled])
  source               = "../modules/ec2"
  releases_bucket_name = local.releases_bucket_name
  collector_distro     = each.key
  nr_ingest_key        = var.nr_ingest_key
  # reuse vpc to avoid having to pay for second NAT gateway for this simple use case
  vpc_id              = data.aws_eks_cluster.eks_cluster.vpc_config[0].vpc_id
  deploy_id           = random_string.deploy_id.result
  permission_boundary = local.required_permissions_boundary_arn_for_new_roles
}
