variable "project_name" {
  type        = string
  description = "Project name"
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
