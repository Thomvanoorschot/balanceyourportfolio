resource "digitalocean_database_firewall" "example-fw" {
  cluster_id = digitalocean_database_cluster.postgres-example.id

  rule {
    type  = "ip_addr"
    value = "192.168.1.1"
  }
  rule {
    type  = "droplet"
    value = digitalocean_droplet.web.id
  }
}
