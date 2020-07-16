# CloudKnox Role-Policy Resource (GCP Usage)

## Create a Role based on Activity of User(s)

A GCP Role Definition is created based on the Activity of User(s) provided

### Example

#### Terraform Resource

The following block declares a `cloudknox_role_policy` named `user-activity-gcp-role`. `identity_type` should be set to "USER" and all `identity_ids` should be set to a user. The policy is generated from the history of the activity of thoose users from 90 days as set in `filter_history_days`. 

```terraform
resource "cloudknox_role_policy" "user-activity-gcp-role" {
    name = "user-activity-gcp-role"
    output_path = "./"
    auth_system_info = {
         id = "carbide-bonsai-205017",
         type = "GCP"
     }
    identity_type = "USER"
    identity_ids = ["geeta@cloudknox.io"]
    filter_history_days = 90
}
```

#### Output

An `google_project_iam_custom_role` resource is outputted to a file `./user-activity-gcp-role.tf` containing the following Terraform Resource. Policies are named automatically according to the response from the CloudKnox API.

```terraform
resource "google_project_iam_custom_role" "user-activity-gcp-role" {
		role_id     = "ck_activity_1594935056222"
		title		= "user-activity-gcp-role"
		description = "Cloudknox Generated IAM Role-Policy for GCP at 2020-07-16 14:30:55.374293 -0700 PDT m=+0.348074801"
		permissions = [
			"storage.buckets.list",
			"compute.disks.list",

            // Permissions Truncated

			"storage.buckets.setIamPolicy",
			"compute.diskTypes.list",

		]
}
```

