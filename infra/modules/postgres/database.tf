resource "digitalocean_database_db" "database_balanceyourportfolio" {
  cluster_id = digitalocean_database_cluster.postgres_cluster.id
  name       = "${var.cluster_name}-${var.environment}-db"

  depends_on = [
    digitalocean_database_cluster.postgres_cluster
  ]
}