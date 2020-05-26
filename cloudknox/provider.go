package cloudknox

import (
	"cloudknox/terraform-provider-cloudknox/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider creates and returns a Terraform Provider with populated Schema
func Provider() terraform.ResourceProvider {
	log := common.GetLogger()
	log.Info("Building Provider")
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{

			"shared_credentials_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["shared_credentials_file"],
			},
			"profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["profile"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudknox_policy": resourcePolicy(),
		},

		ConfigureFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log := common.GetLogger()
	log.Info("Configuring Provider")

	parameters := &common.ClientParameters{
		SharedCredentialsFile: d.Get("shared_credentials_file").(string),
		Profile:               d.Get("profile").(string),
	}

	common.SetConfiguration(parameters) //Build Client Struct using parameters
	return common.GetClient()           //Return the Client Struct and the Error
}

var descriptions map[string]string

func init() {

	log := common.GetLogger()

	log.Info("Building Descriptions")
	descriptions = map[string]string{

		"shared_credentials_file": "Filename of the HOCON credentials file including path.",

		"profile": "Profile for accessKey/secretKey pair you would like to use.",
	}
}
