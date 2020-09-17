provider "azurerm" {
  features {} 
}

variable "auth_system_id" {
    type = string
    description = "Azure Auth System ID"
    default = "12abcd34-56e7-890f-gh12-34i5678901jk"
}

variable "identity_ids" {
    type = list(string)
    description = "A list of Azure user Ids"
    default     = []
}

variable "params_scope" {
    type = string
    description = "Azure scope of permission"
    default     = "/subscriptions/12abcd34-56e7-890f-gh12-34i5678901jk"
}

provider "cloudknox" {
    profile = "default" // See docs/guides/authentication_reference.md
}

resource "cloudknox_role_policy" "app-activity-azure-role" {
    name = "app-activity-azure-role"
    output_path = "./"
    auth_system_info = {
         id = var.auth_system_id,
         type = "AZURE"
     }
    identity_type = "USER"
    identity_ids = var.identity_ids
    filter_history_days = 90
    filter_preserve_reads = true
    request_params_scope = var.params_scope
}