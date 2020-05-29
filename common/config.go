package common

import (
	"cloudknox/terraform-provider-cloudknox/utils"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/go-akka/configuration"
	"github.com/go-kit/kit/log/level"
	"github.com/mitchellh/go-homedir"
)

/* Private Variables */
var configOnce sync.Once
var creds *Credentials

func setConfiguration(parameters *ClientParameters) {
	configOnce.Do(
		func() {
			/* Initialize Configuration */

			var configurationType string
			creds = new(Credentials)
			logger := GetLogger()
			level.Info(logger).Log("msg", "Identifying Configuration Type")
			// === Configuration Hierarchy ===
			// Static Credentials
			// Shared Credentials File (Profile)
			// Shared Credentials File (Default)
			// Default Path Credentials File (Profile)
			// Default Path Credentials File (Default)
			// Environment Variables
			// No Credentials Provided Panic

			// Set Default Value for Profile if not provided
			if parameters.Profile == "" {
				parameters.Profile = "default"
			} else {
				parameters.Profile = strings.ToLower(parameters.Profile)
			}

			// Check Shared Credentials File
			if parameters.SharedCredentialsFile == "" {
				level.Warn(logger).Log("msg", "Shared Credentials File Not Provided")
			} else {
				level.Info(logger).Log("msg", "Searching for Shared Credentials File", "path", parameters.SharedCredentialsFile)

				if utils.CheckIfPathExists(parameters.SharedCredentialsFile) {
					level.Info(logger).Log("msg", "Shared Credentials File exists")
					level.Info(logger).Log("msg", "Checking Profile", "profile", parameters.Profile)

					err := readHOCON(parameters.SharedCredentialsFile, parameters.Profile)

					if err == nil {
						configurationType = "Shared Credentials File"
						buildClient(creds, configurationType)
						return
					} else {
						level.Error(logger).Log("msg", "Unable to Read HOCON File", "hocon_parse_error", err.Error())
					}
				}
			}

			// Check Default Path
			homedir, _ := homedir.Dir()
			defaultCredentialsPath := homedir + "//.cnx//creds.conf"

			if utils.CheckIfPathExists(defaultCredentialsPath) {
				level.Info(logger).Log("msg", "Default Credentials File Exists", "path", defaultCredentialsPath)
				level.Info(logger).Log("msg", "Checking Profile", "profile", parameters.Profile)

				err := readHOCON(defaultCredentialsPath, parameters.Profile)

				if err == nil {
					configurationType = "Default Credentials File"
					buildClient(creds, configurationType)
					return
				} else {
					level.Error(logger).Log("msg", "Unable to Read HOCON File", "hocon_parse_error", err.Error())
				}
			} else {
				level.Warn(logger).Log("msg", "Default Credentials File Not Provided")
				level.Info(logger).Log("msg", "Checking Environment Variables")
			}

			// Check Environment Variables
			if os.Getenv("CNX_SERVICE_ACCOUNT_ID") == "" || os.Getenv("CNX_ACCESS_KEY") == "" || os.Getenv("CNX_SECRET_KEY") == "" {
				level.Warn(logger).Log("msg", "All Enviornment Variables Not Correctly Set")
				level.Error(logger).Log("msg", "No Credentials Exist")
				client = nil
				clientErr = errors.New("No Credentials Found")
				return
			} else {
				level.Info(logger).Log("msg", "Environment Variables Located")

				creds.ServiceAccountID = os.Getenv("CNX_SERVICE_ACCOUNT_ID")
				creds.AccessKey = os.Getenv("CNX_ACCESS_KEY")
				creds.SecretKey = os.Getenv("CNX_SECRET_KEY")

				configurationType = "Environment Variables"

				buildClient(creds, configurationType)

				return

			}

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

	conf := configuration.ParseString(text)

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

func SetConfiguration(parameters *ClientParameters) {
	setConfiguration(parameters)
}
