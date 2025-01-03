resource "digitalocean_loadbalancer" "ingress_load_balancer" {
  name      = "${var.project_name}-lb"
  region    = var.region
  algorithm = var.loadbalancer_algorithm
  size      = var.loadbalancer_size

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

  lifecycle {
    ignore_changes = [
      forwarding_rule,
    ]
  }
}
