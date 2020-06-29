package cloudknox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var TestAccProviders map[string]terraform.ResourceProvider
var TestAccProvider terraform.ResourceProvider

func init() {
	TestAccProvider = Provider().(terraform.ResourceProvider)
	TestAccProviders = map[string]terraform.ResourceProvider{
		"cloudknox": TestAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ = Provider()
}
