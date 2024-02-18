resource "digitalocean_database_cluster" "postgres_cluster" {
  name       = "${var.cluster_name}_postgres_cluster"
  engine     = var.postgres_cluster_engine
  version    = var.postgres_cluster_version
  size       = var.postgres_cluster_size
  region     = var.region
  node_count = var.postgres_cluster_node_count
}