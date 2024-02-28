resource "kubernetes_namespace" "cert-manager" {
  metadata {
    name = "cert-manager"
  }
}
resource "helm_release" "cert-manager" {
  name       = "cert-manager"
  repository = "https://charts.jetstack.io"
  chart      = "cert-manager"
  version    = "v1.14.3"
  namespace  = "cert-manager"
  timeout    = 120

  set {
    name  = "createCustomResource"
    value = "true"
  }
  set {
    name  = "installCRDs"
    value = "true"
  }
  depends_on = [
    kubernetes_namespace.cert-manager
  ]
}

resource "helm_release" "cluster-issuer" {
  name      = "cluster-issuer"
  chart     = "../../../helm/cluster-issuer"
  namespace = "cert-manager"
  depends_on = [
    helm_release.cert-manager,
    kubernetes_namespace.cert-manager
  ]
  set {
    name  = "letsencrypt_email"
    value = var.letsencrypt_email
  }
}
