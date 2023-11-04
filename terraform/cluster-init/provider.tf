terraform {
  backend "s3" {
    key    = "terraform/cluster-init/terraform.tfstate"
    region = "us-east-1"
  }
}

provider "helm" {
  kubernetes {
    config_path = "config"
  }
}

provider "kubernetes" {
  config_path = "config"
}
