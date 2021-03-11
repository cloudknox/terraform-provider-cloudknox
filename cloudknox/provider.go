package cloudknox

import (
	"github.com/cloudknox/terraform-provider-cloudknox/cloudknox/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider creates and returns a Terraform Provider with populated Schema
func Provider() terraform.ResourceProvider {
	logger := common.GetLogger()
	logger.Debug("msg", "initializing Cloudknox terraform provider")

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
		DataSourcesMap: map[string]*schema.Resource{
			common.RolePolicy: dataSourceRolePolicy(),
		},
		ConfigureFunc: providerConfigure,
	}

	return provider
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	logger := common.GetLogger()
	logger.Info("msg", "setting Cloudknox terraform provider parameters")
	parameters := &common.ClientParameters{
		SharedCredentialsFile: d.Get("shared_credentials_file").(string),
		Profile:               d.Get("profile").(string),
	}
	credentials := common.GetCredentials(parameters) //Build Client Struct using parameters
	return common.NewClient(credentials)             //Return the Client Struct and the Error
}

var descriptions map[string]string

func init() {
	logger := common.GetLogger()
	logger.Debug("msg", "running initialization function")
	descriptions = map[string]string{
		"shared_credentials_file": "Path/Filename of the credentials file.",
		"profile":                 "Profile for (SERVICE_ACCOUNT_ID, ACCESS_KEY, SECRET_KEY) triplet you would like to use in a credentials file.",
	}
}
