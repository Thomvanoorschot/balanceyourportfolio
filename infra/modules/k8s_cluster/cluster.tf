resource "digitalocean_kubernetes_cluster" "default" {
  name   = var.cluster_name
  region = var.region
  version = var.k8s_version

  node_pool {
    name       = "worker-pool"
    size       = var.node_pool_size
    node_count = var.node_pool_node_count
  }
}