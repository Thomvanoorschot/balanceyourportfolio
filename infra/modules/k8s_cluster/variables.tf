// Digital Ocean
variable "do_variables" {
  description = "Digital ocean variables"
  type = map(string)
}
variable "loadbalancer_id" {
  description = "The id of the loadbalancer"
  type = string
}

// Cluster
variable "cluster_name" {
  description = "Name of the k8s cluster"
  type = string
}
variable "region" {
  description = "Region of the k8s cluster"
  type = string
  default = "ams3"
}
variable "k8s_version" {
  description = "Version of the k8s cluster"
  type = string
  default = "1.26.13-do.0"
}
variable "node_pool_size" {
  description = "Size of the node pool of the k8s cluster"
  type = string
  default = "s-1vcpu-2gb"
}

variable min_nodes {
  description = "The minimum number of nodes in the default pool"
  type        = number
  default     = 1
}

variable max_nodes {
  description = "The maximum number of nodes in the default pool"
  type        = number
  default     = 3
}

// Domains
variable frontend_domain {
  description = "The frontend domain"
  type    = string
}

variable frontend_port {
  description = "The frontend port"
  type    = number
}

// Environment
variable "environment" {
  description = "The current environment"
  type = string
}

// Services
variable frontend_service_name {
  description = "The name of the frontend service"
  type    = string
}