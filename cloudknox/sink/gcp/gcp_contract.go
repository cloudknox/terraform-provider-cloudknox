package gcp

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"
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
	template := fmt.Sprintf(
		`data "google_iam_policy" %s {
		binding {
			  %s
	}`, gcp.Name, gcp.Policy)

	suffix := "\n}"

	err := ioutil.WriteFile(gcp.OutputPath+"cloudknox-gcp-"+gcp.Name+".tf", []byte(template+suffix), 0644)

	if err != nil {
		logger.Error("msg", "FileIO Error", "file_error", err)
		return err
	}
	return nil
}
