package common

import (
	"cloudknox/terraform-provider-cloudknox/utils"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/go-akka/configuration"
	"github.com/mitchellh/go-homedir"
)

/* Private Variables */

var configOnce sync.Once
var creds *Credentials

func setConfiguration(parameters *ClientParameters) (*Client, error) {
	configOnce.Do(
		func() {
			/* Initialize Configuration */

			var configurationType string
			creds = new(Credentials)
			log := GetLogger()

			// log.Info(parameters.ServiceAccountID)
			// log.Info(parameters.AccessKey)
			// log.Info(parameters.SecretKey)
			// log.Info(parameters.SharedCredentialsFile)
			// log.Info(parameters.Profile)

			// === Configuration Hierarchy ===
			// Static Credentials
			// Shared Credentials File (Profile)
			// Shared Credentials File (Default)
			// Default Path Credentials File (Profile)
			// Default Path Credentials File (Default)
			// Environment Variables
			// No Credentials Provided Panic

			homedir, _ := homedir.Dir()

			defaultCredentialsPath := homedir + "//.cnx//creds.conf"
			// Check Static Credentials
			if parameters.ServiceAccountID == "" || parameters.AccessKey == "" || parameters.SecretKey == "" {
				log.Info("Static Credentials not provided or incomplete")
				if configurationType != "" {
					log.Info("Continuing to use " + configurationType)
				}
			} else {

				if configurationType == "" {
					log.Info("Static Credentials provided")

				} else {
					log.Info("Static Credentials provided, overriding " + configurationType)
				}

				configurationType = "Static Credentials"

				creds.ServiceAccountID = parameters.ServiceAccountID
				creds.AccessKey = parameters.AccessKey
				creds.SecretKey = parameters.SecretKey

				buildClient(creds, configurationType)

				return

			}

			// Set Default Value for Profile if not provided
			if parameters.Profile == "" {
				parameters.Profile = "default"
			} else {
				parameters.Profile = strings.ToLower(parameters.Profile)
			}

			// Check Shared Credentials File
			if parameters.SharedCredentialsFile == "" {
				log.Info("Custom Shared Credentials File not provided")
			} else {
				log.Info("Searching for Shared Credentials File at" + parameters.SharedCredentialsFile)

				if utils.CheckIfPathExists(parameters.SharedCredentialsFile) {
					log.Info("Shared Credentials File located")
					log.Info("Checking " + parameters.Profile + " Profile")

					err := readHOCON(parameters.SharedCredentialsFile, parameters.Profile)

					if err == nil {
						configurationType = "Shared Credentials File"
						buildClient(creds, configurationType)
						return
					} else {
						log.Info(err)
					}
				}
			}

			//Check Default Path

			if utils.CheckIfPathExists(defaultCredentialsPath) {
				log.Info("Default Credentials File located")
				log.Info("Checking " + parameters.Profile + " Profile")

				err := readHOCON(defaultCredentialsPath, parameters.Profile)

				if err == nil {
					configurationType = "Default Credentials File"
					buildClient(creds, configurationType)
					return
				} else {
					log.Info(err)
				}
			} else {
				log.Info("Default Credentials File not provided at ~/.cnx/creds.conf")
				log.Info("Checking Environment Variables")
			}

			if os.Getenv("CNX_SERVICE_ACCOUNT_ID") == "" || os.Getenv("CNX_ACCESS_KEY") == "" || os.Getenv("CNX_SECRET_KEY") == "" {
				log.Info("All Environment Variables not set")
				log.Info("No Credentials Exist")
				client = nil
				clientErr = errors.New("No Credentials Found")
				return
			} else {
				log.Info("Environment Variables Located")

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

	return client, clientErr
}

func readHOCON(path string, profile string) error {
	log := GetLogger()
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
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
