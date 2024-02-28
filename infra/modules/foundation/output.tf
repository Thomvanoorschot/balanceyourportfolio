output "project_id" {
  description = "The ID of the project"
  value       = digitalocean_project.balanceyourportfolio.id
}
output "loadbalancer_id" {
  description = "The ID of the loadbalancer"
  value       = digitalocean_loadbalancer.ingress_load_balancer.id
}
output "loadbalancer_urn" {
  description = "The URN of the loadbalancer"
  value       = digitalocean_loadbalancer.ingress_load_balancer.urn
}
output "frontend_domain_urn" {
  description = "The URN of the frontend domain"
  value       = digitalocean_domain.frontend_domain.urn
}