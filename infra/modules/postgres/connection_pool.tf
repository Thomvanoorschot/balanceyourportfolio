resource "digitalocean_database_connection_pool" "postgres_connection_pool" {
  cluster_id = digitalocean_database_cluster.postgres_cluster.id
  name       = "${var.cluster_name}-${var.environment}-postgres-connection-pool"
  mode       = var.connection_pool_mode
  size       = var.connection_pool_size
  db_name    = digitalocean_database_db.database_balanceyourportfolio.name

  depends_on = [
    digitalocean_database_db.database_balanceyourportfolio
  ]
}
