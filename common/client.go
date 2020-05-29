package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/log/level"
)

type ClientParameters struct {
	SharedCredentialsFile string
	Profile               string
}

type Credentials struct {
	ServiceAccountID string `json:"serviceAccountId"`
	AccessKey        string `json:"accessKey"`
	SecretKey        string `json:"secretKey"`
}

type Client struct {
	AccessToken string
}

/* Private Variables */
var credentials *Credentials
var configType string

var client *Client
var clientErr error

const (
	AuthURL = "https://olympus.aws-staging.cloudknox.io/api/v2/service-account/authenticate"
)

func credentialsToJSON(credentials *Credentials) []byte {
	c, _ := json.Marshal(credentials)
	return c
}

/* Private Functions */
func buildClient(credentials *Credentials, configurationType string) {
	logger := GetLogger()
	level.Info(logger).Log("msg", "Building Client", "config_type", configurationType)

	configType = configurationType

	// Make POST Request for API Token

	// Setup HTTP Request

	// Parameters
	var jsonBytes = credentialsToJSON(credentials)

	// Request Configuration
	req, err := http.NewRequest("POST", AuthURL, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	// Setup Client and Make Request
	hclient := &http.Client{}
	resp, err := hclient.Do(req)
	if err != nil {
		level.Error(logger).Log("resp", resp, "http_error", err.Error())
		client = nil
		clientErr = errors.New("Unable to make HTTP Client Request")
		return
	}
	defer resp.Body.Close()

	// Get Response
	level.Info(logger).Log("msg", "Got HTTP Response")
	if resp.StatusCode != http.StatusOK {
		level.Error(logger).Log("msg", "HTTP Response status != 200 OK", "resp", resp.Status, "credentials", "invalid")
		client = nil
		clientErr = errors.New("Invalid Credentials")
		return
	} else {
		level.Info(logger).Log("msg", "HTTP Response status == 200 OK", "resp", resp.Status, "credentials", "valid")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	jsonBody := string(body)

	// Create Map from Body of Response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonBody), &responseMap)

	if err != nil {
		level.Error(logger).Log("msg", "Unable to extract response from body", "unmarshal_error", err)
		client = nil
		clientErr = errors.New("Unable to read HTTP Response")
		return
	}

	var accessToken = responseMap["accessToken"].(string)

	client = &Client{
		AccessToken: accessToken,
	}

	return
}

// func authorize() {
// 	TODO
// }

/* Public Functions */
func GetClient() (*Client, error) {

	if clientErr == nil {
		if client != nil {
			return client, nil
		} else {
			return nil, errors.New("Unexpected Error")
		}
	} else {
		return nil, errors.New(clientErr.Error() + " | ConfigType: " + configType)
	}

}
