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

			// Set Default Value for Profile if not provided
			parameters.UpdateProfile()

			// Check Shared Credentials File
			if parameters.SharedCredentialsFile == "" {
				logger.Warn("msg", "shared credentials file not provided in provider configuration")
			} else {
				logger.Debug("msg", "searching for shared credentials file", "path", parameters.SharedCredentialsFile)
				if utils.CheckIfPathExists(parameters.SharedCredentialsFile) {
					logger.Debug("msg", "shared credentials file found")
					err := updateCredentialsFromFile(parameters.SharedCredentialsFile, parameters.Profile)
					if err == nil {
						return
					}
				} else {
					logger.Error("msg", "provided shared credentials file does not exist")
					return
				}
			}

			// Check Default Path
			homeDir, _ := homedir.Dir()
			defaultCredentialsPath := homeDir + "//.cloudknox//credentials.conf"
			logger.Debug("msg", "searching for default credentials file")
			if utils.CheckIfPathExists(defaultCredentialsPath) {
				logger.Debug("msg", "default credentials file found")
				err := updateCredentialsFromFile(defaultCredentialsPath, parameters.Profile)
				if err == nil {
					return
				}
			} else {
				logger.Warn("msg", "default credentials file not provided")
			}

			logger.Info("msg", "checking environment variables")
			// Check Environment Variables
			if os.Getenv("CNX_SERVICE_ACCOUNT_ID") == "" || os.Getenv("CNX_ACCESS_KEY") == "" || os.Getenv("CNX_SECRET_KEY") == "" {
				logger.Warn("msg", "all environment variables not correctly set")
				logger.Error("msg", "no credentials exist")
				return
			}
			logger.Info("msg", "environment variables located")

			credentials = &Credentials{
				ServiceAccountID: os.Getenv("CNX_SERVICE_ACCOUNT_ID"),
				AccessKey:        os.Getenv("CNX_ACCESS_KEY"),
				SecretKey:        os.Getenv("CNX_SECRET_KEY"),
			}
		},
	)
	return credentials
}

func updateCredentialsFromFile(filename string, profile string) error {
	logger := GetLogger()
	logger.Debug("msg", "checking profile", "profile", profile)
	if profile == "" {
		err := fmt.Errorf("profile is not set")
		return err
	}
	logger.Debug("msg", "reading file", "filename", filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// Convert []byte to string and print to screen
	logger.Debug("msg", "reading content into text")
	text := string(content)

	logger.Debug("msg", "parsing string text")
	conf := config.ParseString(text)

	logger.Debug("msg", "reading SAI")
	serviceAccountID := conf.GetString("profiles." + profile + ".service_account_id")

	logger.Debug("msg", "reading AK")
	accessKey := conf.GetString("profiles." + profile + ".access_key")

	logger.Debug("msg", "reading SK")
	secretKey := conf.GetString("profiles." + profile + ".secret_key")

	logger.Debug("msg", "building credentials object")
	credentials = &Credentials{
		ServiceAccountID: serviceAccountID,
		AccessKey:        accessKey,
		SecretKey:        secretKey,
	}

	if credentials.ServiceAccountID == "" || credentials.AccessKey == "" || credentials.SecretKey == "" {
		err := fmt.Errorf("malformed configuration")
		logger.Error("msg", "unable to read configuration file", "error", err.Error())
		return err
	}
	return nil
}

// GetCredentials is the public function used to get a struct containing required API credentials
func GetCredentials(parameters *ClientParameters) *Credentials {
	return getCredentials(parameters)
}
