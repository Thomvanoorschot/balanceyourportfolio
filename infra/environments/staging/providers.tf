provider "kubernetes" {
  host  = module.k8s_cluster.cluster_endpoint
  token = module.k8s_cluster.cluster_token
  cluster_ca_certificate = base64decode(
    module.k8s_cluster.cluster_ca_certificate
  )
}
provider "helm" {
  kubernetes {
    host  = module.k8s_cluster.cluster_endpoint
    token = module.k8s_cluster.cluster_token
    cluster_ca_certificate = base64decode(
      module.k8s_cluster.cluster_ca_certificate
    )
  }
}