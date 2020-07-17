package routes

import (
	"encoding/json"
	"terraform-provider-cloudknox/cloudknox/api/helpers"
	"terraform-provider-cloudknox/cloudknox/common"
	"terraform-provider-cloudknox/cloudknox/utils"
	"time"
)

const (
	// RolePolicyCreateRoute has the route used to get new role-policies with the CloudKnox API
	RolePolicyCreateRoute string = "/api/v2/role-policy/new"
)

// CreateRolePolicy creates a role_policy
func CreateRolePolicy(platform string, name string, outputPath string, payload *RolePolicyData) error {
	logger := common.GetLogger()

	logger.Info("msg", "creating new role-policy", "name", "cloudknox-"+platform+"-"+name+".tf", "output_path", outputPath)

	client, err := common.GetClient()
	if err != nil {
		logger.Error("msg", "unable to get client access token", "client_error", err.Error())
		return err
	}
	logger.Debug("msg", "Payload pre-marshal")
	payloadBytes, _ := json.Marshal(payload)
	logger.Debug("msg", "Payload post-marshal", "payload", string(payloadBytes))

	url := "https://olympus.aws-staging.cloudknox.io" + RolePolicyCreateRoute

	rolePolicy, err := client.POST(url, payloadBytes)
	if err != nil {
		logger.Error("msg", "unable to complete POST request", "error", err.Error())
		return err
	}
	logger.Info("msg", "post request successful")

	rolePolicyJSONBytes, err := json.Marshal(rolePolicy["data"])
	rolePolicyJSONString := string(rolePolicyJSONBytes)

	logger.Debug("rolePolicyJSONString", utils.Truncate(rolePolicyJSONString, 30, true))

	if err != nil {
		logger.Error("msg", "JSON marshaling error while preparing data", "json_error", err)
	}

	args := map[string]string{
		"name":        name,
		"description": "Cloudknox Generated IAM Role-Policy for " + platform + " at " + time.Now().String(),
		"output_path": outputPath,
		"aws_path":    "/",
		"data":        rolePolicyJSONString,
	}

	logger.Info("msg", "Begin Write Sequence")
	err = helpers.WriteResource("cloudknox_role_policy", platform, args)

	if err != nil {
		logger.Error("msg", "unable to write role_policy", "write_error", err.Error())
		return err
	}

	logger.Info("msg", "write sequence completed successfully")

	return nil
}
