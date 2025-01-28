output "ecr_repository_urls" {
  value = [for key, value in module.ecr : value.repository_url]
}

output "distros" {
  value = local.distros
}
