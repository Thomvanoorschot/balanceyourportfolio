module "foundation" {
  source      = "../../modules/foundation"
  environment = var.environment
  do_variables = {
    TOKEN : var.digital_ocean_token
  }
  project_description = var.project_description
  project_name        = var.project_name
}

resource "digitalocean_project_resources" "resources" {
  project = module.foundation.project_id
  resources = [
    module.k8s_cluster.cluster_urn,
    module.k8s_cluster.loadbalancer_urn,
    module.postgres.cluster_urn,
  ]
}