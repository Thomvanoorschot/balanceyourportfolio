// Digital Ocean
variable "do_variables" {
  description = "Digital ocean variables"
  type = map(string)
}

// Project
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
variable "project_resources" {
  description = "The resources for the project"
  type = list(string)
}
variable "region" {
  description = "Region of the infrastructure"
  type = string
  default = "ams3"
}

// DNS
variable frontend_domain {
  description = "The frontend domain"
  type    = string
}

// Loadbalancer
variable "loadbalancer_algorithm" {
  description = "The algorithm the loadbalancer users"
  type = string
  default = "round_robin"
}
variable "loadbalancer_size" {
  description = "Size of the loadbalancer"
  type = string
  default = "lb-small"
}

// Certificate manager
variable "letsencrypt_email" {
  description = "The email for LetsEncrypt"
  type = string
}
