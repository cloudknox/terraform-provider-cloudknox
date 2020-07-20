# CloudKnox Role-Policy Data Source (AWS Usage)

## Create a Policy based on Activity of Resource(s)

An AWS IAM Policy is created based on the Activity of Resource(s) provided

### Example

#### Terraform Data Source 

The following block declares a `cloudknox_role_policy` named `resource-activity-aws-policy`. `identity_type` should be set to `RESOURCE` and all `identity_ids` should be set to resources ids such as the ids of the EC2 instances in this example. The policy is generated from the history of the activity of those resources between the millisecond timestamps specified. 

```terraform
data "cloudknox_role_policy" "resource-activity-aws-policy" {
    name = "resource-activity-aws-policy"
    output_path = "./resource_policies/"
    auth_system_info = {
        id = "123456789012"
        type = "AWS"
    }
    identity_type = "RESOURCE"
    identity_ids = [
        "arn:aws:ec2:us-east-1:123456789012:instance/i-0a5a0012f0asd7de0",
    	"arn:aws:ec2:us-east-1:123456789012:instance/i-03asda213hjkj3329"]
    filter_history_start_time_millis = 123456789012
    filter_history_end_time_millis = 123456789012
}
```

#### Output

An `aws_iam_policy` resource is outputted to a file `./resource_policies/resource-activity-aws-policy.tf` containing the following AWS Terraform Provider Resources. If the policy exceeded 6144 characters, multiple `aws_iam_policy` policies would be generated in the same output file. Policies are named automatically according to the response from the CloudKnox API.

```terraform
resource "aws_iam_policy" "ck_activity_1593123412453_0" {
			name        = "ck_activity_1593123412453_0"
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

			{
				"Action": [
					"rds:ListTagsForResource"
				],
				"Effect": "Allow",
				"Resource": [
					"*"
				],
				"Sid": "rdsReadActions"
			}
		],
		"Version": "2012-10-17"
	}
EOF
}


```

