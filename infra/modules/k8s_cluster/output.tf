output "cluster_urn" {
  description = "The URN of the cluster"
  value       = digitalocean_kubernetes_cluster.default.urn
}
output "loadbalancer_urn" {
  description = "The URN of the loadbalancer"
  value       = digitalocean_loadbalancer.ingress_load_balancer.urn
}