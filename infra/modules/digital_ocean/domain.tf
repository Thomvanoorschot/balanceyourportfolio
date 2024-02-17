resource "digitalocean_domain" "default" {
  name       = "example.com"
  ip_address = digitalocean_droplet.foo.ipv4_address
}