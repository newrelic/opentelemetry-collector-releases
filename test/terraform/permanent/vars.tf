variable "aws_account_id" {
  type        = string
  description = "AWS account id to deploy to"
}

variable "aws_region" {
  type        = string
  description = "AWS region to deploy to"
  default     = "us-east-1"
}

variable "nr_backend_url" {
  type        = string
  description = "NR endpoint used in test cluster"
  sensitive   = true
}

variable "nr_ingest_key" {
  type        = string
  description = "NR ingest key used in test cluster"
  sensitive   = true
}

variable "test_environment" {
  type        = string
  description = "Name of test environment to distinguish entities"
  default     = "nightly"
}

