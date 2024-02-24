locals {
  labels = merge({ app = var.project_name }, { component = var.component_name })
}

resource "kubernetes_service_account" "default" {
  metadata {
    name      = "k8s-${var.component_name}"
    namespace = var.namespace
  }
}

resource "random_id" "config_map_suffix" {
  byte_length = 8
  keepers     = var.app_vars
}

resource "kubernetes_config_map" "default" {
  metadata {
    name      = "${var.component_name}-${random_id.config_map_suffix.hex}"
    namespace = var.namespace
  }
  data = var.app_vars
}

resource "kubernetes_deployment" "default" {
  metadata {
    name      = var.component_name
    namespace = var.namespace
    labels    = local.labels
  }

  spec {
    selector {
      match_labels = local.labels
    }


    strategy {
      rolling_update {
        max_surge       = 1
        max_unavailable = 1
      }
    }

    template {
      metadata {
        labels = local.labels
      }

      spec {
        image_pull_secrets {
          name = "balanceyourportfolio"
        }
        restart_policy       = "Always"
        service_account_name = kubernetes_service_account.default.metadata[0].name

        container {
          name              = var.component_name
          image             = var.image
          image_pull_policy = "Always"
          command           = ["/app"]

          resources {
            requests = {
              cpu    = var.cpu_request
              memory = var.memory_request
            }
            limits = {
              cpu    = var.cpu_limit
              memory = var.memory_limit
            }
          }

          env_from {
            config_map_ref {
              name = kubernetes_config_map.default.metadata[0].name
            }
          }

          port {
            container_port = var.port
          }

          liveness_probe {
            initial_delay_seconds = "30"
            period_seconds        = "10"
            timeout_seconds       = "21"
            success_threshold     = "1"
            failure_threshold     = "5"
            http_get {
              path   = "healthz"
              port   = var.port
              scheme = "HTTP"
            }
          }

          readiness_probe {
            initial_delay_seconds = "30"
            period_seconds        = "10"
            timeout_seconds       = "21"
            success_threshold     = "1"
            failure_threshold     = "5"
            http_get {
              path   = "healthz"
              port   = var.port
              scheme = "HTTP"
            }
          }
        }
      }
    }
  }
}


resource "kubernetes_service" "default" {
  metadata {
    name      = format("%s-svc", var.component_name)
    namespace = var.namespace
  }

  spec {
    selector = local.labels
    type     = "ClusterIP"

    port {
      name        = "http"
      port        = var.port
      target_port = tostring(var.port)
    }
  }
}