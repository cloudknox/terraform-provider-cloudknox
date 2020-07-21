package common

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kit/kit/log"
)

//CustomLogger is a wrapper for go-kit's kit logger
type CustomLogger struct {
	logger log.Logger
}

//ClientParameters holds parameters required to create a client to interact with Cloudknox resources
type ClientParameters struct {
	SharedCredentialsFile string
	Profile               string
}

func (c *ClientParameters) UpdateProfile() {
	if c.Profile == "" {
		c.Profile = "default"
	} else {
		c.Profile = strings.ToLower(c.Profile)
	}
}

//Credentials holds parameters required to recieve an accessToken
type Credentials struct {
	ServiceAccountID string `json:"serviceAccountId"`
	AccessKey        string `json:"accessKey"`
	SecretKey        string `json:"secretKey"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	AccessToken string
	BaseURL     *url.URL
	httpClient  HttpClient
}
