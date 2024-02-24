output "cluster_urn" {
  description = "The URN of the cluster"
  value       = digitalocean_kubernetes_cluster.default.urn
}
output "loadbalancer_urn" {
  description = "The URN of the loadbalancer"
  value       = digitalocean_loadbalancer.ingress_load_balancer.urn
}
output "cluster_endpoint" {
  description = "The endpoint of the cluster"
  value       = digitalocean_kubernetes_cluster.default.endpoint
}
output "cluster_token" {
  description = "The token of the cluster"
  value       = digitalocean_kubernetes_cluster.default.kube_config[0].token
}
output "cluster_ca_certificate" {
  description = "The CA certificate of the cluster"
  value       = digitalocean_kubernetes_cluster.default.kube_config[0].cluster_ca_certificate
}