package cloudknox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	//testProvider = Provider().(*schema.Provider)
}

func TestProvider(t *testing.T) {

}

func TestGoodCredentials(t *testing.T) {

}

func TestBadCredentials(t *tessting.T) {

}
