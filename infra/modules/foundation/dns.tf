resource "digitalocean_domain" "frontend_domain" {
  name = var.frontend_domain
}

resource "digitalocean_record" "a_records" {
  domain = var.frontend_domain
  type   = "A"
  ttl = 60
  name   = "@"
  value  = digitalocean_loadbalancer.ingress_load_balancer.ip
  depends_on = [
    digitalocean_domain.frontend_domain,
  ]
}

resource "digitalocean_record" "cname_redirects" {
  domain = var.frontend_domain
  type   = "CNAME"
  ttl = 60
  name   = "www"
  value  = "@"
  depends_on = [
    digitalocean_domain.frontend_domain,
  ]
}
