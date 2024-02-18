resource "digitalocean_database_user" "db_user" {
  cluster_id = digitalocean_database_cluster.postgres_cluster.id
  name       = var.postgres_database_user
}
