package azure

import "cloudknox/terraform-provider-cloudknox/cloudknox/common"

type ContractWriter struct {
}

func (azure *ContractWriter) WritePolicy() error {
	logger := common.GetLogger()
	logger.Info("msg", "Writing Azure Policy")
	return nil
}
