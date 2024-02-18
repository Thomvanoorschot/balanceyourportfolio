resource "digitalocean_loadbalancer" "ingress_load_balancer" {
  name   = "${var.cluster_name}_load_balancer"
  region = var.region

  forwarding_rule {
    entry_port     = 80
    entry_protocol = "http"

    target_port     = 80
    target_protocol = "http"
  }

  healthcheck {
    port     = 22
    protocol = "tcp"
  }
}
