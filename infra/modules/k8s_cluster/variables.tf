variable "do_variables" {
  description = "Digital ocean variables"
  type = map(string)
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
  default = "s-2vcpu-2gb"
}
variable "node_pool_node_count" {
  description = "Node pool node count of the k8s cluster"
  type = number
  default = 2
}

// Environment
variable "environment" {
  description = "The current environment"
  type = string
}