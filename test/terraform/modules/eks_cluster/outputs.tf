output "cluster_name" {
  value = module.eks_cluster.cluster_name
}

output "cluster_endpoint" {
  value = module.eks_cluster.cluster_endpoint
}

output "cluster_certificate_authority_data" {
  value = module.eks_cluster.cluster_certificate_authority_data
}

output "cluster_iam_role_arn" {
  value = module.eks_cluster.cluster_iam_role_arn
}

output "eks_vpc_id" {
  value = module.vpc.vpc_id
}
