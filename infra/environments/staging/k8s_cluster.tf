module "k8s_cluster" {
  source = "../../modules/k8s_cluster"
  do_variables = {
    TOKEN : var.digital_ocean_token
  }
  cluster_name    = var.project_name
  environment     = var.environment
  frontend_domain = var.frontend_domain
  loadbalancer_id = module.foundation.loadbalancer_id
  frontend_service_name = module.deployment_frontend.service_name
  frontend_port = var.frontend_port
  min_nodes = 3
}