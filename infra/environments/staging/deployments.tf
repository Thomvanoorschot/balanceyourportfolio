module "deployment_backend" {
  source = "../../modules/deployment-backend"
  app_vars = var.backend_vars
  component_name = var.backend_component_name
  image = var.backend_image
  port = var.backend_port
}
module "deployment_frontend" {
  source = "../../modules/deployment-frontend"
  app_vars = merge(var.frontend_vars, {
    KINDE_REDIRECT_URL= "https://localhost:5173/api/auth/kinde_callback"
    KINDE_POST_LOGOUT_REDIRECT_URL= "https://localhost:5173/portfolio/overview"
    KINDE_POST_LOGIN_REDIRECT_URL= "https://localhost:5173/portfolio/overview"
    GRPC_API_URL= "${module.deployment_backend.service_name}:${var.backend_port}"
  })
  component_name = var.frontend_component_name
  image = var.frontend_image
  port = var.frontend_port
}