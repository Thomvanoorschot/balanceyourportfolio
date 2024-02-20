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

// Postgres Cluster
variable "cluster_engine" {
  description = "The database engine"
  type = string
  default = "pg"
}
variable "cluster_version" {
  description = "The version of postgres"
  type = string
  default = "15"
}
variable "cluster_size" {
  description = "The size of the postgres cluster"
  type = string
  default = "db-s-1vcpu-1gb"
}
variable "cluster_node_count" {
  description = "The postgres cluster node count"
  type = number
  default = 1
}

// Postgres connection pool
variable "connection_pool_mode" {
  description = "The postgres connection pool mode"
  type = string
  default = "transaction"
}
variable "connection_pool_size" {
  description = "The postgres connection pool size"
  type = number
  default = 20
}

// Environment
variable "environment" {
  description = "The current environment"
  type = string
}