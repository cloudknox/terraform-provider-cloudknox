package cloudknox

import (
	"cloudknox/terraform-provider-cloudknox/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider creates and returns a Terraform Provider with populated Schema
func Provider() terraform.ResourceProvider {
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
	log.Info("Setting up client parameters")
	parameters := &common.ClientParameters{
		ServiceAccountId:      d.Get("service_account_id").(string),
		AccessKey:             d.Get("access_key").(string),
		SecretKey:             d.Get("secret_key").(string),
		SharedCredentialsFile: d.Get("shared_credentials_file").(string),
		Profile:               d.Get("profile").(string),
	}

	common.SetConfiguration(parameters)

	return common.GetClient()
}

var descriptions map[string]string

func init() {

	descriptions = map[string]string{
		"service_account_id": "Cloudknox Service Account ID",

		"access_key": "Access key that you would like the service account to use",

		"secret_key": "Associated secret key for the access key",

		"shared_credentials_file": "Filename of the HOCON credentials file including path.",

		"profile": "Profile for accessKey/secretKey pair you would like to use.",
	}

	Log := common.GetLogger()
	Log.Info("we here")

	//Set the configuration for the provider based on given paramaters.

	//figure out how to read the terraform provider properties then pass this map into configurator
	//also figure out the type of provider to use based on given stuff
}
