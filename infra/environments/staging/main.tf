module "foundation" {
  source = "../../modules/foundation"
  do_variables = {
    TOKEN: var.digital_ocean_token
  }
  environment = var.environment
  project_description = var.project_description
  project_name = var.project_name
}