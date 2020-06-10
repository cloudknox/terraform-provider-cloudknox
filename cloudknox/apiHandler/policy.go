package apiHandler

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"
	"encoding/json"
	"time"
)

func NewPolicy(platform string, name string, outputPath string, payload *PolicyData) error {
	logger := common.GetLogger()

	logger.Info("msg", "Creating New Policy", "name", "cloudknox-"+platform+"-"+name+".tf", "output_path", outputPath)

	client, err := common.GetClient()
	if err != nil {
		logger.Error("msg", "Unable to Get Client Access Token", "client_error", err.Error())
		return err
	}
	// logger.Debug("msg", "Payload pre-marshal")
	payload_bytes, _ := json.Marshal(payload)
	// logger.Debug("msg", "Payload post-marshal", "payload", string(payload_bytes))

	policy, err := client.POST(common.NEW_POLICY(), payload_bytes)
	if err != nil {
		logger.Error("msg", "Unable to complete POST Request", "error", err.Error())
		return err
	} else {
		logger.Info("msg", "Post Request Successful")
	}

	policyJsonString, err := json.MarshalIndent(policy["data"], "\t", "\t")

	// logger.Debug("payload", jsonString)

	if err != nil {
		logger.Error("msg", "JSON Marshaling Error While Preparing Data", "json_error", err)
	}

	args := map[string]string{
		"name":        name,
		"description": "Cloudknox Generated IAM Policy for " + platform + " at " + time.Now().String(),
		"output_path": outputPath,
		"aws_path":    "/",
		"data":        string(policyJsonString),
	}

	logger.Info("msg", "Begin Write Sequence")
	err = writeResource("cloudknox_policy", platform, args)

	if err != nil {
		logger.Error("msg", "Unable to Write Policy", "write_error", err.Error())
		return err
	}

	logger.Info("msg", "Write Sequence Completed Successfully")

	return nil
}
