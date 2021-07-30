data "sbercloud_images_image" "myimage" {
  name        = "Ubuntu 20.04 server 64bit"
  most_recent = true
}

resource "sbercloud_compute_instance" "ecs_1" {
  name              = "node-1"
  image_id          = data.sbercloud_images_image.myimage.id
  flavor_id         = "s6.xlarge.2"
  key_pair = "KeyPair-rudometov"
  security_groups   = ["default", "arudometov-sg"]
  availability_zone = "ru-moscow-1a"

  system_disk_type = "SAS"
  system_disk_size = 40

  network {
    uuid = sbercloud_vpc_subnet.subnet_1.id
  }
}

resource "sbercloud_compute_instance" "ecs_2" {
  name              = "node-2"
  image_id          = data.sbercloud_images_image.myimage.id
  flavor_id         = "s6.xlarge.2"
  key_pair = "KeyPair-rudometov"
  security_groups   = ["default", "arudometov-sg"]
  availability_zone = "ru-moscow-1a"
  system_disk_type = "SAS"
  system_disk_size = 40

  network {
    uuid = sbercloud_vpc_subnet.subnet_1.id
  }
}
resource "sbercloud_compute_instance" "ecs_m" {
  name              = "master"      
  image_id          = data.sbercloud_images_image.myimage.id
  flavor_id         = "s6.xlarge.2"
  key_pair = "KeyPair-rudometov"
  security_groups   = ["default", "arudometov-sg"]
  availability_zone = "ru-moscow-1a"
   
  system_disk_type = "SAS"
  system_disk_size = 40   
 
  network {
    uuid = sbercloud_vpc_subnet.subnet_1.id
  }
}  

