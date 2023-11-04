resource "kubernetes_namespace_v1" "hps_logs_parser_staging" {
  metadata {
    name = "hps-logs-parser-staging"
  }
}

resource "kubernetes_namespace_v1" "hps_logs_parser_production" {
  metadata {
    name = "hps-logs-parser-production"
  }
}
