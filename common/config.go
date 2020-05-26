package common

import (
	"cloudknox/terraform-provider-cloudknox/utils"
	"sync"

	"github.com/mitchellh/go-homedir"
)

/* Private Variables */

var configOnce sync.Once

func setConfiguration(parameters *ClientParameters) (*Client, error) {
	configOnce.Do(
		func() {
			/* Initialize Configuration */

			var configurationType string

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

			defaultCredentialspath := homedir + "/.cnx/credentials"

			if parameters.Profile == "" {
				parameters.Profile = "Default"
			}

			if parameters.SharedCredentialsFile == "" {
				log.Info("Custom Shared Credentials File not provided")
			} else {
				log.Info("Searching for Shared Credentials File at" + parameters.SharedCredentialsFile)

				if utils.CheckIfPathExists(parameters.SharedCredentialsFile) {
					log.Info("Shared Credentials File located")
					log.Info("Checking " + parameters.Profile + " Profile")
					/*
						If success
						configurationType = "SharedCredentialsFile"
						else
						continue

					*/
				}
			}

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

				buildClient(parameters.ServiceAccountID, parameters.AccessKey, parameters.SecretKey, configurationType)

				return

			}

			//Check Default Path
			if utils.CheckIfPathExists(defaultCredentialspath) {
				log.Info("Credentials file found at " + defaultCredentialspath)
				log.Info("Checking " + parameters.Profile + " Profile")
				/*
					If success
					configurationType = "DefaultCredentialsFile"
					return
					else
					continue

				*/

				return
			} else {
				log.Info("Checking Environment Variables")

				/*
					If success
					configurationType = "DefaultCredentialsFile"
					return
					else
					error

				*/

				return
			}

			return
		},
	)

	return client, clientErr
}

/* Public Variables */

/* Public Functions */

func SetConfiguration(parameters *ClientParameters) {
	setConfiguration(parameters)
}
