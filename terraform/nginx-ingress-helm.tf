resource "helm_release" "nginx_ingress" {
  repository = "https://kubernetes.github.io/ingress-nginx"
  chart      = "ingress-nginx"

  name             = "ingress-nginx"
  namespace        = "ingress-nginx"
  create_namespace = true
}
