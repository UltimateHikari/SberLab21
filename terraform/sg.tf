locals {
  rules = {
    http-rule = {
      description = "Allow HTTP from anywhere",
      protocol = "tcp",
      port = 80,
      source = "0.0.0.0/0"
    },
    ssh-rule = {
      description = "Allow SSH from only one source",
      protocol = "tcp",
      port = 22,
      source = "0.0.0.0/0"
    },
    udp-rule = {
      description = "allow dns for docker",
      protocol = "udp",
      port = 53,
      source = "0.0.0.0/0"
  
    }
  }
}


resource "sbercloud_networking_secgroup" "secgroup_1" {
  name        = "arudometov-sg"
  description = "My neutron security group"
}

resource "sbercloud_networking_secgroup_rule" "sg_rule_01" {
  for_each = local.rules

  direction         = "ingress"
  ethertype         = "IPv4"
  description       = each.value.description
  protocol          = each.value.protocol
  port_range_min    = each.value.port
  port_range_max    = each.value.port
  remote_ip_prefix  = each.value.source

  security_group_id = sbercloud_networking_secgroup.secgroup_1.id
}
