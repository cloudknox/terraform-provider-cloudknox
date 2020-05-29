package apiHandler

import (
	"cloudknox/terraform-provider-cloudknox/common"
	"encoding/json"
	"io/ioutil"

	"github.com/go-kit/kit/log/level"
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

const (
	NewPolicyURL = "https://olympus.aws-staging.cloudknox.io/api/v2/role-policy/new"
)

func NewPolicy(name string, outputPath string, payload *PolicyData) error {
	logger := common.GetLogger()

	level.Info(logger).Log("msg", "Creating New Policy", "name", name, "output_path", outputPath)

	client, err := common.GetClient()
	if err != nil {
		level.Error(logger).Log("msg", "Unable to Get Client Access Token", "client_error", err.Error())
		return err
	}
	level.Debug(logger).Log("msg", "Payload pre-marshal")
	payload_bytes, _ := json.Marshal(payload)
	level.Debug(logger).Log("msg", "Payload post-marshal", "payload", string(payload_bytes))

	policy, err := MakePOSTRequest(client.AccessToken, NewPolicyURL, payload_bytes)
	if err != nil {
		level.Error(logger).Log("msg", "Unable to make POST Request", "post_error", err.Error())
		return err
	} else {
		level.Info(logger).Log("msg", "Post Request Successful", "post_error")
	}

	level.Info(logger).Log("msg", "Writing Policy", "filename", outputPath)
	err = writePolicy(name, outputPath, policy)

	if err != nil {
		level.Error(logger).Log("msg", "Unable to Write Policy", "write_error", err.Error())
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
