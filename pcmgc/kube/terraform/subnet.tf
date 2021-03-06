
variable "subnet_gateway_ip" {
  default = "192.168.0.1"
}

variable "subnet_name" {
  default = "arudometov-subnet"
}

variable "subnet_cidr" {
  default = "192.168.0.0/16"
}

resource "sbercloud_vpc_subnet" "subnet_1" {
  name       = var.subnet_name
  cidr       = var.subnet_cidr
  gateway_ip = var.subnet_gateway_ip
  vpc_id     = sbercloud_vpc.vpc_v1.id
  primary_dns   = "100.125.13.59"
  secondary_dns = "8.8.8.8"
}

