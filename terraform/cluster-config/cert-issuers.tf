resource "kubernetes_manifest" "letsencrypt-staging" {
  manifest = {
    "apiVersion" = "cert-manager.io/v1"
    "kind"       = "Issuer"
    "metadata" = {
      "name"      = "letsencrypt-staging"
      "namespace" = "hps-logs-parser-staging"
    }
    "spec" = {
      "acme" = {
        "email" = var.acme_email
        "privateKeySecretRef" = {
          "name" = "letsencrypt-staging"
        }
        "server" = "https://acme-staging-v02.api.letsencrypt.org/directory"
        "solvers" = [
          {
            "http01" = {
              "ingress" = {
                "ingressClassName" = "nginx"
              }
            }
          },
        ]
      }
    }
  }
}

resource "kubernetes_manifest" "letsencrypt-prod" {
  manifest = {
    "apiVersion" = "cert-manager.io/v1"
    "kind"       = "Issuer"
    "metadata" = {
      "name"      = "letsencrypt-prod"
      "namespace" = "hps-logs-parser-production"
    }
    "spec" = {
      "acme" = {
        "email" = var.acme_email
        "privateKeySecretRef" = {
          "name" = "letsencrypt-prod"
        }
        "server" = "https://acme-v02.api.letsencrypt.org/directory"
        "solvers" = [
          {
            "http01" = {
              "ingress" = {
                "ingressClassName" = "nginx"
              }
            }
          },
        ]
      }
    }
  }
}
