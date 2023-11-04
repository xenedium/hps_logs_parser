resource "kubernetes_manifest" "hps_logs_parser_staging" {
  manifest = {
    "apiVersion" = "argoproj.io/v1alpha1"
    "kind"       = "Application"
    "metadata" = {
      "name"      = "hps-logs-parser-staging"
      "namespace" = "argo-cd"
    }
    "spec" = {
      "destination" = {
        "name"      = ""
        "namespace" = "hps-logs-parser-staging"
        "server"    = "https://kubernetes.default.svc"
      }
      "project" = "default"
      "source" = {
        "path"           = "kubernetes/staging"
        "repoURL"        = "https://github.com/xenedium/hps_logs_parser"
        "targetRevision" = "HEAD"
        "directory" = {
          "jsonnet" = {
            "tlas" = []
            "extVars" = [
              {
                "name"  = "host"
                "value" = var.ingress_staging_host
              }
            ]
          }
        }

      }
      "sources" = []
      "syncPolicy" = {
        "automated" = {
          "prune"    = true
          "selfHeal" = false
        }
        "syncOptions" = [
          "CreateNamespace=true",
        ]
      }
    }
  }
}
