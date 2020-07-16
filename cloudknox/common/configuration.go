package common

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

/* Private Variables */

var configuration Configuration
var configurationError error
var configOnce sync.Once

/* Public Variables */

const (
	RolePolicy string = "cloudknox_role_policy"
)

/* Public Functions */

func SetConfiguration(resource_path string) error {
	configOnce.Do(
		func() {
			logger := GetLogger()
			logger.Debug("msg", "setting constants")

			yamlFile, err := ioutil.ReadFile(resource_path)
			if err != nil {
				logger.Error("msg", "error reading configuration file", "file_read_error", err)
				configurationError = err
			}
			err = yaml.Unmarshal(yamlFile, &configuration)
			if err != nil {
				logger.Error("msg", "unable to decode into struct", "yaml_decode_error", err)
				configurationError = err
			}

		},
	)
	return configurationError
}

func GetConfiguration() Configuration {
	return configuration
}
