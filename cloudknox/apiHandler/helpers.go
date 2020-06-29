package apiHandler

import (
	"terraform-provider-cloudknox/cloudknox/common"
	"terraform-provider-cloudknox/cloudknox/sink"
)

func writeResource(resource string, platform string, args map[string]string) error {

	logger := common.GetLogger()

	contract, err := sink.BuildContractWriter(resource, platform, args)

	if err != nil {
		logger.Error("msg", "Error while Building Contract", "contract_error", err)
		return err
	}

	err = contract.Write()

	if err != nil {
		logger.Error("msg", "Error while Writing Policy", "fileio_error", err)
		return err
	}

	return nil
}
