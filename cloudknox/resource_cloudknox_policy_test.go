package cloudknox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var resources = [...]string{"cloudknox_policy.test_aws_policy",
	"cloudknox_policy.test_gcp_policy",
	"cloudknox_policy.test_azure_policy"}

func TestAccPolicy_Basic(t *testing.T) {

	// Test AWS
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyConfigAWS(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resources[0], "name", "test_aws_policy",
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
				Config: testAccPolicyConfigGCP(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resources[1], "name", "test_gcp_policy",
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
				Config: testAccPolicyConfigAZURE(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						resources[2], "name", "test_azure_policy",
					),
				),
			},
		},
	})

}

// configs
func testAccPolicyConfigAWS() string {
	return `
resource "cloudknox_policy" "test_aws_policy" {
	name = "test_aws_policy"
	output_path = "./"
	auth_system_info = {
		id = "377596131774"
		type = "AWS"
	}
	identity_type = "RESOURCE"
	identity_ids = [
		"arn:aws:ec2:us-east-1:377596131774:instance/i-0a5e0048fb0237de0",
		"arn:aws:ec2:us-east-1:377596131774:instance/i-03689effa30f70329"]
	filter_history_start_time_millis = 1585071573512
	filter_history_end_time_millis = 1592847573512
}
`
}

func testAccPolicyConfigGCP() string {
	return `
resource "cloudknox_policy" "test_gcp_policy" {
	name = "test_gcp_policy"
	output_path = "./"
	auth_system_info = {
			id = "carbide-bonsai-205017",
			type = "GCP"
		}
	identity_type = "USER"
	identity_ids = ["geeta@cloudknox.io"]
	filter_history_days = 90
	filter_preserve_reads = true
}
`
}

func testAccPolicyConfigAZURE() string {
	return `
resource "cloudknox_policy" "test_azure_policy" {
	name = "test_azure_policy"
	output_path = "./"
	auth_system_info = {
			id = "87eefd90-95a3-480a-ba42-56ff299a05ee",
			type = "AZURE"
		}
	identity_type = "USER"
	identity_ids = ["aislam@cloudknoxsecurity.io"]
	filter_history_days = 90
	filter_preserve_reads = true
	request_params_scope = "/subscriptions/87eefd90-95a3-480a-ba42-56ff299a05ee"
}
`
}
