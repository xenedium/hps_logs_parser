resource "helm_release" "argocd" {
  repository = "https://argoproj.github.io/argo-helm"
  chart      = "argo-cd"

  name             = "argo-cd"
  namespace        = "argo-cd"
  create_namespace = true

  set {
    name  = "crds.keep"
    value = false
  }
}
