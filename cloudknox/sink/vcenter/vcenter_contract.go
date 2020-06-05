package vcenter

import "cloudknox/terraform-provider-cloudknox/cloudknox/common"

type ContractWriter struct {
}

func (vCenter ContractWriter) WritePolicy() error {
	logger := common.GetLogger()
	logger.Info("msg", "Writing vCenter Policy")
	return nil
}
