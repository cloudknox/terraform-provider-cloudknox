package common

import "github.com/go-kit/kit/log"

//CustomLogger is a wrapper for go-kit's kit logger
type CustomLogger struct {
	logger log.Logger
}

//ClientParameters holds parameters required to create a client to interact with Cloudknox resources
type ClientParameters struct {
	SharedCredentialsFile string
	Profile               string
}

//Credentials holds parameters required to recieve an accessToken
type Credentials struct {
	ServiceAccountID string `json:"serviceAccountId"`
	AccessKey        string `json:"accessKey"`
	SecretKey        string `json:"secretKey"`
}

//Client object is used to interact with client functions using an AccessToken
type Client struct {
	AccessToken string
}

//Configuration stores the parameters set in the configuration file
type Configuration struct {
	BaseURL string `yaml:"base_url"`
	Routes  struct {
		Auth   string `yaml:"authentication"`
		Policy struct {
			Create string `yaml:"create"`
		} `yaml:"policy"`
	} `yaml:"routes"`
}
