resource "helm_release" "nginx_ingress_chart" {
  name       = "nginx-ingress-controller"
  namespace  = kubernetes_namespace.balanceyourportfolio.metadata[0].name
  repository = "https://charts.bitnami.com/bitnami"
  chart      = "nginx-ingress-controller"
  set {
    name  = "service.type"
    value = "LoadBalancer"
  }
  set {
    name  = "service.annotations.kubernetes\\.digitalocean\\.com/load-balancer-id"
    value = var.loadbalancer_id
  }
}

resource "kubernetes_ingress_v1" "default_cluster_ingress" {
  depends_on = [
    helm_release.nginx_ingress_chart,
  ]
  metadata {
    name      = "${var.cluster_name}-ingress"
    namespace = kubernetes_namespace.balanceyourportfolio.metadata[0].name
    annotations = {
      "kubernetes.io/ingress.class"          = "nginx"
      "ingress.kubernetes.io/rewrite-target" = "/"
      "cert-manager.io/cluster-issuer"       = "letsencrypt-production"
    }
  }
  spec {
    rule {
      host = var.frontend_domain
      http {
        path {
          backend {
            service {
              name = var.frontend_service_name
              port {
                number = var.frontend_port
              }
            }
          }
          path = "/"
        }
      }
    }
    tls {
      hosts       = [var.frontend_domain]
      secret_name = "frontend-tls"
    }
  }
}
