resource "digitalocean_database_connection_pool" "postgres_connection_pool" {
  cluster_id = digitalocean_database_cluster.postgres_cluster.id
  name       = "${var.cluster_name}_${var.environment}_postgres_connection_pool"
  mode       = var.postgres_connection_pool_mode
  size       = var.postgres_connection_pool_size
  db_name    = digitalocean_database_db.database_etfinsight.name
  user       = var.postgres_database_user
}
