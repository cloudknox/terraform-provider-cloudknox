package apiHandler

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"
	"cloudknox/terraform-provider-cloudknox/cloudknox/sink"
	"encoding/json"
	"time"
)

type PolicyData struct {
	AuthSystemInfo struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"authSystemInfo"`
	IdentityType string      `json:"identityType"`
	IdentityIds  interface{} `json:"identityIds"`
	Filter       struct {
		HistoryDays     int  `json:"historyDays"`
		PreserveReads   bool `json:"preserveReads"`
		HistoryDuration struct {
			StartTime int `json:"startTime"`
			EndTime   int `json:"endTime"`
		} `json:"historyDuration"`
	} `json:"filter"`
	RequestParams struct {
		Scope     string      `json:"scope"`
		Resource  string      `json:"resource"`
		Resources interface{} `json:"resources"`
		Condition string      `json:"condition"`
	}
}

func NewPolicy(platform string, name string, outputPath string, payload *PolicyData) error {
	logger := common.GetLogger()

	logger.Info("msg", "Creating New Policy", "name", name, "output_path", outputPath)

	client, err := common.GetClient()
	if err != nil {
		logger.Error("msg", "Unable to Get Client Access Token", "client_error", err.Error())
		return err
	}
	logger.Debug("msg", "Payload pre-marshal")
	payload_bytes, _ := json.Marshal(payload)
	logger.Debug("msg", "Payload post-marshal", "payload", string(payload_bytes))

	policy, err := client.POST(common.NEW_POLICY(), payload_bytes)
	if err != nil {
		logger.Error("msg", "Unable to make POST Request", "post_error", err.Error())
		return err
	} else {
		logger.Info("msg", "Post Request Successful")
	}

	logger.Info("msg", "Writing Policy to Disk", "filename", outputPath+name)
	err = writePolicy(platform, name, outputPath, policy)

	if err != nil {
		logger.Error("msg", "Unable to Write Policy", "write_error", err.Error())
		return err
	}

	return nil
}

func writePolicy(platform string, name string, outputPath string, policy map[string]interface{}) error {

	logger := common.GetLogger()

	jsonString, err := json.MarshalIndent(policy["data"], "\t", "\t")

	if err != nil {
		logger.Error("msg", "JSON Marshaling Error while Preparing Policy", "json_error", err)
	}

	args := map[string]string{
		"name":        name,
		"description": "Cloudknox Generated IAM Policy for " + platform + " at " + time.Now().String(),
		"output_path": outputPath,
		"aws_path":    "/",
		"policy":      string(jsonString),
	}

	contract, err := sink.BuildContract(platform, args)

	if err != nil {
		logger.Error("msg", "Error while Building Contract", "contract_error", err)
	}

	err = contract.WritePolicy()

	if err != nil {
		logger.Error("msg", "Error while Writing Policy", "fileio_error", err)
	}

	return nil
}
