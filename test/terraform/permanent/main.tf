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

resource "helm_release" "ci_e2e_nightly" {
  depends_on = [module.ci_e2e_cluster]

  name      = "ci-e2etest-nightly"
  chart     = "../../charts/nr_backend"

  create_namespace = true
  namespace = "nightly"

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
}
