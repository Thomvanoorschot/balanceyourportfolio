resource "digitalocean_project" "balanceyourportfolio" {
  name        = var.project_name
  description = var.project_description
  environment = var.environment

  resources = var.project_resources
}
