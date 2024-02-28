module "foundation" {
  source      = "../../modules/foundation"
  environment = var.environment
  do_variables = {
    TOKEN : var.digital_ocean_token
  }
  project_description = var.project_description
  project_name        = var.project_name
  project_resources = [
    module.k8s_cluster.cluster_urn,
    module.postgres.cluster_urn,
    module.foundation.loadbalancer_urn,
    module.foundation.frontend_domain_urn
  ]
  letsencrypt_email = var.letsencrypt_email
  frontend_domain = var.frontend_domain
}