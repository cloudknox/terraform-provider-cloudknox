provider "aws" {
  region     = "us-west-2"
  access_key = "my-access-key"
  secret_key = "my-secret-key"
}

variable "auth_system_id" {
    type = string
    description = "AWS Auth System ID"
    default = "123456789012"
}

variable "identity_ids" {
    type = list(string)
    description = "A list of ARNs that represent roles"
    default     = []
}

provider "cloudknox" {
    profile = "default" // See docs/guides/authentication_reference.md
}

data "cloudknox_role_policy" "role-activity-aws-policy" {
    name = "role-activity-large-aws-policy"
    output_path = "./"
    auth_system_info = {
        id = var.auth_system_id,
        type = "AWS"
    }
    identity_type = "ROLE" [
    identity_ids = var.identity_ids
]

    filter_history_days = 90
}