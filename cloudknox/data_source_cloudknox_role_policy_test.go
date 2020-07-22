package cloudknox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var resources = [...]string{
	"cloudknox_role_policy.test_aws_policy",
	"cloudknox_role_policy.test_gcp_policy",
	"cloudknox_role_policy.test_azure_policy",
}

func TestAccRolePolicy_Basic(t *testing.T) {

	// Test AWS
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRolePolicyConfigAWS(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resources[0], "name", "resource-activity-aws-policy",
					),
				),
			},
		},
	})

	// Test GCP
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRolePolicyConfigGCP(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resources[1], "name", "resource-activity-aws-policy",
					),
				),
			},
		},
	})

	// Test AZURE
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccRolePolicyConfigAZURE(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resources[2], "name", "user-activity-azure-role",
					),
				),
			},
		},
	})

}

// configs
func testAccRolePolicyConfigAWS() string {
	return `
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
`
}

func testAccRolePolicyConfigGCP() string {
	return `
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
`
}

func testAccRolePolicyConfigAZURE() string {
	return `
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
`
}
