package helpers

import (
	"github.com/cloudknox/terraform-provider-cloudknox/cloudknox/common"
	"github.com/cloudknox/terraform-provider-cloudknox/cloudknox/sink"
)

// WriteResource interfaces with sinks to create desirec local output
func WriteResource(resource string, platform string, args map[string]string) error {
	logger := common.GetLogger()
	logger.Debug("msg", "preparing writer", "resource", resource, "platform", platform)
	contract, err := sink.BuildContractWriter(resource, platform, args)
	if err != nil {
		logger.Error("msg", "error while building contract", "contract_error", err)
		return err
	}
	err = contract.Write()
	if err != nil {
		logger.Error("msg", "error while writing policy", "fileio_error", err)
		return err
	}
	return nil
}
