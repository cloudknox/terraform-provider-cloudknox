package common

import (
	"cloudknox/terraform-provider-cloudknox/cloudknox/utils"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	config "github.com/go-akka/configuration"
	"github.com/mitchellh/go-homedir"
)

/* Private Variables */
var clientConfigOnce sync.Once
var creds *Credentials

func setClientConfiguration(parameters *ClientParameters) {
	clientConfigOnce.Do(
		func() {
			/* Initialize Configuration */

			var configurationType string
			creds = new(Credentials)
			logger := GetLogger()
			logger.Info("msg", "Identifying Configuration Type")
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
				logger.Warn("msg", "Shared Credentials File Not Provided")
			} else {
				logger.Info("msg", "Searching for Shared Credentials File", "path", parameters.SharedCredentialsFile)

				if utils.CheckIfPathExists(parameters.SharedCredentialsFile) {
					logger.Info("msg", "Shared Credentials File exists")
					logger.Info("msg", "Checking Profile", "profile", parameters.Profile)

					err := readHOCON(parameters.SharedCredentialsFile, parameters.Profile)

					if err == nil {
						configurationType = "Shared Credentials File"
						buildClient(creds, configurationType)
						return
					} else {
						logger.Error("msg", "Unable to Read HOCON File", "hocon_parse_error", err.Error())
					}
				}
			}

			// Check Default Path
			homedir, _ := homedir.Dir()
			defaultCredentialsPath := homedir + "//.cnx//creds.conf"

			if utils.CheckIfPathExists(defaultCredentialsPath) {
				logger.Info("msg", "Default Credentials File Exists", "path", defaultCredentialsPath)
				logger.Info("msg", "Checking Profile", "profile", parameters.Profile)

				err := readHOCON(defaultCredentialsPath, parameters.Profile)

				if err == nil {
					configurationType = "Default Credentials File"
					buildClient(creds, configurationType)
					return
				} else {
					logger.Error("msg", "Unable to Read HOCON File", "hocon_parse_error", err.Error())
				}
			} else {
				logger.Warn("msg", "Default Credentials File Not Provided")
				logger.Info("msg", "Checking Environment Variables")
			}

			// Check Environment Variables
			if os.Getenv("CNX_SERVICE_ACCOUNT_ID") == "" || os.Getenv("CNX_ACCESS_KEY") == "" || os.Getenv("CNX_SECRET_KEY") == "" {
				logger.Warn("msg", "All Enviornment Variables Not Correctly Set")
				logger.Error("msg", "No Credentials Exist")
				return
			} else {
				logger.Info("msg", "Environment Variables Located")

				creds.ServiceAccountID = os.Getenv("CNX_SERVICE_ACCOUNT_ID")
				creds.AccessKey = os.Getenv("CNX_ACCESS_KEY")
				creds.SecretKey = os.Getenv("CNX_SECRET_KEY")

				configurationType = "Environment Variables"

				buildClient(creds, configurationType)

				return

			}
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

func SetClientConfiguration(parameters *ClientParameters) {
	setClientConfiguration(parameters)
}
