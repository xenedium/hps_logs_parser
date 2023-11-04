terraform {
  backend "s3" {
    key    = "terraform/cluster-config/terraform.tfstate"
    region = "us-east-1"
  }
}

provider "kubernetes" {
  config_path = "config"
}


data "kubernetes_service_v1" "ingress-nginx-external-ip" {
  metadata {
    name = "ingress-nginx-controller"
    namespace = "ingress-nginx"
  }
}

output "ingress-nginx-external-ip-output" {
  value = data.kubernetes_service_v1.ingress-nginx-external-ip.status.0.load_balancer.0.ingress.0.ip
}
