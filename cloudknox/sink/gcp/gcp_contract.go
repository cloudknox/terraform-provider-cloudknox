package gcp

import "cloudknox/terraform-provider-cloudknox/cloudknox/common"

type ContractWriter struct {
}

func (gcp *ContractWriter) WritePolicy() error {
	logger := common.GetLogger()
	logger.Info("msg", "Writing GCP Policy")
	return nil
}
