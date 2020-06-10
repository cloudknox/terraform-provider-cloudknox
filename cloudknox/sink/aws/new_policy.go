package aws

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"
	"fmt"
	"io/ioutil"
)

type PolicyContractWriter struct {
	Args map[string]string
}

func (aws PolicyContractWriter) Write() error {
	logger := common.GetLogger()
	logger.Info("msg", "Writing AWS Policy")

	template := fmt.Sprintf(
		`resource "aws_iam_policy" "%s" {
	    name        = "%s"
		path        = "%s"
		description = "%s"
		policy = <<EOF
		%s'`, aws.Args["name"], aws.Args["name"], aws.Args["aws_path"], aws.Args["description"], aws.Args["data"])

	suffix := "\nEOF\n}"

	filename := fmt.Sprintf("%scloudknox-aws-%s.tf", aws.Args["output_path"], aws.Args["name"])

	err := ioutil.WriteFile(filename, []byte(template+suffix), 0644)

	if err != nil {
		logger.Error("msg", "FileIO Error", "file_error", err)
		return err
	}
	return nil
}
