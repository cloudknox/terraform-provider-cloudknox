package cloudknox

import (
	"terraform-provider-cloudknox/cloudknox/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mitchellh/go-homedir"
)

const (
	resourcePath = "/opt/cloudknox/terraform-provider-cloudknox-config.yaml"
)

// Provider creates and returns a Terraform Provider with populated Schema
func Provider() terraform.ResourceProvider {
	logger := common.GetLogger()
	logger.Info("msg", "Building Cloudknox Terraform Provider")
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
			common.NewPolicy: resourcePolicy(),
		},

		ConfigureFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	logger := common.GetLogger()
	logger.Info("msg", "Configuring Cloudknox Terraform Provider")

	parameters := &common.ClientParameters{
		SharedCredentialsFile: d.Get("shared_credentials_file").(string),
		Profile:               d.Get("profile").(string),
	}

	home, _ := homedir.Dir()

	err := common.SetConfiguration(home + resourcePath)

	if err != nil {
		return nil, err
	}

	common.SetClientConfiguration(parameters) //Build Client Struct using parameters
	return common.GetClient()                 //Return the Client Struct and the Error
}

var descriptions map[string]string

func init() {

	logger := common.GetLogger()
	logger.Debug("msg", "Running Initialization Function")
	descriptions = map[string]string{

		"shared_credentials_file": "Path/Filename of the HOCON credentials file.",

		"profile": "Profile for (SERVICE_ACCOUNT_ID, ACCESS_KEY, SECRET_KEY) triplet you would like to use in a HOCON credentials file.",
	}
}
