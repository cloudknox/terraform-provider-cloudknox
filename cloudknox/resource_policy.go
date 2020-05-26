package cloudknox

import (
	"cloudknox/terraform-provider-cloudknox/common"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourcePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyCreate,
		Read:   resourcePolicyRead,
		Update: resourcePolicyUpdate,
		Delete: resourcePolicyDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePolicyCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*common.Client)
	err := common.ValidateClient(client)
	log := common.GetLogger()

	if err != nil {
		log.Info(err)
		return err
	}
	log.Info("Creating New Policy")
	log.Info("Dummy resource property test " + d.Get("address").(string))

	return nil
}

func resourcePolicyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourcePolicyDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
