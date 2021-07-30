variable "vpc_name" {
  default = "arudometov-vpc"
}

variable "vpc_cidr" {
  default = "192.168.0.0/16"
}

resource "sbercloud_vpc" "vpc_v1" {
  name = var.vpc_name
  cidr = var.vpc_cidr
}

