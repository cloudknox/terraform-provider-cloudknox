package azure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"terraform-provider-cloudknox/cloudknox/common"
)

type RolePolicyContractWriter struct {
	Args map[string]string
}

func (azure RolePolicyContractWriter) Write() error {
	logger := common.GetLogger()
	logger.Info("msg", "writing azure role")

	//Turn the given policy into a map so that we can extract even more fields
	policy := make(map[string]interface{})

	err := json.Unmarshal([]byte(azure.Args["data"]), &policy)

	if err != nil {
		logger.Error("msg", "unable to extract response from body", "unmarshal_error", err)
		logger.Error("role", azure.Args["data"])
		return err
	}

	var actions_str, not_actions_str string

	if policy["Actions"] != nil {
		actions := policy["Actions"]

		//Convert actions to an array
		actions_arr := make([]string, 0)
		for _, v := range actions.([]interface{}) {
			actions_arr = append(actions_arr, v.(string))
		}
		actions_str = linePrint(actions_arr)
	}

	if policy["NotActions"] != nil {
		not_actions := policy["NotActions"]

		//Convert NotActions to an array
		not_actions_arr := make([]string, 0)
		for _, v := range not_actions.([]interface{}) {
			not_actions_arr = append(not_actions_arr, v.(string))
		}
		not_actions_str = linePrint(not_actions_arr)
	}

	scopes := policy["AssignableScopes"]

	//Convert scopes to an array
	scopes_arr := make([]string, 0)
	for _, v := range scopes.([]interface{}) {
		scopes_arr = append(scopes_arr, v.(string))
	}

	scopes_str := linePrint(scopes_arr)

	policy_name := policy["Name"]

	//We set the resource scope to the first available scope, reccomended by AZURE terraform provider
	template := fmt.Sprintf(
		`resource "azurerm_role_definition" "%s" {
			name        = "%s"
			scope       = "%s"
			description = "%s"
		  
			permissions {
			  actions     = [%s
			  ]
			  not_actions = [%s
			  ]
			}
		  
			assignable_scopes = [%s
			]
		`,
		azure.Args["name"],
		policy_name,
		scopes_arr[0],
		azure.Args["description"],
		actions_str,
		not_actions_str,
		scopes_str,
	)

	suffix := "\r\n}"

	//Write the template to file after filling out the fields

	filename := fmt.Sprintf("%s%s.tf", azure.Args["output_path"], azure.Args["name"])
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
		str += "\t\t\t\t" + fmt.Sprintf(`"%s",`, i) + "\n"
	}
	return str
}
