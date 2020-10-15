---
subcategory: "Privilege Management"
layout: "cloudknox"
page_title: "CloudKnox: cloudknox_role_policy"
description: |-
   The CloudKnox provider is used to interact with actions supported in the CloudKnox API
---

# Data Source: cloudknox_role_policy

Creates a `<name>.tf` file containing a Cloud Provider Terraform Resource with a Least-Privilege Role-Policy for AWS, GCP or Azure for the provided identities.

The `cloudknox_role_policy` data source will create different outputs based on the authorization system. Parameters can be set based on authorization system type to support functionality available through the [CloudKnox Portal](https://app.cloudknox.io)'s JEP Controller.

## Example Usage (AWS)

### Activity of Resources Example Declaration

The following block declares a `cloudknox_role_policy` named `resource-activity-aws-policy`. `identity_type` should be set to `RESOURCE` and all `identity_ids` should be set to resources ids such as the ids of the EC2 instances in this example. The policy is generated from the history of the activity of those resources between the millisecond timestamps specified. 

```hcl
data "cloudknox_role_policy" "resource-activity-aws-policy" {
    name = "resource-activity-aws-policy"
    output_path = "./resource_policies/"
    auth_system_info = {
        id = "123456789012"
        type = "AWS"
    }
    identity_type = "RESOURCE"
    identity_ids = [
        "arn:aws:ec2:us-east-1:123456789012:instance/i-0a1a2345b6cde7fg8",
    	"arn:aws:ec2:us-east-1:123456789012:instance/i-0a1a2345b6cde7fg9"]
    filter_history_start_time_millis = 123456789012
    filter_history_end_time_millis = 123456789012
}
```

### Activity of Resources Example Output

An `aws_iam_policy` resource is outputted to a file `./resource_policies/resource-activity-aws-policy.tf` containing the following AWS Terraform Provider Resources. If the policy exceeded 6144 characters, multiple `aws_iam_policy` policies would be generated in the same output file. Policies are named automatically according to the response from the CloudKnox API.

```hcl
resource "aws_iam_policy" "ck_activity_1234567890123_0" {
			name        = "ck_activity_1234567890123_0"
			path        = "/"
			description = "Cloudknox Generated IAM Role-Policy for AWS at 2020-07-16 12:21:21.7822427 -0700 PDT m=+0.381123001"
			policy = <<EOF
			{
		"Statement": [
			{
				"Action": [
					"route53:ListTagsForResource"
				],
				"Effect": "Allow",
				"Resource": [
					"*"
				],
				"Sid": "route53ReadActions"
			},
			
            // Truncated Policy Actions

		],
		"Version": "2012-10-17"
	}
EOF
}
```

### Activity of Roles Example Declaration

The following block declares a `cloudknox_role_policy` named `role-activity-aws-policy`. `identity_type` should be set to `ROLE` and `identity_ids` should be a list containing a role id. The policy is generated from the history of the activity of the roles from 90 days as set in `filter_history_days`. 

```hcl
data "cloudknox_role_policy" "role-activity-aws-policy" {
    name = "role-activity-aws-policy"
    output_path = "./"
    auth_system_info = {
        id = "123456789012"
        type = "AWS"
    }
    identity_type = "ROLE"
    identity_ids = [
    "arn:aws:iam::123456789012:role/IAM_R_KNOX_SECURITY"
  ]

    filter_history_days = 90
}
```

### Activity of Roles Example Output


An `aws_iam_policy` resource is outputted to a file `./role-activity-large-aws-policy.tf` containing the following resources. Since the AWS Policy exceeds 6144 characters, the Policy is automatically split across multiple resources denoted with the underscore in the policy name. If the Policy was less than 6144 characters, only a single resource will be created in the output file. Policies are named automatically according to the response from the CloudKnox API.

```hcl
resource "aws_iam_policy" "ck_activity_1234567890123_0" {
			name        = "ck_activity_1234567890123_0"
			path        = "/"
			description = "Cloudknox Generated IAM Role-Policy for AWS at 2020-07-16 16:41:10.6657102 -0700 PDT m=+0.846067101"
			policy = <<EOF
			{
		"Statement": [
            // Statements Truncated
		],
		"Version": "2012-10-17"
	}
EOF
}

// Resources Truncated

resource "aws_iam_policy" "ck_activity_1234567890123_2" {
			name        = "ck_activity_1234567890123_2"
			path        = "/"
			description = "Cloudknox Generated IAM Role-Policy for AWS at 2020-07-16 16:41:10.6657102 -0700 PDT m=+0.846067101"
			policy = <<EOF
			{
		"Statement": [
            // Statements Truncated
		],
		"Version": "2012-10-17"
	}
EOF
}
```

### Activity of Users Example Declaration

The following block declares a `cloudknox_role_policy` named `user-activity-aws-policy`. `identity_type` should be set to `USER` and all `identity_ids` should be user ids. The policy is generated from the history of the activity of those users from 90 days as set in `filter_history_days`. 

```hcl
data "cloudknox_role_policy" "user-activity-large-aws-policy" {
    name = "user-activity-large-aws-policy"
    output_path = "./"
    auth_system_info = {
        id = "123456789012"
        type = "AWS"
    }
    identity_type = "USER"
    identity_ids = [
    "arn:aws:iam::123456789012:user/bob",
    "arn:aws:iam::123456789012:user/carol",

    // Identities Truncated

    "arn:aws:iam::123456789012:user/judy-policy-boundary-test",
    "arn:aws:iam::123456789012:user/judy-policy-boundary-test-direct"
  ]

    filter_history_days = 90
}
```

### Activity of Users Example Output

```terraform
resource "aws_iam_policy" "ck_activity_1234567890123_0" {
			name        = "ck_activity_1234567890123_0"
			path        = "/"
			description = "Cloudknox Generated IAM Role-Policy for AWS at 2020-07-16 12:21:21.9465109 -0700 PDT m=+0.545391201"
			policy = <<EOF
			{
		"Statement": [
			// Statements Truncated
		],
		"Version": "2012-10-17"
	}
EOF
}

// Resources Truncated
```

## Example Usage (Azure)

### Activity of Apps Example Declaration

The following block declares a `cloudknox_role_policy` named `app-activity-azure-role`. `identity_type` should be set to `APP` and all `identity_ids` should be app ids. The policy is generated from the history of the activity of those apps from 90 days as set in `filter_history_days`. 

`filter_preserve_reads` is set to `true` meaning that any read permissions granted before are preserved. 

Azure requires that the parameter `request_params_scope` be set to the scope of permission.

```hcl
resource "cloudknox_role_policy" "app-activity-azure-role" {
    name = "app-activity-azure-role"
    output_path = "./"
    auth_system_info = {
         id = "12abcd34-56e7-890f-gh12-34i5678901jk",
         type = "AZURE"
     }
    identity_type = "APP"
    identity_ids = ["12abcd34-56e7-890f-gh12-34i5678901jk"]
    filter_history_days = 90
    filter_preserve_reads = true
    request_params_scope = "/subscriptions/12abcd34-56e7-890f-gh12-34i5678901jk"
}
```

### Activity of Apps Example Output

An `azurerm_role_definition` resource is outputted to a file `./app-activity-azure-role.tf` containing the following Terraform Resource. Policies are named automatically according to the response from the CloudKnox API.

```hcl
data "azurerm_role_definition" "app-activity-azure-role" {
			name        = "ck_activity_1234567890123"
			scope       = "/subscriptions/12abcd34-56e7-890f-gh12-34i5678901jk"
			description = "Cloudknox Generated IAM Role-Policy for AZURE at 2020-07-16 15:40:44.0841773 -0700 PDT m=+0.864027401"
		  
			permissions {
			  actions     = [
				"Microsoft.VMwareCloudSimple/*/read",
				"Microsoft.OffAzure/*/read",
                // Actions Truncated

			  ]
			  not_actions = [
			  ]
			}
		  
			assignable_scopes = [
				"/subscriptions/12abcd34-56e7-890f-gh12-34i5678901jk",

			]
		
}
```

### Activity of Users Example Declaration

The following block declares a `cloudknox_role_policy` named `user-activity-azure-role`. `identity_type` should be set to `USER` and all `identity_ids` should be set user ids. The policy is generated from the history of the activity of those users from 90 days as set in `filter_history_days`. 

```hcl
data "cloudknox_role_policy" "user-activity-azure-role" {
    name = "user-activity-azure-role"
    output_path = "./"
    auth_system_info = {
         id = "12abcd34-56e7-890f-gh12-34i5678901jk",
         type = "AZURE"
     }
    identity_type = "USER"
    identity_ids = ["alice@domain.io"]
    filter_history_days = 90
    filter_preserve_reads = true
    request_params_scope = "/subscriptions/12abcd34-56e7-890f-gh12-34i5678901jk"
}
```

### Activity of Users Example Output

An `azurerm_role_definition` resource is outputted to a file `./user-activity-azure-role.tf` containing the following Terraform Resource. Policies are named automatically according to the response from the CloudKnox API.

```hcl
resource "azurerm_role_definition" "user-activity-azure-role" {
			name        = "ck_activity_1234567890123"
			scope       = "/subscriptions/12abcd34-56e7-890f-gh12-34i5678901jk"
			description = "Cloudknox Generated IAM Role-Policy for AZURE at 2020-07-16 14:30:55.5363074 -0700 PDT m=+0.510089201"
		  
			permissions {
			  actions     = [
				"Microsoft.VMwareCloudSimple/*/read",
				"Microsoft.OffAzure/*/read",
				"Microsoft.Kubernetes/*/read",

                // Actions Truncated
			  ]
			  not_actions = [
			  ]
			}
		  
			assignable_scopes = [
				"/subscriptions/12abcd34-56e7-890f-gh12-34i5678901jk",
			]
		
}
```

## Example Usage (GCP)

### Activity of Users Example Declaration

The following block declares a `cloudknox_role_policy` named `user-activity-gcp-role`. `identity_type` should be set to `USER` and all `identity_ids` should be set to a user. The policy is generated from the history of the activity of thoose users from 90 days as set in `filter_history_days`. 

```hcl
data "cloudknox_role_policy" "user-activity-gcp-role" {
    name = "user-activity-gcp-role"
    output_path = "./"
    auth_system_info = {
         id = "silicon-banana-123456",
         type = "GCP"
     }
    identity_type = "USER"
    identity_ids = ["grace@domain.io"]
    filter_history_days = 90
}
```

### Activity of Users Example Output

A `google_project_iam_custom_role` resource is outputted to a file `./user-activity-gcp-role.tf` containing the following Terraform Resource. Policies are named automatically according to the response from the CloudKnox API.

```hcl
resource "google_project_iam_custom_role" "user-activity-gcp-role" {
		role_id     = "ck_activity_1234567890123"
		title		= "user-activity-gcp-role"
		description = "Cloudknox Generated IAM Role-Policy for GCP at 2020-07-16 14:30:55.374293 -0700 PDT m=+0.348074801"
		permissions = [
			"storage.buckets.list",
			"compute.disks.list",

            // Permissions Truncated
		]
}
```





## Argument Reference

The following arguments are supported:

### Required
* `name` - (Required) Name of the policy, can match the terraform data source name.
* `output_path` - (Required) Directory where the terraform script will be outputted.
* `auth_system_info` - (Required) Set to the following map.

```
{
    id = Enter the ID as a string. AWS Account ID, GCP Project ID, or Azure Subcription ID. 
    type = Enter AWS, GCP or AZURE as a string
}
```

### AWS

* `identity_type` - (Required) Identity type of the ids. Enter `RESOURCE`, `ROLE` or `USER`.
* `identity_ids` - (Required) Provide a comma seperated list of ARNs containing of type `identity_type`.
* `filter_history_days` - (Optional) Number of days in the past to look at the actions of `identity_ids` to generate a policy.
* `filter_history_start_time_millis` - (Optional) Start time in unix time milliseconds to look at actions of `identity_ids`.
* `filter_history_end_time_millis` - (Optional) End time in unix time milliseconds to look at actions of `identity_ids`.

~> **Note** Although these parameters are optional, you must use `filter_history_days` xor (`filter_history_start_time_millis` and `filter_history_end_time_millis`) as only one parameter will be considered when generating an AWS policy. 

### Azure

* `identity_type` - (Required) Identity type of the ids. Enter `APP` or `USER`.
* `identity_ids` - (Required) Provide a comma seperated list of strings containing Azure `ids` of type `identity_type`.
* `filter_history_days` - (Optional) Number of days in the past to look at the actions of `identity_ids` to generate a policy.
* `filter_preserve_reads` - (Optional) Boolean to indicate preserve read permissions that the identity already has.
* `filter_history_start_time_millis` - (Optional) Start time in unix time milliseconds to look at actions of `identity_ids`.
* `filter_history_end_time_millis` - (Optional) End time in unix time milliseconds to look at actions of `identity_ids`.
* `request_params_scope` - (Required) Enter the Azure scope =.

~> **Note** Although these parameters are optional, you must use `filter_history_days` xor (`filter_history_start_time_millis` and `filter_history_end_time_millis`) as only one parameter will be considered when generating a Azure role. 

### GCP

* `identity_type` - (Required) Identity type of the ids. Enter `USER`
* `identity_ids` - (Required) Provide a comma seperated list of strings containing GCP `ids` of type `identity_type`
* `filter_history_days` - (Optional) Number of days in the past to look at the actions of `identity_ids` to generate a policy
* `filter_history_start_time_millis` - (Optional) Start time in unix time milliseconds to look at actions of `identity_ids`
* `filter_history_end_time_millis` - (Optional) End time in unix time milliseconds to look at actions of `identity_ids`

~> **Note** Although these parameters are optional, you must use `filter_history_days` xor (`filter_history_start_time_millis` and `filter_history_end_time_millis`) as only one parameter will be considered when generating a GCP role. 

### Optional
* `request_params_resource` - Optional parameter for Cloudknox API
* `request_params_resources` - Optional list of parameters for Cloudknox API
* `request_params_condition` - Optional parameter for Cloudknox API




