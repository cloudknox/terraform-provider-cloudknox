package gcp

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ContractWriter struct {
	Name        string
	OutputPath  string
	Description string
	Policy      string
}

func (gcp ContractWriter) WritePolicy() error {

	logger := common.GetLogger()
	logger.Info("msg", "Writing GCP Policy")

	policy := make(map[string]interface{})

	err := json.Unmarshal([]byte(gcp.Policy), &policy)

	if err != nil {
		logger.Error("msg", "Unable to extract response from body", "unmarshal_error", err)
		logger.Error("policy", gcp.Policy)
		return err
	}

	permissions := policy["role"].(map[string]interface{})["includedPermissions"]

	permissions_arr := make([]string, 0)
	for _, v := range permissions.([]interface{}) {
		permissions_arr = append(permissions_arr, v.(string))
	}

	permissions_str := linePrint(permissions_arr)

	template := fmt.Sprintf(
		`resource "google_project_iam_custom_role" "%s" {
		role_id     = "%s"
		title		= "%s"
		description = "%s"
		permissions = [%s
		`, gcp.Name, policy["roleId"], gcp.Name, gcp.Description, permissions_str)

	suffix := "]\r}"

	err = ioutil.WriteFile(gcp.OutputPath+"cloudknox-gcp-"+gcp.Name+".tf", []byte(template+suffix), 0644)

	if err != nil {
		logger.Error("msg", "FileIO Error", "file_error", err)
		return err
	}
	return nil
}

func linePrint(arr []string) string {
	var str string
	str += "\n"
	for _, i := range arr {
		str += "\t\t\t" + fmt.Sprintf(`"%s",`, i) + "\n"
	}
	return str
}
