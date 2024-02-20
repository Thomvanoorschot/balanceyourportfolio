output "cluster_urn" {
  description = "The URN of the cluster"
  value       = digitalocean_database_cluster.postgres_cluster.urn
}