variable "digital_ocean_token" {
  description = "The token used to authenticate with DigitalOcean"
  type        = string
}
variable "project_name" {
  description = "The name of the project"
  type        = string
}
variable "project_description" {
  description = "The project description"
  type        = string
}
variable "environment" {
  description = "Environment of the cluster"
  type        = string
}

// Domains
variable frontend_domain {
  description = "The frontend domain"
  type    = string
}

// Deployments
// Backend
variable "backend_vars" {
  description = "Backend component variables"
  type        = map(string)
}
variable "backend_component_name" {
  description = "Backend component variables"
  type        = string
}
variable "backend_image" {
  description = "Image for the backend pod"
  type        = string
}
variable "backend_port" {
  description = "Port which the backend runs on"
  type        = number
}

// Frontend
variable "frontend_vars" {
  description = "Frontend component variables"
  type        = map(string)
}
variable "frontend_component_name" {
  description = "Frontend component variables"
  type        = string
}
variable "frontend_image" {
  description = "Image for the frontend pod"
  type        = string
}
variable "frontend_port" {
  description = "Port which the frontend runs on"
  type        = number
}

// LetsEncrypt
variable "letsencrypt_email" {
  description = "The email used for LetsEncrypt"
  type        = string
}
