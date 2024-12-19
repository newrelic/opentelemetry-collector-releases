locals {
  iam_role_permissions_boundary = "arn:aws:iam::${var.account_id}:policy/resource-provisioner-boundary"
}

module "eks_cluster" {
  source  = "terraform-aws-modules/eks/aws"
  version = "20.31.4"

  cluster_name    = var.name
  cluster_version = "1.31"

  cluster_endpoint_private_access = true
  cluster_endpoint_public_access  = true

  vpc_id     = var.vpc_id
  subnet_ids = var.subnet_ids

  enable_cluster_creator_admin_permissions = true
  iam_role_permissions_boundary            = local.iam_role_permissions_boundary

  eks_managed_node_group_defaults = {
    iam_role_permissions_boundary = local.iam_role_permissions_boundary

    instance_types = ["m5.large"]

    min_size     = 2
    max_size     = 2
    desired_size = 2
  }

  eks_managed_node_groups = {
    default = {
      # Starting on 1.30, AL2023 is the default AMI type for EKS managed node groups
      ami_type = "AL2023_x86_64_STANDARD"
    }
  }
}


