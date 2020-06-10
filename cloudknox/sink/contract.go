package sink

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"
	"cloudknox/terraform-provider-cloudknox/cloudknox/sink/aws"
	"cloudknox/terraform-provider-cloudknox/cloudknox/sink/azure"
	"cloudknox/terraform-provider-cloudknox/cloudknox/sink/gcp"
	"errors"
	"strings"
)

type ContractWriter interface {
	Write() error
}

func BuildContractWriter(resource string, platform string, args map[string]string) (ContractWriter, error) {
	logger := common.GetLogger()
	logger.Info("msg", "Getting Contract")

	resource = strings.ToLower(resource)
	platform = strings.ToLower(platform)

	switch resource {
	case "cloudknox_policy":
		logger.Info("resource", "cloudknox_policy")
		return getPolicyContract(platform, args)
	default:
		logger.Error("msg", "Invalid Resource", "resource", "default")
	}

	return nil, errors.New("Invalid Platform")
}

func getPolicyContract(platform string, args map[string]string) (ContractWriter, error) {
	logger := common.GetLogger()
	logger.Debug("msg", "Getting contract associated with platform for policy resource")
	switch platform {
	case "aws":
		logger.Info("platform", "aws")
		var aws = aws.PolicyContractWriter{Args: args}
		return aws, nil
	case "azure":
		logger.Info("platform", "azure")
		var azure = azure.PolicyContractWriter{Args: args}
		return azure, nil
	case "gcp":
		logger.Info("platform", "gcp")
		var gcp = gcp.PolicyContractWriter{Args: args}
		return gcp, nil
	case "vcenter":
		logger.Info("platform", "vcenter")
		return nil, nil
	}

	return nil, nil
}
