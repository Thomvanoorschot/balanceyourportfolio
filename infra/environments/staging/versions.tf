terraform {
  required_version = ">= 1.2"
  required_providers {
    digitalocean = {
      source = "digitalocean/digitalocean"
      version = "2.34.1"
    },
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.26.0"
    }
  }
}