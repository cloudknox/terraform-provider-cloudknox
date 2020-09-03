package common

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/go-kit/kit/log"
)

// CustomLogger is a wrapper for go-kit's kit logger
type CustomLogger struct {
	logger log.Logger
}

// ClientParameters holds parameters required to create a client to interact with Cloudknox resources
type ClientParameters struct {
	SharedCredentialsFile string
	Profile               string
}

// UpdateProfile ensures a profile is set for the client credentials
func (c *ClientParameters) UpdateProfile() {
	if c.Profile == "" {
		c.Profile = "default"
	} else {
		c.Profile = strings.ToLower(c.Profile)
	}
}

// Credentials holds parameters required to recieve an accessToken
type Credentials struct {
	ServiceAccountID string `json:"serviceAccountId"`
	AccessKey        string `json:"accessKey"`
	SecretKey        string `json:"secretKey"`
}

// HTTPClient is a single method interface used to perform HTTP actions
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the struct used to interface with the CloudKnox API
type Client struct {
	AccessToken      string
	APIID            string
	ServiceAccountID string
	BaseURL          *url.URL
	httpClient       HTTPClient
}
