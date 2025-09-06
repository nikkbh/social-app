data "azurerm_key_vault" "akv" {
  name                = var.key_vault_name
  resource_group_name = var.akv_resource_group_name
}

data "azurerm_key_vault_secret" "supabase_access_token" {
  name         = var.akv_supabase_secret_name
  key_vault_id = data.azurerm_key_vault.akv.id
}

data "azurerm_key_vault_secret" "dbusername" {
  name         = "username"
  key_vault_id = data.azurerm_key_vault.akv.id
}
data "azurerm_key_vault_secret" "dbpassword" {
  name         = "password"
  key_vault_id = data.azurerm_key_vault.akv.id
}
