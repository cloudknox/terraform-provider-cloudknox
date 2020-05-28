package cloudknox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func init() {
	testProvider = Provider().(*schema.Provider)
}

func TestProvider(t *testing.T) {

}

func TestGoodCredentials(t *testing.T) {

}

func TestBadCredentials(t *testing.T) {

}
