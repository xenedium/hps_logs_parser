resource "helm_release" "kube_prom_stack" {
  repository = "https://prometheus-community.github.io/helm-charts"
  chart      = "kube-prometheus-stack"

  name             = "kube-prom-stack"
  namespace        = "kube-prom-stack"
  create_namespace = true
}
