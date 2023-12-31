terraform {
  backend "s3" {
    key    = "terraform/digitalocean/terraform.tfstate"
    region = "us-east-1"
  }

  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {
  token = var.do_token
}
