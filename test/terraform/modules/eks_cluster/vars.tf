variable "name" {
  description = "name of the cluster"
  type        = string
}

variable "vpc_id" {
  description = "VPC id"
  type        = string
}

variable "subnet_ids" {
  description = "subnet ids"
  type = list(string)
}

variable "account_id" {
  description = "account id"
  type        = string
}
