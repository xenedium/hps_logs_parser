resource "digitalocean_kubernetes_cluster" "logparser" {
  name    = "logparser"
  region  = "fra1"
  version = "1.28.2-do.0"

  node_pool {
    name       = "main-pool"
    size       = "s-1vcpu-2gb"
    node_count = 3
  }
}

resource "local_file" "kubeconfig" {
  content  = digitalocean_kubernetes_cluster.logparser.kube_config.0.raw_config
  filename = "config"
}
