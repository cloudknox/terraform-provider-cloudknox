package cloudknox

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccPolicy_Basic(t *testing.T) {
	//resourceName := "cloudknox_policy.test_policy"
	resource.Test(t, resource.TestCase{
		Providers: TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccPolicyConfig(),
				// Check: resource.ComposeTestCheckFunc(
				// 	resource.TestCheckResourceAttr(
				// 		resourceName, "name", "test_policy"),
				// ),
			},
		},
	})
}

// configs
func testAccPolicyConfig() string {
	return `
resource "cloudknox_policy" "test_policy" {
	name = "test_policy"
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
