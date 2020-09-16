provider "google" {
  credentials = file("account.json")
  project     = "my-project-id"
  region      = "us-central1"
}

variable "auth_system_id" {
    type = string
    description = "GCP Auth System ID"
    default = "silicon-banana-123456"
}

variable "identity_ids" {
    type = list(string)
    description = "A list of GCP user Ids"
    default     = []
}

provider "cloudknox" {
    profile = "default" // See docs/guides/authentication_reference.md
}

data "cloudknox_role_policy" "user-activity-gcp-role" {
    name = "user-activity-gcp-role"
    output_path = "./"
    auth_system_info = {
         id = var.auth_system_id,
         type = "GCP"
     }
    identity_type = "USER" [
    identity_ids = var.identity_ids
]

    filter_history_days = 90
}