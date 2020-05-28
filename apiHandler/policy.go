package apiHandler

import (
	"cloudknox/terraform-provider-cloudknox/common"
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

const (
	NewPolicyURL = "https://olympus.aws-staging.cloudknox.io/api/v2/role-policy/new"
)

func NewPolicy(name string, outputPath string, payload *PolicyData) error {
	log := common.GetLogger()
	log.Info("apiHandler is getting client")
	client, err := common.GetClient()
	if err != nil {
		log.Info(err)
		return err
	}
	log.Info("Payload pre marshall")
	payload_bytes, _ := json.Marshal(payload)
	log.Info(string(payload_bytes))
	log.Info("Payload has been marshalled")

	policy, err := MakePOSTRequest(client.AccessToken, NewPolicyURL, payload_bytes)
	if err != nil {
		log.Info(err)
		return err
	}
	log.Info(policy)
	err = writePolicy(name, outputPath, policy)

	if err != nil {
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
