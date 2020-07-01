package apiHandler

import (
	"encoding/json"
	"terraform-provider-cloudknox/cloudknox/common"
	"terraform-provider-cloudknox/cloudknox/utils"
	"time"
)

func NewPolicy(platform string, name string, outputPath string, payload *PolicyData) error {
	logger := common.GetLogger()

	logger.Info("msg", "creating new policy", "name", "cloudknox-"+platform+"-"+name+".tf", "output_path", outputPath)

	client, err := common.GetClient()
	if err != nil {
		logger.Error("msg", "unable to get client access token", "client_error", err.Error())
		return err
	}
	// logger.Debug("msg", "Payload pre-marshal")
	payload_bytes, _ := json.Marshal(payload)
	// logger.Debug("msg", "Payload post-marshal", "payload", string(payload_bytes))

	url := common.GetConfiguration().BaseURL + common.GetConfiguration().Routes.Policy.Create

	policy, err := client.POST(url, payload_bytes)
	if err != nil {
		logger.Error("msg", "unable to complete POST request", "error", err.Error())
		return err
	} else {
		logger.Info("msg", "post request successful")
	}

	policyJsonBytes, err := json.Marshal(policy["data"])
	policyJsonString := string(policyJsonBytes)

	logger.Debug("policyJsonString", utils.Truncate(policyJsonString, 30, true))

	if err != nil {
		logger.Error("msg", "JSON marshaling error while preparing data", "json_error", err)
	}

	args := map[string]string{
		"name":        name,
		"description": "Cloudknox Generated IAM Policy for " + platform + " at " + time.Now().String(),
		"output_path": outputPath,
		"aws_path":    "/",
		"data":        policyJsonString,
	}

	logger.Info("msg", "Begin Write Sequence")
	err = writeResource("cloudknox_policy", platform, args)

	if err != nil {
		logger.Error("msg", "unable to write policy", "write_error", err.Error())
		return err
	}

	logger.Info("msg", "write sequence completed successfully")

	return nil
}
