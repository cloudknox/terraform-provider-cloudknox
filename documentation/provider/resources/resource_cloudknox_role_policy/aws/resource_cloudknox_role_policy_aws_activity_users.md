# CloudKnox Role-Policy Resource (AWS Usage)

## Create a Policy based on Activity of User(s)

An AWS IAM Policy is created based on the Activity of User(s) provided

### Example

#### Terraform Resource

The following block declares a `cloudknox_role_policy` named `user-activity-aws-policy`. `identity_type` should be set to "USER" and all `identity_ids` should be set to a user. The policy is generated from the history of the activity of thoose users from 90 days as set in `filter_history_days`. 

```terraform
resource "cloudknox_role_policy" "user-activity-large-aws-policy" {
    name = "user-activity-large-aws-policy"
    output_path = "./"
    auth_system_info = {
        id = "377596131774"
        type = "AWS"
    }
    identity_type = "USER"
    identity_ids = [
    "arn:aws:iam::377596131774:user/balaji",
    "arn:aws:iam::377596131774:user/carl",
    "arn:aws:iam::377596131774:user/devops",
    "geeta+okta@cloudknox.io",
    "arn:aws:iam::377596131774:user/geeta.alapati",
    "arn:aws:iam::377596131774:user/geetatest2",
    "arn:aws:iam::377596131774:user/geetauser3",
    "guru+okta@cloudknox.io",
    "guru+okta01@cloudknox.io",
    "arn:aws:iam::377596131774:user/guru.ramprakash",
    "arn:aws:iam::377596131774:user/josh",
    "josh+okta@cloudknox.io",
    "arn:aws:iam::377596131774:user/josh-policy-boundary-test",
    "arn:aws:iam::377596131774:user/josh-policy-boundary-test-direct"
  ]

    filter_history_days = 90
}
```

#### Output

An `aws_iam_policy` resource is outputted to a file `./user-activity-large-aws-policy.tf` containing the following resources. Since the AWS Policy exceeds 6144 characters, the Policy is automatically split across multiple resources denoted with the underscore in the policy name. If the Policy was less than 6144 characters, only a single resource will be created in the output file. Policies are named automatically according to the response from the CloudKnox API.

```terraform
resource "aws_iam_policy" "ck_activity_1594927282578_0" {
			name        = "ck_activity_1594927282578_0"
			path        = "/"
			description = "Cloudknox Generated IAM Role-Policy for AWS at 2020-07-16 12:21:21.9465109 -0700 PDT m=+0.545391201"
			policy = <<EOF
			{
		"Statement": [
			{
				"Action": [
					"cloudformation:DeleteStack"
				],
				"Effect": "Allow",
				"Resource": [
					"*"
				],
				"Sid": "cloudformationDeleteActions"
			},

	        // Truncated Policy Actions

			{
				"Action": [
					"ec2:DeleteSecurityGroup",
					"ec2:TerminateInstances"
				],
				"Effect": "Allow",
				"Resource": [
					"*"
				],
				"Sid": "ec2DeleteActions"
			}
		],
		"Version": "2012-10-17"
	}
EOF
}

resource "aws_iam_policy" "ck_activity_1594927282578_1" {
			name        = "ck_activity_1594927282578_1"
			path        = "/"
			description = "Cloudknox Generated IAM Role-Policy for AWS at 2020-07-16 12:21:21.9465109 -0700 PDT m=+0.545391201"
			policy = <<EOF
			{
		"Statement": [
			{
				"Action": [
					"redshift:DescribeClusters",

                    // Truncated Policy Action Items

					"redshift:DescribeHsmClientCertificates"
				],
				"Effect": "Allow",
				"Resource": [
					"*"
				],
				"Sid": "redshiftReadActions"
			},

            // Truncated Policy Actions

			{
				"Action": [
					"s3:PutBucketPublicAccessBlock",
					"s3:PutBucketAcl",
					"s3:PutBucketPolicy"
				],
				"Effect": "Allow",
				"Resource": [
					"*"
				],
				"Sid": "s3PermissionsActions"
			}
		],
		"Version": "2012-10-17"
	}
EOF
}
```

