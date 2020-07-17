package apiHandler

import (
	"encoding/json"
	"terraform-provider-cloudknox/cloudknox/common"
	"terraform-provider-cloudknox/cloudknox/utils"
	"time"
)

const (
	RolePolicyCreateRoute string = "/api/v2/role-policy/new"
)

func CreateRolePolicy(platform string, name string, outputPath string, payload *RolePolicyData) error {
	logger := common.GetLogger()

	logger.Info("msg", "creating new role-policy", "name", "cloudknox-"+platform+"-"+name+".tf", "output_path", outputPath)

	client, err := common.GetClient()
	if err != nil {
		logger.Error("msg", "unable to get client access token", "client_error", err.Error())
		return err
	}
	// logger.Debug("msg", "Payload pre-marshal")
	payload_bytes, _ := json.Marshal(payload)
	// logger.Debug("msg", "Payload post-marshal", "payload", string(payload_bytes))

	url := "https://olympus.aws-staging.cloudknox.io" + RolePolicyCreateRoute

	rolePolicy, err := client.POST(url, payload_bytes)
	if err != nil {
		logger.Error("msg", "unable to complete POST request", "error", err.Error())
		return err
	} else {
		logger.Info("msg", "post request successful")
	}

	rolePolicyJsonBytes, err := json.Marshal(rolePolicy["data"])
	rolePolicyJsonString := string(rolePolicyJsonBytes)

	logger.Debug("rolePolicyJsonString", utils.Truncate(rolePolicyJsonString, 30, true))

	if err != nil {
		logger.Error("msg", "JSON marshaling error while preparing data", "json_error", err)
	}

	args := map[string]string{
		"name":        name,
		"description": "Cloudknox Generated IAM Role-Policy for " + platform + " at " + time.Now().String(),
		"output_path": outputPath,
		"aws_path":    "/",
		"data":        rolePolicyJsonString,
	}

	logger.Info("msg", "Begin Write Sequence")
	err = writeResource("cloudknox_role_policy", platform, args)

	if err != nil {
		logger.Error("msg", "unable to write role_policy", "write_error", err.Error())
		return err
	}

	logger.Info("msg", "write sequence completed successfully")

	return nil
}
