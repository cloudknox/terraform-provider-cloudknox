package aws

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"
	"fmt"
	"io/ioutil"
)

type ContractWriter struct {
	Name        string
	OutputPath  string
	AWSPath     string
	Description string
	Policy      string
}

func (aws ContractWriter) WritePolicy() error {
	logger := common.GetLogger()
	logger.Info("msg", "Writing AWS Policy")

	template := fmt.Sprintf(
		`resource "aws_iam_policy" "%s" {
	    name        = "%s"
		path        = "%s"
		description = "%s"
		policy = <<EOF
		%s'`, aws.Name, aws.Name, aws.AWSPath, aws.Description, aws.Policy)

	suffix := "\nEOF\n}"

	err := ioutil.WriteFile(aws.OutputPath+"cloudknox-aws-"+aws.Name+".tf", []byte(template+suffix), 0644)

	if err != nil {
		logger.Error("msg", "FileIO Error", "file_error", err)
		return err
	}
	return nil
}
