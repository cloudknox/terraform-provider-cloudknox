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

func setConfiguration(resource_path string) {
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

			logger.Debug("base_url", BASE_URL())

			logger.Debug("routes.authentication", AUTH())

			logger.Debug("routes.policy.create", NEW_POLICY())

		},
	)

}

func BASE_URL() string {
	return configuration.BaseURL
}

func AUTH() string {
	return BASE_URL() + configuration.Routes.Auth
}

func NEW_POLICY() string {
	return BASE_URL() + configuration.Routes.Policy.Create
}
