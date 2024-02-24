variable "do_variables" {
  description = "Digital ocean variables"
  type = map(string)
}
variable top_level_domains {
  description = "Top level domains to create records and pods for"
  type    = list(string)
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
variable "node_pool_node_count" {
  description = "Node pool node count of the k8s cluster"
  type = number
  default = 2
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

// Environment
variable "environment" {
  description = "The current environment"
  type = string
}