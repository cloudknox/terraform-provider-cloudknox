package azure

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

func (azure ContractWriter) WritePolicy() error {
	logger := common.GetLogger()
	logger.Info("msg", "Writing Azure Policy")

	//Turn the given policy into a map so that we can extract even more fields
	policy := make(map[string]interface{})

	err := json.Unmarshal([]byte(azure.Policy), &policy)

	if err != nil {
		logger.Error("msg", "Unable to extract response from body", "unmarshal_error", err)
		logger.Error("policy", azure.Policy)
		return err
	}

	actions := policy["Actions"]

	//Convert actions to an array
	actions_arr := make([]string, 0)
	for _, v := range actions.([]interface{}) {
		actions_arr = append(actions_arr, v.(string))
	}

	actions_str := linePrint(actions_arr)

	scopes := policy["AssignableScopes"]

	//Convert scopes to an array
	scopes_arr := make([]string, 0)
	for _, v := range scopes.([]interface{}) {
		scopes_arr = append(scopes_arr, v.(string))
	}

	scopes_str := linePrint(scopes_arr)

	name := policy["Name"]

	//We set the resource scope to the first available scope, reccomended by AZURE terraform provider
	template := fmt.Sprintf(
		`resource "azurerm_role_definition" "%s" {
			name        = "%s"
			scope       = "%s"
			description = "%s"
		  
			permissions {
			  actions     = [%s

			  ]
			  not_actions = []
			}
		  
			assignable_scopes = [%s
			]
		`, azure.Name, name, scopes_arr[0], azure.Description, actions_str, scopes_str)

	suffix := "\r\n}"

	//Write the template to file after filling out the fields
	err = ioutil.WriteFile(azure.OutputPath+"cloudknox-gcp-"+azure.Name+".tf", []byte(template+suffix), 0644)

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
		str += "\t\t\t\t" + fmt.Sprintf(`"%s",`, i) + "\n"
	}
	return str
}
