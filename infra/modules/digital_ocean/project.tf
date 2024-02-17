resource "digitalocean_project" "playground" {
  name        = "playground"
  description = "A project to represent development resources."
  purpose     = "Web Application"
  environment = "Development"
}