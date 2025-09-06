terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "4.43.0"
    }
    supabase = {
      source  = "supabase/supabase"
      version = "~> 1.0"
    }
  }
  required_version = ">= 1.1.0"
}
provider "azurerm" {
  features {}
}

provider "supabase" {
  access_token = data.azurerm_key_vault_secret.supabase_access_token.value
}
