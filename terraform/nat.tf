resource "sbercloud_nat_gateway" "nat_1" {
  name                = "arudometov-nat"
  spec                = "1"
  vpc_id = sbercloud_vpc.vpc_v1.id
  subnet_id = sbercloud_vpc_subnet.subnet_1.id
}

resource "sbercloud_nat_snat_rule" "snat_subnet_01" {
  nat_gateway_id = sbercloud_nat_gateway.nat_1.id
  subnet_id      = sbercloud_vpc_subnet.subnet_1.id
  floating_ip_id = sbercloud_vpc_eip.eip_1.id
}

resource "sbercloud_nat_dnat_rule" "dnat_223" {
  floating_ip_id        = sbercloud_vpc_eip.eip_1.id
  nat_gateway_id        = sbercloud_nat_gateway.nat_1.id
  private_ip            = sbercloud_compute_instance.ecs_1.access_ip_v4
  protocol              = "tcp"
  internal_service_port = 22
  external_service_port = 22
}

resource "sbercloud_nat_dnat_rule" "dnat_222" {
  floating_ip_id        = sbercloud_vpc_eip.eip_1.id
  nat_gateway_id        = sbercloud_nat_gateway.nat_1.id
  private_ip            = sbercloud_compute_instance.ecs_2.access_ip_v4
  protocol              = "tcp"
  internal_service_port = 22
  external_service_port = 23
}

resource "sbercloud_nat_dnat_rule" "dnat_80" {
  floating_ip_id        = sbercloud_vpc_eip.eip_1.id
  nat_gateway_id        = sbercloud_nat_gateway.nat_1.id
  private_ip            = sbercloud_compute_instance.ecs_1.access_ip_v4
  protocol              = "tcp"
  internal_service_port = 8080
  external_service_port = 80
}

