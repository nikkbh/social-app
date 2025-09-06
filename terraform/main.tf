# Creates an Azure Resource Group
resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location

  tags = var.tags
}

# Creates an Azure App Service Plan for hosting web apps
resource "azurerm_service_plan" "app_service_plan" {
  name                = "${var.resource_group_name}-plan"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  sku_name            = var.app_service_plan_sku
  os_type             = var.app_service_plan_os_type

  tags = var.tags
}

# Creates an Azure Container Registry for storing Docker images
resource "azurerm_container_registry" "acr" {
  name                = "${var.project_name}ACR"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  sku                 = var.acr_sku
  admin_enabled       = true
  tags                = var.tags
}

# Deploys a Linux Web App using a Docker image from the Container Registry
resource "azurerm_linux_web_app" "web_app" {
  name                = "${var.resource_group_name}-app"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
  service_plan_id     = azurerm_service_plan.app_service_plan.id

  site_config {
    always_on = true
    application_stack {
      docker_image_name        = "${var.image_name}:latest"
      docker_registry_url      = "https://${azurerm_container_registry.acr.login_server}"
      docker_registry_username = azurerm_container_registry.acr.admin_username
      docker_registry_password = azurerm_container_registry.acr.admin_password
    }
  }

  tags = var.tags
}

# Create a supabase project resource
resource "supabase_project" "db" {
  organization_id   = var.supabase_org_id
  name              = "${var.project_name}-db"
  database_password = data.azurerm_key_vault_secret.dbpassword.value
  region            = var.supabase_region
  lifecycle {
    ignore_changes = [database_password]
  }
}
