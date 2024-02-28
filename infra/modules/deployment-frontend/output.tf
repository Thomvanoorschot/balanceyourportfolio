output "service_name" {
  description = "The frontend service name"
  value       = kubernetes_service.default.metadata[0].name
}