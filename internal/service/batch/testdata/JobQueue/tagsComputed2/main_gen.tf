# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

provider "null" {}

resource "aws_batch_job_queue" "test" {
  name     = var.rName
  priority = 1
  state    = "DISABLED"

  compute_environments = [aws_batch_compute_environment.test.arn]

  tags = {
    (var.unknownTagKey) = null_resource.test.id
    (var.knownTagKey)   = var.knownTagValue
  }
}

resource "aws_batch_compute_environment" "test" {
  compute_environment_name = var.rName
  service_role             = aws_iam_role.batch_service.arn
  type                     = "UNMANAGED"

  depends_on = [aws_iam_role_policy_attachment.batch_service]
}

data "aws_partition" "current" {}

resource "aws_iam_role" "batch_service" {
  name = "${var.rName}-batch-service"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": "batch.${data.aws_partition.current.dns_suffix}"
      }
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "batch_service" {
  role       = aws_iam_role.batch_service.name
  policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/service-role/AWSBatchServiceRole"
}

resource "aws_iam_role" "ecs_instance" {
  name = "${var.rName}-ecs-instance"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
        "Service": "ec2.${data.aws_partition.current.dns_suffix}"
        }
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ecs_instance" {
  role       = aws_iam_role.ecs_instance.name
  policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

resource "aws_iam_instance_profile" "ecs_instance" {
  name = aws_iam_role.ecs_instance.name
  role = aws_iam_role_policy_attachment.ecs_instance.role
}

resource "null_resource" "test" {}

variable "rName" {
  type     = string
  nullable = false
}

variable "unknownTagKey" {
  type     = string
  nullable = false
}

variable "knownTagKey" {
  type     = string
  nullable = false
}

variable "knownTagValue" {
  type     = string
  nullable = false
}

