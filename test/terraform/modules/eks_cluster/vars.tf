variable "name" {
  description = "name of the cluster"
  type        = string
}

variable "permission_boundary" {
  description = "ARN of the IAM policy that is used to set the permissions boundary for the IAM roles created by this module"
  type        = string
}
