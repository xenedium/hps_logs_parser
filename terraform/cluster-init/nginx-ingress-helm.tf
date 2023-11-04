resource "helm_release" "nginx_ingress" {
  repository = "https://kubernetes.github.io/ingress-nginx"
  chart      = "ingress-nginx"

  name             = "ingress-nginx"
  namespace        = "ingress-nginx"
  create_namespace = true
}


#data "kubernetes_service" "nginx_ingress" {
#  metadata {
#    name = "ingress-nginx"
#    namespace = "ingress-nginx"
#  }
#}

#output "nginx_ingress_ip" {
#  value = helm_release.nginx_ingress.status.
#}
