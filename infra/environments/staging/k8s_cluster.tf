module "k8s_cluster" {
  source = "../../modules/k8s_cluster"
  do_variables = {
    TOKEN : var.digital_ocean_token
  }
  cluster_name      = var.project_name
  environment       = var.environment
  top_level_domains = var.top_level_domains
}