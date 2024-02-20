module "postgres" {
  source = "../../modules/postgres"
  do_variables = {
    TOKEN: var.digital_ocean_token
  }
  cluster_name = var.project_name
  environment = var.environment
}