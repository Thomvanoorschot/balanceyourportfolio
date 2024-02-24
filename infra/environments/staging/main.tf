module "foundation" {
  source      = "../../modules/foundation"
  environment = var.environment
  do_variables = {
    TOKEN : var.digital_ocean_token
  }
  project_description = var.project_description
  project_name        = var.project_name
  project_resources = var.project_resources
}