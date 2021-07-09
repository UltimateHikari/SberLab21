data "sbercloud_images_image" "myimage" {
  name        = "Ubuntu 20.04 server 64bit"
  most_recent = true
}

resource "sbercloud_compute_instance" "ecs_1" {
  name              = "arudometov-1"
  image_id          = data.sbercloud_images_image.myimage.id
  flavor_id         = "s6.small.1"
  key_pair = "KeyPair-rudometov"
  security_groups   = ["default", "arudometov-sg"]
  availability_zone = "ru-moscow-1a"
  user_data = "#!/bin/bash\n echo 'hello' > /root/hello.txt"

  system_disk_type = "SAS"
  system_disk_size = 40

  network {
    uuid = sbercloud_vpc_subnet.subnet_1.id
  }
}

resource "sbercloud_compute_instance" "ecs_2" {
  name              = "arudometov-2"
  image_id          = data.sbercloud_images_image.myimage.id
  flavor_id         = "s6.small.1"
  key_pair = "KeyPair-rudometov"
  security_groups   = ["default", "arudometov-sg"]
  availability_zone = "ru-moscow-1a"
  user_data = file("./dockerfront.sh")
  system_disk_type = "SAS"
  system_disk_size = 40

  network {
    uuid = sbercloud_vpc_subnet.subnet_1.id
  }
}

