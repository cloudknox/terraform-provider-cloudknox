package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"terraform-provider-cloudknox/cloudknox/utils"

	config "github.com/go-akka/configuration"
	"github.com/mitchellh/go-homedir"
)

/* Private Variables */
var clientConfigOnce sync.Once
var credentials *Credentials

/* Private Functions */
func getCredentials(parameters *ClientParameters) *Credentials {
	clientConfigOnce.Do(
		func() {
			/* Initialize Configuration */
			logger := GetLogger()
			logger.Info("msg", "attempting to identifying configuration type")
			// === Configuration Hierarchy ===
			// Shared Credentials File (Profile)
			// Shared Credentials File (Default)
			// Default Path Credentials File (Profile)
			// Default Path Credentials File (Default)
			// Environment Variables
			// No Credentials Provided => Error

			// Set Default Value for Profile if not provided
			parameters.UpdateProfile()

			// Check Shared Credentials File
			if parameters.SharedCredentialsFile == "" {
				logger.Warn("msg", "shared credentials file not provided")
			}

			logger.Info("searching for shared credentials file")
			if utils.CheckIfPathExists(parameters.SharedCredentialsFile) {
				err := updateCredentialsFromFile(parameters.SharedCredentialsFile, parameters.Profile)
				if err == nil { return }
			}

			// Check Default Path
			homeDir, _ := homedir.Dir()
			defaultCredentialsPath := homeDir + "//.cloudknox//credentials.conf"
			logger.Info("searching for default credentials file")
			if utils.CheckIfPathExists(defaultCredentialsPath) {
				err := updateCredentialsFromFile(defaultCredentialsPath, parameters.Profile)
				if err == nil { return }
			}

			logger.Info("msg", "checking environment variables")
			// Check Environment Variables
			if os.Getenv("CNX_SERVICE_ACCOUNT_ID") == "" || os.Getenv("CNX_ACCESS_KEY") == "" || os.Getenv("CNX_SECRET_KEY") == "" {
				logger.Warn("msg", "all environment variables not correctly set")
				logger.Error("msg", "no credentials exist")
				return
			}
			logger.Info("msg", "environment variables located")
			credentials.ServiceAccountID = os.Getenv("CNX_SERVICE_ACCOUNT_ID")
			credentials.AccessKey = os.Getenv("CNX_ACCESS_KEY")
			credentials.SecretKey = os.Getenv("CNX_SECRET_KEY")
		},
	)
	return credentials
}

func updateCredentialsFromFile(filename string, profile string) error {
	logger := GetLogger()
	logger.Info("msg", "shared credentials file exists", "path", filename)
	logger.Debug("msg", "checking profile", "profile", profile)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// Convert []byte to string and print to screen
	text := string(content)
	conf := config.ParseString(text)
	credentials.ServiceAccountID = conf.GetString("profiles." + profile + ".service_account_id")
	credentials.AccessKey = conf.GetString("profiles." + profile + ".access_key")
	credentials.SecretKey = conf.GetString("profiles." + profile + ".secret_key")
	if credentials.ServiceAccountID == "" || credentials.AccessKey == "" || credentials.SecretKey == "" {
		err := fmt.Errorf("malformed configuration")
		logger.Error("msg", "unable to read configuration file", err.Error())
		return err
	}
	return nil
}

// SetClientConfiguration is the public function used to set Client Configuration
func GetCredentials(parameters *ClientParameters) *Credentials {
	return getCredentials(parameters)
}
