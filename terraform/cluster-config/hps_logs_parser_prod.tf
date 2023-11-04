resource "kubernetes_manifest" "hps_logs_parser_production" {
  manifest = {
    "apiVersion" = "argoproj.io/v1alpha1"
    "kind"       = "Application"
    "metadata" = {
      "name"      = "hps-logs-parser-production"
      "namespace" = "argo-cd"
    }
    "spec" = {
      "destination" = {
        "name"      = ""
        "namespace" = "hps-logs-parser-production"
        "server"    = "https://kubernetes.default.svc"
      }
      "project" = "default"
      "source" = {
        "path"           = "kubernetes/production"
        "repoURL"        = "https://github.com/xenedium/hps_logs_parser"
        "targetRevision" = "HEAD"
        "directory" = {
          "jsonnet" = {
            "tlas" = []
            "extVars" = [
              {
                "name"  = "host"
                "value" = var.ingress_prod_host
              }
            ]
          }
        }
      }
      "sources" = []
      "syncPolicy" = {
        "syncOptions" = [
          "CreateNamespace=true",
        ]
      }
    }
  }
}
