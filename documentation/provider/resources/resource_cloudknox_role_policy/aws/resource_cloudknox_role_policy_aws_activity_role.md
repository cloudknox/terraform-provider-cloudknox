# CloudKnox Role-Policy Resource (AWS Usage)

## Create a Policy based on Activity of a Role

An AWS IAM Policy is created based on the Activity of a Role provided

### Example

#### Terraform Resource

The following block declares a `cloudknox_role_policy` named `role-activity-aws-policy`. `identity_type` should be set to `ROLE` and `identity_ids` should be a list containing a role id. The policy is generated from the history of the activity of the roles from 90 days as set in `filter_history_days`. 

```terraform
resource "cloudknox_role_policy" "role-activity-aws-policy" {
    name = "role-activity-aws-policy"
    output_path = "./"
    auth_system_info = {
        id = "377596131774"
        type = "AWS"
    }
    identity_type = "ROLE"
    identity_ids = [
    "arn:aws:iam::377596131774:role/IAM_R_KNOX_SECURITY"
  ]

    filter_history_days = 90
}
```

#### Output

An `aws_iam_policy` resource is outputted to a file `./role-activity-large-aws-policy.tf` containing the following resources. Since the AWS Policy exceeds 6144 characters, the Policy is automatically split across multiple resources denoted with the underscore in the policy name. If the Policy was less than 6144 characters, only a single resource will be created in the output file. Policies are named automatically according to the response from the CloudKnox API.

```terraform
resource "aws_iam_policy" "ck_activity_1594942871390_0" {
			name        = "ck_activity_1594942871390_0"
			path        = "/"
			description = "Cloudknox Generated IAM Role-Policy for AWS at 2020-07-16 16:41:10.6657102 -0700 PDT m=+0.846067101"
			policy = <<EOF
			{
		"Statement": [
			{
				"Action": [
					"eks:ListClusters"
				],
				"Effect": "Allow",
				"Resource": [
					"*"
				],
				"Sid": "eksReadActions"
			},
			
            // Statements Truncated

		],
		"Version": "2012-10-17"
	}
EOF
}

// Resources Truncated

resource "aws_iam_policy" "ck_activity_1594942871390_2" {
			name        = "ck_activity_1594942871390_2"
			path        = "/"
			description = "Cloudknox Generated IAM Role-Policy for AWS at 2020-07-16 16:41:10.6657102 -0700 PDT m=+0.846067101"
			policy = <<EOF
			{
		"Statement": [
			{
				"Action": [
					"s3:GetBucketLocation",
    
                    // Actions Truncated

					"s3:GetAccountPublicAccessBlock"
				],
				"Effect": "Allow",
				"Resource": [
					"*"
				],
				"Sid": "s3ReadActions"
			},
			
            // Statements Truncated
		],
		"Version": "2012-10-17"
	}
EOF
}


```

