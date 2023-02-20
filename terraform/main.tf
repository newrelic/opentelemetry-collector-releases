variable "ec2_otels" {
  default = ""
}
variable "nr_license_key" {
  default = ""
}
variable "otlp_endpoint" {
  default = ""
}
variable "pvt_key" {
  default = ""
}
variable "ssh_pub_key" {
  default = ""
}
variable "ansible_playbook" {
  default = ""
}

variable "region" {
  default = "us-east-2"
}

provider "aws" {
  region = var.region
}

module "otel-env-provisioner" {
  source         = "git::https://github.com/newrelic-experimental/otel-env-provisioner//terraform/otel-ec2"
  ec2_otels      = var.ec2_otels
  nr_license_key = var.nr_license_key
  otlp_endpoint  = var.otlp_endpoint
  pvt_key        = var.pvt_key
  ssh_pub_key    = var.ssh_pub_key
}