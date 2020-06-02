package cloudknox

import (
	"cloudknox/terraform-provider-cloudknox/apiHandler"
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

			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"output_path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"auth_system_info": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					//id -> string
					//resource -> string ie AWS GCP ETC
				},
				Required: true,
			},
			"identity_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identity_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"filter_history_days": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"filter_preserve_reads": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"filter_history_start_time_millis": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"filter_history_end_time_millis": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"request_params_scope": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_params_resource": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_params_resources": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"request_params_condition": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePolicyCreate(d *schema.ResourceData, m interface{}) error {
	logger := common.GetLogger()
	logger.Info("msg", "Building Policy Payload")

	var payload apiHandler.PolicyData

	logger.Info("msg", "Reading Resource Data")
	payload.AuthSystemInfo.ID = d.Get("auth_system_info").(map[string]interface{})["id"].(string)
	payload.AuthSystemInfo.Type = d.Get("auth_system_info").(map[string]interface{})["type"].(string)
	payload.IdentityType = d.Get("identity_type").(string)
	payload.IdentityIds = d.Get("identity_ids")
	payload.Filter.HistoryDays = d.Get("filter_history_days").(int)
	payload.Filter.PreserveReads = d.Get("filter_preserve_reads").(bool)
	payload.Filter.HistoryDuration.StartTime = d.Get("filter_history_start_time_millis").(int)
	payload.Filter.HistoryDuration.EndTime = d.Get("filter_history_end_time_millis").(int)

	payload.RequestParams.Scope = d.Get("request_params_scope").(string)
	payload.RequestParams.Resource = d.Get("request_params_resource").(string)
	payload.RequestParams.Resources = d.Get("request_params_resources")
	payload.RequestParams.Condition = d.Get("request_params_condition").(string)

	logger.Info("msg", "Payload Successfully Built")
	err := apiHandler.NewPolicy(d.Get("name").(string), d.Get("output_path").(string), &payload)

	if err != nil {
		return err
	}

	d.SetId(d.Get("name").(string))

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
