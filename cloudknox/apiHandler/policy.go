package apiHandler

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"
	"encoding/json"
	"io/ioutil"
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

func NewPolicy(name string, outputPath string, payload *PolicyData) error {
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
	err = writePolicy(name, outputPath, policy)

	if err != nil {
		logger.Error("msg", "Unable to Write Policy", "write_error", err.Error())
		return err
	}

	return nil
}

func writePolicy(name string, outputPath string, policy map[string]interface{}) error {
	jsonString, err := json.MarshalIndent(policy["data"], "", "\t")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outputPath+name, []byte(jsonString), 0644)

	if err != nil {
		return err
	}
	return nil
}
