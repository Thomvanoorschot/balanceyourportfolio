module "k8s_cluster" {
  source = "../../modules/k8s_cluster"
  cluster_name = var.project_name
  environment = var.environment
}