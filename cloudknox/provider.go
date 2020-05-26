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
			"service_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["service_account_id"],
			},

			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["access_key"],
			},

			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions["secret_key"],
			},

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
		ServiceAccountID:      d.Get("service_account_id").(string),
		AccessKey:             d.Get("access_key").(string),
		SecretKey:             d.Get("secret_key").(string),
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
		"service_account_id": "Cloudknox Service Account ID",

		"access_key": "Access key that you would like the service account to use",

		"secret_key": "Associated secret key for the access key",

		"shared_credentials_file": "Filename of the HOCON credentials file including path.",

		"profile": "Profile for accessKey/secretKey pair you would like to use.",
	}
}
