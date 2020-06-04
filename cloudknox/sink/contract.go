package sink

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/common"
	"cloudknox/terraform-provider-cloudknox/cloudknox/sink/aws"
	"cloudknox/terraform-provider-cloudknox/cloudknox/sink/gcp"
	"errors"
	"strings"
)

type ContractWriter interface {
	WritePolicy() error
}

func BuildContract(platform string, args map[string]string) (ContractWriter, error) {
	logger := common.GetLogger()
	logger.Info("msg", "Getting Contract")
	switch strings.ToLower(platform) {
	case "aws":
		logger.Info("platform", "aws")
		acw := aws.ContractWriter{
			Name:        args["name"],
			OutputPath:  args["output_path"],
			AWSPath:     args["aws_path"],
			Description: args["description"],
			Policy:      args["policy"],
		}
		return acw, nil
	case "azure":
		logger.Info("platform", "azure")
		return nil, nil
	case "gcp":
		logger.Info("platform", "gcp")
		gcp := gcp.ContractWriter{
			Name:        args["name"],
			OutputPath:  args["output_path"],
			Description: args["description"],
			Policy:      args["policy"],
		}
		return gcp, nil
	case "vcenter":
		logger.Info("platform", "vcenter")
		return nil, nil
	}

	return nil, errors.New("Invalid Platform")
}
