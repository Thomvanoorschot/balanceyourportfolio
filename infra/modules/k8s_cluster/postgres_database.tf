resource "digitalocean_database_db" "database_etfinsight" {
  cluster_id = digitalocean_database_cluster.postgres_cluster.id
  name       = "${var.cluster_name}_${var.environment}_db"
}