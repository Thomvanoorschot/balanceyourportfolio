terraform {
  cloud {
    organization = "balanceyourportfolio"

    workspaces {
      name = "staging"
    }
  }

  required_version = ">= 1.1.2"
}