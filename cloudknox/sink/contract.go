package sink

import (
	"errors"
	"strings"
	"terraform-provider-cloudknox/cloudknox/common"
	"terraform-provider-cloudknox/cloudknox/sink/aws"
	"terraform-provider-cloudknox/cloudknox/sink/azure"
	"terraform-provider-cloudknox/cloudknox/sink/gcp"
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
	case common.NewPolicy:
		logger.Info("resource", common.NewPolicy)
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
	case AWS:
		logger.Info("platform", AWS)
		var aws = aws.PolicyContractWriter{Args: args}
		return aws, nil
	case AZURE:
		logger.Info("platform", AZURE)
		var azure = azure.PolicyContractWriter{Args: args}
		return azure, nil
	case GCP:
		logger.Info("platform", GCP)
		var gcp = gcp.PolicyContractWriter{Args: args}
		return gcp, nil
	case VCENTER:
		logger.Info("platform", VCENTER)
		return nil, nil
	}

	return nil, nil
}
