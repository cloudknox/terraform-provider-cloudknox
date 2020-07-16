package vcenter

import "terraform-provider-cloudknox/cloudknox/common"

type RolePolicyContractWriter struct {
	Args map[string]string
}

func (vCenter RolePolicyContractWriter) Write() error {
	logger := common.GetLogger()
	logger.Info("msg", "writing vCenter role")
	return nil
}
