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

variable "supabase_region" {
  type        = string
  description = "Supabase DB region"
}

variable "supabase_org_id" {
  type        = string
  description = "Supabase organization ID"
}

variable "key_vault_name" {
  type        = string
  description = "Name of the Azure Key Vault"
}

variable "akv_resource_group_name" {
  type        = string
  description = "Resource group name where the Key Vault is located"
}

variable "akv_supabase_secret_name" {
  type        = string
  description = "Name of the secret in Key Vault that contains the Supabase access token"
}
variable "environment_variables" {
  type        = map(string)
  description = "Environment variables for the application"
  default     = {}
}
