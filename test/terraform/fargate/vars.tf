#######################################
# Global vars
#######################################

variable "region" {
  default = "us-east-2"
}

variable "accountId" {
  default = "018789649883"
}

variable "vpc_id" {
  default = "vpc-0a3c00f5dc8645fe0"
}

variable "vpc_subnet" {
  default = "subnet-09b64de757828cdd4"
}

variable "cluster_name" {
  default = "caos_otel_releases"
}

#######################################
# Task definition
#######################################

variable "task_command" {
  default = [
    "test/automated-run"
  ]
}

variable "secret_name_ssh" {
  default = "caos/canaries/ssh_key-UBSKNA"
}

variable "secret_name_license" {
  default = "caos/canaries/license-f9eYwe"
}

variable "secret_name_license_canaries" {
  default = "caos/canaries/license_canaries-1DCE1L"
}

variable "secret_name_account" {
  default = "caos/canaries/account-kKFMGP"
}

variable "secret_name_api" {
  default = "caos/canaries/api-9q0NPb"
}

variable "secret_name_nr_api_key" {
  default = "caos/canaries/nr_api_key-xadBYJ"
}

variable "task_container_image" {
  default = "ghcr.io/newrelic/fargate-runner-action:latest"
}

variable "task_logs_group" {
  default = "/ecs/test-prerelease-otel-releases"
}

variable "task_container_name" {
  default = "test-otel-releases"
}

variable "task_name_prefix" {
  default = "otel-releases"
}

variable "task_logs_prefix" {
  default = "ecs-otel-releases"
}

#######################################
# EFS volume
#######################################

variable "efs_volume_mount_point" {
  default = "/srv/runner/inventory"
}

variable "efs_volume_name" {
  default = "shared-otel-releases"
}

variable "canaries_security_group" {
  default = "sg-044ef7bc34691164a"
}

#######################################
# OIDC permissions
#######################################

variable "oidc_repository" {
  default = "repo:newrelic/opentelemetry-collector-releases:*"
}

variable "oidc_role_name" {
  default = "caos-pipeline-oidc-otel-releases"
}
