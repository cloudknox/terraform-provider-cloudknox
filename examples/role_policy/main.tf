provider "cloudknox" {
    profile = "default" // ENSURE A DEFAULT PROFILE EXISTS IN YOUR DEFAULT CREDENTIALS FILE
}

data "cloudknox_role_policy" "user-activity-aws-policy" {
    name = "user-activity-large-aws-policy"
    output_path = "./"
    auth_system_info = {
        id = "<AWS ID AUTH SYSTEM ID HERE>" // ENTER AUTH SYSTEM ID HERE
        type = "AWS"
    }
    identity_type = "USER"
    identity_ids = [
       // ARN  HERE
]

    filter_history_days = 90
}