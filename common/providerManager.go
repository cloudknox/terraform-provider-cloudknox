package common

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
	oauth2 "golang.org/x/oauth2"
)

type Client struct {
	token oauth2.Token
}

/* Private Variables */
var client *Client
var clientOnce sync.Once

var config *viper.Viper
var configOnce sync.Once

var logger *log.Logger
var loggerOnce sync.Once

var oAuthConfig *oauth2.Config

const (
	ConfFileDefaultPath = "/cloudknox"
	ConfFileName        = "terraform-provider-cloudknox-config"
)

// func getClient() *Client {

// 	//read config once

// 	//popualte the client object once

// 	clientOnce.Do(
// 		func() {
// 			//create client struct, populate
// 			oAuthConfig = &oauth2.Config{
// 				ClientID:     ClientID,
// 				ClientSecret: ClientSecret,
// 				Scopes:       []string{"all"}, //Unsure about this line here
// 				RedirectURL:  Hostname + "/oauth2",
// 				Endpoint: oauth2.Endpoint{
// 					AuthURL:  Hostname + "/authorize",
// 					TokenURL: Hostname + "/token",
// 				},
// 			}

// 			//OAuth2.0 Exchange Goes here

// 			//t, err := oAuthConfig.Exchange(context.Background())

// 			//Build the Client Struct

// 			client = &Client{
// 				token: nil,
// 			}

// 		},
// 	)

// 	return client

// }

func getLogger() *log.Logger {
	loggerOnce.Do(
		func() {
			/* Initialize Logger */
			logger = log.New()
			// formatter := &log.TextFormatter{
			// 	DisableColors: false,
			// 	FullTimestamp: true,
			// }
			logger.SetFormatter(&log.JSONFormatter{})
			//ogger.SetLevel(log.DebugLevel)

			file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				logger.Fatal(err)
			}

			defer file.Close()

			logger.SetOutput(file)

			logger.Info("Successfully Created Logger Instance!")
		},
	)
	return logger
}

// func getConfiguration(configType string, paramters map[string]) *viper.Viper {
// 	log := GetLogger()

// 	switch configType {
// 	case "static":
// 		log.Info("Using static terraform properties")
// 	case "credentials_default_location":
// 		log.Info("Using default credentials file location")
// 	case "environment":
// 		log.Info("Using environment variables")
// 	default:
// 		log.Panic("Improper properties configurtation")
// 	}

// 	configOnce.Do(
// 		func() {
// 			/* Initialize Configuration */

// 			log.Info("Creating Viper Configuration")
// 			config = viper.New()
// 			confFullPath := fmt.Sprintf("%s/%s.yaml", ConfFilePath, ConfFileName)
// 			if utils.CheckIfPathExists(confFullPath) {
// 				log.Info("Config filepath exists")
// 				config.SetConfigType("yaml")
// 				config.SetConfigName(ConfFileName)
// 				config.AddConfigPath(ConfFilePath)
// 			} else { //Why is this going into this case?
// 				log.Info("Config filepath does not exist. Creating config path!")
// 				config.SetConfigType("yaml")
// 				config.SetConfigName("terraform-provider-cloudknox-config")
// 				config.AddConfigPath("./conf")
// 			}

// 			if err := config.ReadInConfig(); err != nil {
// 				logrus.Panicf("Failed To Load Configuration: %s", err)
// 				os.Exit(1)
// 			}
// 		},
// 	)
// 	return config
// }

/* Public Variables */
var ClientID string
var ClientSecret string
var Hostname string

/* Public Functions */
// func GetClient() *Client {
// 	return getClient()
// }

func GetLogger() *logrus.Logger {
	return getLogger()
}
