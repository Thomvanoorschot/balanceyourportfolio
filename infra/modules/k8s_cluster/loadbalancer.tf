resource "digitalocean_loadbalancer" "ingress_load_balancer" {
  name   = "${var.cluster_name}-load-balancer"
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
