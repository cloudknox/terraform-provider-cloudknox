# CloudKnox Role-Policy Resource (Azure Usage)

## Create a Role based on Activity of App(s)

An Azure Role Definition is created based on the Activity of App(s) provided

### Example

#### Terraform Resource

The following block declares a `cloudknox_role_policy` named `app-activity-azure-role`. `identity_type` should be set to `APP` and all `identity_ids` should be app ids. The policy is generated from the history of the activity of those apps from 90 days as set in `filter_history_days`. 

`filter_preserve_reads` is set to `true` meaning that any read permissions granted before are preserved. 

Azure requires that the parameter `request_params_scope` be set to the scope of permission.

```terraform
resource "cloudknox_role_policy" "app-activity-azure-role" {
    name = "app-activity-azure-role"
    output_path = "./"
    auth_system_info = {
         id = "12abcd90-95a3-123a-ab12-56f1234565ee",
         type = "AZURE"
     }
    identity_type = "APP"
    identity_ids = ["123ab012-f123-123b-a143-23184e64d6a0"]
    filter_history_days = 90
    filter_preserve_reads = true
    request_params_scope = "/subscriptions/12abcd90-95a3-123a-ab12-56f1234565ee"
}
```

#### Output

An `azurerm_role_definition` resource is outputted to a file `./app-activity-azure-role.tf` containing the following Terraform Resource. Policies are named automatically according to the response from the CloudKnox API.

```terraform
resource "azurerm_role_definition" "app-activity-azure-role" {
			name        = "ck_activity_1594939244616"
			scope       = "/subscriptions/12abcd90-95a3-123a-ab12-56f1234565ee"
			description = "Cloudknox Generated IAM Role-Policy for AZURE at 2020-07-16 15:40:44.0841773 -0700 PDT m=+0.864027401"
		  
			permissions {
			  actions     = [
				"Microsoft.VMwareCloudSimple/*/read",
				"Microsoft.OffAzure/*/read",

                // Actions Truncated

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

