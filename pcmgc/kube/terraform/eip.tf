resource "sbercloud_vpc_eip" "eip_1" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "arudometov-eip"
    size        = 5
    share_type  = "PER"
    charge_mode = "traffic"
  }
}
