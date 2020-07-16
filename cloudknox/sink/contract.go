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
	logger.Info("msg", "getting contract")

	resource = strings.ToLower(resource)
	platform = strings.ToLower(platform)

	switch resource {
	case common.RolePolicy:
		logger.Info("resource", common.RolePolicy)
		return getRolePolicyContract(platform, args)
	default:
		logger.Error("msg", "invalid resource", "resource", "default")
	}

	return nil, errors.New("Invalid Platform")
}

func getRolePolicyContract(platform string, args map[string]string) (ContractWriter, error) {
	logger := common.GetLogger()
	logger.Debug("msg", "getting contract associated with platform for role_policy resource")
	switch platform {
	case AWS:
		logger.Info("platform", AWS)
		var aws = aws.RolePolicyContractWriter{Args: args}
		return aws, nil
	case AZURE:
		logger.Info("platform", AZURE)
		var azure = azure.RolePolicyContractWriter{Args: args}
		return azure, nil
	case GCP:
		logger.Info("platform", GCP)
		var gcp = gcp.RolePolicyContractWriter{Args: args}
		return gcp, nil
	case VCENTER:
		logger.Info("platform", VCENTER)
		return nil, nil
	}

	return nil, nil
}
