package gcp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"terraform-provider-cloudknox/cloudknox/common"
)

type RolePolicyContractWriter struct {
	Args map[string]string
}

func (gcp RolePolicyContractWriter) Write() error {

	logger := common.GetLogger()
	logger.Info("msg", "writing gcp role")

	//Turn the given policy into a map so that we can extract even more fields
	logger.Debug("msg", "creating policy map")
	policy := make(map[string]interface{})

	err := json.Unmarshal([]byte(gcp.Args["data"]), &policy)

	if err != nil {
		logger.Error("msg", "unable to extract response from body", "unmarshal_error", err)
		logger.Error("role", gcp.Args["data"])
		return err
	}

	//Extract the permissions from the policy map
	logger.Debug("msg", "extracting permissions from policy map")
	permissions := policy["role"].(map[string]interface{})["includedPermissions"]

	//Convert permissions to an array
	logger.Debug("msg", "converting permissions to an array", "permissions", permissions)
	permissions_arr := make([]string, 0)
	if permissions == nil {
		logger.Debug("msg", "no permissions found")
	} else {
		for _, v := range permissions.([]interface{}) {
			permissions_arr = append(permissions_arr, v.(string))
		}
	}

	//Format the permissions array into a string with new lines after every permission
	logger.Debug("msg", "formatting permissions string")
	permissions_str := linePrint(permissions_arr)

	//Create the template for the resource
	template := fmt.Sprintf(
		`resource "google_project_iam_custom_role" "%s" {
		role_id     = "%s"
		title		= "%s"
		description = "%s"
		permissions = [%s
		`, gcp.Args["name"], policy["roleId"], gcp.Args["name"], gcp.Args["description"], permissions_str)

	suffix := "]\r\n}"

	//Write the template to file after filling out the fields

	filename := fmt.Sprintf("%s%s.tf", gcp.Args["output_path"], gcp.Args["name"])

	logger.Debug("msg", "writing template to file")
	err = ioutil.WriteFile(filename, []byte(template+suffix), 0644)

	if err != nil {
		logger.Error("msg", "fileIO error", "file_error", err)
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
