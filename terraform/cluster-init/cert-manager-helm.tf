resource "helm_release" "cert_manager" {
  repository = "https://charts.jetstack.io"
  chart      = "cert-manager"

  namespace        = "cert-manager"
  name             = "cert-manager"
  create_namespace = true

  set {
    name  = "installCRDs"
    value = true
  }
}
