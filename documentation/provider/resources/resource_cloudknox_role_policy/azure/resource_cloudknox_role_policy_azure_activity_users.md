# CloudKnox Role-Policy Resource (Azure Usage)

## Create a Role based on Activity of User(s)

An Azure Role Definition is created based on the Activity of User(s) provided

### Example

#### Terraform Resource

The following block declares a `cloudknox_role_policy` named `user-activity-azure-role`. `identity_type` should be set to `USER` and all `identity_ids` should be set user ids. The policy is generated from the history of the activity of those users from 90 days as set in `filter_history_days`. 

`filter_preserve_reads` is set to `true` meaning that any read permissions granted before are preserved. 

Azure requires that the parameter `request_params_scope` be set to the scope of permission.

```terraform
resource "cloudknox_role_policy" "user-activity-azure-role" {
    name = "user-activity-azure-role"
    output_path = "./"
    auth_system_info = {
         id = "12abcd90-95a3-123a-ab12-56f1234565ee",
         type = "AZURE"
     }
    identity_type = "USER"
    identity_ids = ["alice@cloudknoxsecurity.io"]
    filter_history_days = 90
    filter_preserve_reads = true
    request_params_scope = "/subscriptions/12abcd90-95a3-123a-ab12-56f1234565ee"
}
```

#### Output

An `azurerm_role_definition` resource is outputted to a file `./user-activity-azure-role.tf` containing the following Terraform Resource. Policies are named automatically according to the response from the CloudKnox API.

```terraform
resource "azurerm_role_definition" "user-activity-azure-role" {
			name        = "ck_activity_1594935056278"
			scope       = "/subscriptions/12abcd90-95a3-123a-ab12-56f1234565ee"
			description = "Cloudknox Generated IAM Role-Policy for AZURE at 2020-07-16 14:30:55.5363074 -0700 PDT m=+0.510089201"
		  
			permissions {
			  actions     = [
				"Microsoft.VMwareCloudSimple/*/read",
				"Microsoft.OffAzure/*/read",
				"Microsoft.Kubernetes/*/read",

                // Actions Truncated
		
				"Microsoft.ChangeAnalysis/*/read",
				"paraleap.cloudmonix/*/read",
				"Microsoft.Compute/*/read",

			  ]
			  not_actions = [
			  ]
			}
		  
			assignable_scopes = [
				"/subscriptions/12abcd90-95a3-123a-ab12-56f1234565ee",

			]
		
}
```

