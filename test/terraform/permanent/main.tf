# read default vpc and create security group to validate basic setup and role permissions
data "aws_vpc" "example_read_resource" {
  id = "vpc-015d2f927c8b5dea7"
}

resource "aws_security_group" "example_write_resource" {
  name   = "example_write_resource_from_terraform"
  vpc_id = data.aws_vpc.example_read_resource.id

  tags = {
    created_by = "terraform"
  }
}