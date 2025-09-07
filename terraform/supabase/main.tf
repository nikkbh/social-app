resource "supabase_project" "db" {
  organization_id   = var.supabase_org_id
  name              = "${var.project_name}-db"
  database_password = data.azurerm_key_vault_secret.dbpassword.value
  region            = var.supabase_region
  lifecycle {
    ignore_changes = [database_password]
  }
}
