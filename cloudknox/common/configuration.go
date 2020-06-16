package common

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

/* Private Variables */

var configuration Configuration
var configOnce sync.Once

/* Public Variables */

const (
	NewPolicy string = "cloudknox_policy"
)

/* Public Functions */

func SetConfiguration(resource_path string) {
	configOnce.Do(
		func() {
			logger := GetLogger()
			logger.Debug("msg", "Setting Constants")

			yamlFile, err := ioutil.ReadFile(resource_path)
			if err != nil {
				logger.Debug("msg", "Error Reading Configuration File", "file_read_error", err)
			}
			err = yaml.Unmarshal(yamlFile, &configuration)
			if err != nil {
				logger.Debug("msg", "Unable to Decode Into Struct", "yaml_decode_error", err)
			}

		},
	)

}

func GetConfiguration() Configuration {
	return configuration
}
