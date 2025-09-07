variable "project_name" {
  type        = string
  description = "Project name"
}

variable "resource_group_name" {
  type        = string
  description = "Resource group name"
}

variable "location" {
  type        = string
  description = "Resource group location"
}

variable "image_name" {
  type        = string
  description = "Docker image name"
}

variable "app_service_plan_sku" {
  type        = string
  description = "App Service Plan SKU"
  default     = "B1" # Basic tier
}

variable "app_service_plan_os_type" {
  type        = string
  description = "App Service Plan OS type"
  default     = "Linux"
}

variable "acr_sku" {
  type        = string
  description = "Azure Container Registry SKU"
  default     = "Standard"
}

variable "tags" {
  type        = map(string)
  description = "Tags to apply to resources"
}
variable "environment_variables" {
  type        = map(string)
  description = "Environment variables for the application"
  default     = {}
}
