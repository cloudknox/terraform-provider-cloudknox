package cloudknox

import (
	"encoding/json"
	"fmt"
	"terraform-provider-cloudknox/cloudknox/api/helpers"
	"terraform-provider-cloudknox/cloudknox/api/models"
	"terraform-provider-cloudknox/cloudknox/common"
	"terraform-provider-cloudknox/cloudknox/utils"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceRolePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePolicyRead,

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
			},
			"request_params_condition": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  nil,
			},
		},
	}
}

func dataSourcePolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.Client)
	logger := common.GetLogger()

	logger.Info("msg", "creating role_policy data source")

	name := d.Get("name").(string)
	outputPath := d.Get("output_path").(string)
	logger.Info("msg", "Building Policy Payload")
	payload := getRolePolicyPayload(d)
	logger.Debug("msg", "payload successfully built", "role_policy", name)
	logger.Info("msg", "creating new role-policy", "name", name+".tf", "output_path", outputPath)

	payloadBytes, _ := json.Marshal(payload)
	response, err := client.POST("api/v2/role-policy/new", payloadBytes)
	if err != nil {
		logger.Error("msg", "unable to complete POST request", "error", err.Error())
		return err
	}
	rolePolicyDataBytes, err := json.Marshal(response["data"])
	if err != nil {
		logger.Error("msg", "JSON marshaling error while preparing data", "json_error", err)
	}
	rolePolicyDataString := string(rolePolicyDataBytes)
	logger.Debug("rolePolicyJSONString", utils.Truncate(rolePolicyDataString, 30, true))
	args := map[string]string{
		"name":        name,
		"description": "Cloudknox Generated IAM Role-Policy for " + payload.AuthSystemInfo.Type + " at " + time.Now().String(),
		"output_path": outputPath,
		"aws_path":    "/",
		"data":        rolePolicyDataString,
	}

	logger.Debug("msg", "Begin Write Sequence")
	err = helpers.WriteResource("cloudknox_role_policy", payload.AuthSystemInfo.Type, args)
	if err != nil {
		logger.Error("msg", "unable to write role_policy", "write_error", err.Error())
		return err
	}
	logger.Debug("msg", "write sequence completed successfully")
	d.SetId(name)
	return nil
}

func getRolePolicyPayload(d *schema.ResourceData) models.RolePolicyData {
	logger := common.GetLogger()
	var payload models.RolePolicyData

	logger.Debug("msg", "Reading Resource Data")
	payload.AuthSystemInfo.ID = d.Get("auth_system_info").(map[string]interface{})["id"].(string)
	payload.AuthSystemInfo.Type = d.Get("auth_system_info").(map[string]interface{})["type"].(string)
	payload.IdentityType = d.Get("identity_type").(string)
	payload.IdentityIds = d.Get("identity_ids")

	var days = d.Get("filter_history_days").(int)
	var start = d.Get("filter_history_start_time_millis").(int)
	var end = d.Get("filter_history_end_time_millis").(int)

	if days != 0 {
		logger.Debug("msg", "Filter History Days Given", "days", days)
		payload.Filter.HistoryDays = days
	}

	if start != 0 && end != 0 {
		logger.Debug("msg", "Filter History Bounds Given")
		payload.Filter.HistoryDuration = &models.HistoryDuration{
			StartTime: start,
			EndTime:   end,
		}
	}

	payload.Filter.PreserveReads = d.Get("filter_preserve_reads").(bool)

	var scope = d.Get("request_params_scope")
	var resource = d.Get("request_params_resource")
	var resources = d.Get("request_params_resources")
	var condition = d.Get("request_params_condition")

	resourcesString := fmt.Sprintf("%v", resources)

	logger.Debug(
		"scope",
		scope.(string),
		"resource",
		resource.(string),
		"resources",
		resourcesString,
		"condition",
		condition.(string),
	)

	if scope == "" && resource == "" && resourcesString == "[]" && condition == "" {
	} else {
		var requestParams models.RequestParams

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

		if resourcesString == "[]" {
			requestParams.Resources = nil
		} else {
			requestParams.Resources = resources
		}

		if condition.(string) == "" {
			requestParams.Condition = nil
		} else {
			requestParams.Condition = condition.(string)
		}

		payload.RequestParams = &requestParams
	}
	return payload
}
