terraform {
  cloud {
    organization = "etfinsight"

    workspaces {
      name = "staging"
    }
  }

  required_version = ">= 1.1.2"
}