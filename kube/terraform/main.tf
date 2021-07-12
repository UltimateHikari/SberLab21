terraform {
  required_providers {
    sbercloud = {
      source = "sbercloud-terraform/sbercloud"
      version = "1.3.0"
    }
  }
}

provider "sbercloud" {}
