resource "digitalocean_project" "etfinsight" {
  name        = var.project_name
  description = var.project_description
  environment = var.environment
}