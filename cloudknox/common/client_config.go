package common

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"terraform-provider-cloudknox/cloudknox/utils"

	config "github.com/go-akka/configuration"
	"github.com/mitchellh/go-homedir"
)

/* Private Variables */
var clientConfigOnce sync.Once
var creds *Credentials

/* Private Functions */
func setClientConfiguration(parameters *ClientParameters) {
	clientConfigOnce.Do(
		func() {
			/* Initialize Configuration */

			var configurationType string
			creds = new(Credentials)
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
			if parameters.Profile == "" {
				parameters.Profile = "default"
			} else {
				parameters.Profile = strings.ToLower(parameters.Profile)
			}

			// Check Shared Credentials File
			if parameters.SharedCredentialsFile == "" {
				logger.Warn("msg", "shared credentials file not provided")
			}
			logger.Debug("msg", "searching for shared credentials file")
			if utils.CheckIfPathExists(parameters.SharedCredentialsFile) {
				logger.Info("msg", "shared credentials file exists", "path", parameters.SharedCredentialsFile)
				logger.Debug("msg", "checking profile", "profile", parameters.Profile)

				err := readHOCON(parameters.SharedCredentialsFile, parameters.Profile)

				if err == nil {
					configurationType = "Shared Credentials File"
					buildClient(creds, configurationType)
					return
				}
				logger.Error("msg", "unable to read HOCON file", "hocon_parse_error", err.Error())
				return
			}

			// Check Default Path
			homedir, _ := homedir.Dir()
			defaultCredentialsPath := homedir + "//.cnx//creds.conf"
			logger.Debug("msg", "searching for default credentials file")
			if utils.CheckIfPathExists(defaultCredentialsPath) {
				logger.Info("msg", "default credentials file exists", "path", defaultCredentialsPath)
				logger.Debug("msg", "checking profile", "profile", parameters.Profile)

				err := readHOCON(defaultCredentialsPath, parameters.Profile)

				if err == nil {
					configurationType = "Default Credentials File"
					buildClient(creds, configurationType)
					return
				}
				logger.Error("msg", "unable to read HOCON file", "hocon_parse_error", err.Error())

			}
			logger.Warn("msg", "default credentials file not provided")
			logger.Info("msg", "checking environment variables")

			// Check Environment Variables
			if os.Getenv("CNX_SERVICE_ACCOUNT_ID") == "" || os.Getenv("CNX_ACCESS_KEY") == "" || os.Getenv("CNX_SECRET_KEY") == "" {
				logger.Warn("msg", "all enviornment variables not correctly set")
				logger.Error("msg", "no credentials exist")
				return
			}
			logger.Info("msg", "environment variables located")

			creds.ServiceAccountID = os.Getenv("CNX_SERVICE_ACCOUNT_ID")
			creds.AccessKey = os.Getenv("CNX_ACCESS_KEY")
			creds.SecretKey = os.Getenv("CNX_SECRET_KEY")

			configurationType = "Environment Variables"

			buildClient(creds, configurationType)

			return

		},
	)

	return
}

func readHOCON(path string, profile string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Convert []byte to string and print to screen
	text := string(content)

	conf := config.ParseString(text)

	creds.ServiceAccountID = conf.GetString("profiles." + profile + ".service_account_id")
	creds.AccessKey = conf.GetString("profiles." + profile + ".access_key")
	creds.SecretKey = conf.GetString("profiles." + profile + ".secret_key")

	if creds.ServiceAccountID == "" || creds.AccessKey == "" || creds.SecretKey == "" {
		return errors.New("Malformed HOCON File")
	}

	return nil
}

/* Public Variables */

/* Public Functions */

// SetClientConfiguration is the public function used to set Client Configuration
func SetClientConfiguration(parameters *ClientParameters) {
	setClientConfiguration(parameters)
}
