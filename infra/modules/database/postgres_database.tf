resource "digitalocean_database_db" "database-example" {
  cluster_id = digitalocean_database_cluster.postgres-example.id
  name       = "foobar"
}