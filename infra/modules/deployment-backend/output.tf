output "service_name" {
  description = "The backend service name"
  value       = kubernetes_service.default.metadata[0].name
}