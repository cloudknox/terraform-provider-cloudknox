package vcenter

import "terraform-provider-cloudknox/cloudknox/common"

type PolicyContractWriter struct {
	Args map[string]string
}

func (vCenter PolicyContractWriter) Write() error {
	logger := common.GetLogger()
	logger.Info("msg", "writing vCenter policy")
	return nil
}
