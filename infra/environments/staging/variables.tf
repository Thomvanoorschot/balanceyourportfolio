variable "digital_ocean_token" {
  description = "The token used to authenticate with DigitalOcean"
  type = string
}
variable "project_name" {
  description = "The name of the project"
  type = string
}
variable "project_description" {
  description = "The project description"
  type = string
}
variable "environment" {
  description = "Environment of the cluster"
  type = string
}
