package aws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"terraform-provider-cloudknox/cloudknox/common"
	"terraform-provider-cloudknox/cloudknox/utils"
)

type PolicyContractWriter struct {
	Args map[string]string
}

type PolicyElement struct {
	PolicyName string      `json:"policyName"`
	Policy     interface{} `json:"policy"`
}

func (aws PolicyContractWriter) Write() error {
	logger := common.GetLogger()
	logger.Info("msg", "Writing AWS Policy")

	var policies []PolicyElement
	err := json.Unmarshal([]byte(aws.Args["data"]), &policies)

	if err != nil {
		logger.Error("error", err)
		return err
	}

	nPolicies := len(policies)
	logger.Debug("nPolicies", nPolicies)

	if err != nil {
		logger.Error("error", err)
		return err
	}

	var resource string

	for i, policy := range policies {

		// Correct i as the policies are put from last to first
		i = len(policies) - i - 1

		policyJsonBytes, err := json.MarshalIndent(policy.Policy, "\t", "\t")
		if err != nil {
			logger.Error("error", err, "policy", i)
			return err
		}
		policyJsonString := string(policyJsonBytes)

		logger.Info("policyName", policy.PolicyName, "policy", utils.Truncate(policyJsonString, 30, true))

		logger.Debug("msg", "Policy Character Count", "count", len(policyJsonString))

		if len(policyJsonString) > 6142 {
			logger.Warn("msg", "Policy character count exceeds 6142 characters")
		}

		template := fmt.Sprintf(
			`resource "aws_iam_policy" "%s" {
			name        = "%s"
			path        = "%s"
			description = "%s"
			policy = <<EOF
			%s`, policy.PolicyName, policy.PolicyName, aws.Args["aws_path"], aws.Args["description"], policyJsonString)

		suffix := "\nEOF\n}"

		resource = (template + suffix + "\n\n") + resource

	}

	filename := fmt.Sprintf("%s%s.tf", aws.Args["output_path"], aws.Args["name"])

	err = ioutil.WriteFile(filename, []byte(resource), 0644)

	if err != nil {
		logger.Error("msg", "FileIO Error", "file_error", err)
		return err
	}
	return nil
}
