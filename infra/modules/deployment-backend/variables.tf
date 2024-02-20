variable "project_id" {
  description = "name of the gcp project (Required)"
  type        = string
}

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

## Variables with defaults
variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "profile"
}

variable "namespace" {
  description = "Namespace where the resources should be provisioned in."
  type        = string
  default     = "profile"
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
  default     = "1600m"
}

variable "cpu_limit" {
  type        = string
  description = "The CPU limit for the pod"
  default     = "2000m"
}

variable "memory_request" {
  type        = string
  description = "The memory request for the pod"
  default     = "800Mi"
}

variable "memory_limit" {
  type        = string
  description = "The memory limit for the pod"
  default     = "1600Mi"
}
