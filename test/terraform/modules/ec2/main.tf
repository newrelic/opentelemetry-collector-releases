locals {
  collector_reported_hostname = "${var.test_environment}-${var.deploy_id}-ec2_ubuntu24_04-0"
  collector_version_specifier = "${var.collector_version == "" ? "" : "="}${var.collector_version}"
}

data "aws_ami" "ubuntu24_04" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

data "aws_vpc" "ec2_vpc" {
  id = var.vpc_id
}

data "aws_subnets" "private_subnets" {
  filter {
    name   = "vpc-id"
    values = [var.vpc_id]
  }
  filter {
    name   = "tag:Name"
    values = ["*private*"]
  }
}

resource "aws_security_group" "ec2_allow_all_egress" {
  name        = "nightly-ec2-all-egress"
  description = "Allow all outbound traffic"
  vpc_id      = data.aws_vpc.ec2_vpc.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "ubuntu24_04" {
  ami                    = data.aws_ami.ubuntu24_04.id
  instance_type          = "t2.micro"
  subnet_id              = data.aws_subnets.private_subnets.ids[0]
  vpc_security_group_ids = [aws_security_group.ec2_allow_all_egress.id]

  user_data = <<-EOF
              #!/bin/bash
              echo 'Configuring swap file to ensure system.paging.usage metric gets published'
              dd if=/dev/zero of=/swapfile bs=1M count=128
              chmod 600 /swapfile
              mkswap /swapfile
              swapon /swapfile
              swapon -s
              echo 'Installing Collector'
              export DEBIAN_FRONTEND=noninteractive
              curl -fsSL https://download.newrelic.com/infrastructure_agent/gpg/newrelic-infra.gpg | gpg --dearmor -o /usr/share/keyrings/newrelic-infra.gpg
              printf '%s\n' 'deb [signed-by=/usr/share/keyrings/newrelic-infra.gpg] http://nr-downloads-ohai-staging.s3-website-us-east-1.amazonaws.com/infrastructure_agent/linux/apt noble main' | tee /etc/apt/sources.list.d/newrelic-infra.list
              apt-get update
              apt-get install nr-otel-collector${local.collector_version_specifier}
              echo 'NEW_RELIC_LICENSE_KEY=${var.nr_ingest_key}' >> /etc/nr-otel-collector/nr-otel-collector.conf
              echo "OTEL_RESOURCE_ATTRIBUTES='host.name=${local.collector_reported_hostname}'" >> /etc/nr-otel-collector/nr-otel-collector.conf
              systemctl reload-or-restart nr-otel-collector.service
              EOF
}
