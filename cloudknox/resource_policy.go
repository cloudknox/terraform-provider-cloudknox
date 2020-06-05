package cloudknox

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/apiHandler"
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"

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
				Default:  nil,
			},
			"filter_preserve_reads": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  nil,
			},
			"filter_history_start_time_millis": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  nil,
			},
			"filter_history_end_time_millis": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  nil,
			},
			"request_params_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"request_params_resource": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
			"request_params_resources": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Default:  nil,
			},
			"request_params_condition": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
		},
	}
}

func resourcePolicyCreate(d *schema.ResourceData, m interface{}) error {
	logger := common.GetLogger()
	logger.Info("msg", "Building Policy Payload")

	var payload apiHandler.PolicyData

	logger.Info("msg", "Reading Resource Data")

	name := d.Get("name").(string)

	payload.AuthSystemInfo.ID = d.Get("auth_system_info").(map[string]interface{})["id"].(string)
	payload.AuthSystemInfo.Type = d.Get("auth_system_info").(map[string]interface{})["type"].(string)
	payload.IdentityType = d.Get("identity_type").(string)
	payload.IdentityIds = d.Get("identity_ids")

	var days = d.Get("filter_history_days").(int)
	var start int = d.Get("filter_history_start_time_millis").(int)
	var end int = d.Get("filter_history_end_time_millis").(int)

	if days != 0 {
		logger.Debug("msg", "Filter History Days Given", "days", days)
		payload.Filter.HistoryDays = days
	}

	if start != 0 && end != 0 {
		logger.Debug("msg", "Filter History Bounds Given")
		payload.Filter.HistoryDuration = &apiHandler.HD{
			StartTime: start,
			EndTime:   end,
		}
	}

	payload.Filter.PreserveReads = d.Get("filter_preserve_reads").(bool)

	var scope interface{} = d.Get("request_params_scope")
	var resource interface{} = d.Get("request_params_resource")
	var resources interface{} = d.Get("request_params_resources")
	var condition interface{} = d.Get("request_params_condition")

	logger.Debug("scope", scope.(string), "resource", resource.(string), "resources", resources, "condition", condition.(string))

	if scope == "" && resource == "" && resources == nil && condition == "" {
		logger.Debug("msg", "No Request Params Given")
	} else {
		logger.Debug("msg", "Request Params Given")

		var requestParams apiHandler.RP

		if scope.(string) == "" {
			requestParams.Scope = nil
		} else {
			requestParams.Scope = scope.(string)
		}

		if resource.(string) == "" {
			requestParams.Resource = nil
		} else {
			requestParams.Resource = resource.(string)
		}

		requestParams.Resources = resources

		if condition.(string) == "" {
			requestParams.Condition = nil
		} else {
			requestParams.Condition = condition.(string)
		}

		payload.RequestParams = &requestParams
	}

	logger.Info("msg", "Payload Successfully Built")
	err := apiHandler.NewPolicy(payload.AuthSystemInfo.Type, name, d.Get("output_path").(string), &payload)

	if err != nil {
		return err
	}

	d.SetId(name)

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
