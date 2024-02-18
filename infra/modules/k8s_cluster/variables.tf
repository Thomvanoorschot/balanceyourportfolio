// Cluster
variable "cluster_name" {
  description = "Name of the k8s cluster"
  type = string
}
variable "region" {
  description = "Region of the k8s cluster"
  type = string
  default = "AMS3"
}
variable "k8s_version" {
  description = "Version of the k8s cluster"
  type = string
  default = "1.22.8-do.1"
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

// Postgres Cluster
variable "postgres_cluster_engine" {
  description = "The database engine"
  type = string
  default = "pg"
}
variable "postgres_cluster_version" {
  description = "The version of postgres"
  type = string
  default = "15"
}
variable "postgres_cluster_size" {
  description = "The size of the postgres cluster"
  type = string
  default = "db-s-1vcpu-1gb"
}
variable "postgres_cluster_node_count" {
  description = "The postgres cluster node count"
  type = number
  default = 1
}

// Postgres connection pool
variable "postgres_connection_pool_mode" {
  description = "The postgres connection pool mode"
  type = string
  default = "transaction"
}
variable "postgres_connection_pool_size" {
  description = "The postgres connection pool size"
  type = number
  default = 20
}

// Environment
variable "environment" {
  description = "The current environment"
  type = string
}

// Postgres database user
variable "postgres_database_user" {
  description = "The database user"
  type = string
  default = "etfinsight"
}