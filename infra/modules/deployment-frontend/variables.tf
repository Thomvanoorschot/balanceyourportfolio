variable "component_name" {
  description = "Name of the component/microservice e.g. core"
  type        = string
}

variable "app_vars" {
  description = "Application specific variables (Required)"
  type        = map(string)
}

variable "image" {
  description = "image full digest (Required)"
  type        = string
}

variable "port" {
  description = "Port which the pod runs on"
  type        = number
}

## Variables with defaults
variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "balanceyourportfolio"
}

variable "namespace" {
  description = "Namespace where the resources should be provisioned in."
  type        = string
  default     = "balanceyourportfolio"
}

variable "min_pods" {
  description = "minimal number of pods (Required)"
  type        = string
  default     = 1
}

variable "max_pods" {
  description = "maximal number of pods (Required)"
  type        = string
  default     = 1
}

variable "cpu_request" {
  type        = string
  description = "The CPU request for the pod"
  default     = "200m"
}

variable "cpu_limit" {
  type        = string
  description = "The CPU limit for the pod"
  default     = "500m"
}

variable "memory_request" {
  type        = string
  description = "The memory request for the pod"
  default     = "250Mi"
}

variable "memory_limit" {
  type        = string
  description = "The memory limit for the pod"
  default     = "1000Mi"
}
