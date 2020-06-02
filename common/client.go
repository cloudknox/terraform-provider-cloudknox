package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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
var clientErr = errors.New("Credentials Error")

func credentialsToJSON(credentials *Credentials) []byte {
	c, _ := json.Marshal(credentials)
	return c
}

/* Private Functions */
func buildClient(credentials *Credentials, configurationType string) {
	logger := GetLogger()
	logger.Info("msg", "Building Client", "config_type", configurationType)

	configType = configurationType

	// Make POST Request for API Token

	// Setup HTTP Request

	// Parameters
	var jsonBytes = credentialsToJSON(credentials)

	// Request Configuration
	req, err := http.NewRequest("POST", AUTH(), bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	// Setup Client and Make Request
	hclient := &http.Client{}
	resp, err := hclient.Do(req)
	if err != nil {
		logger.Error("resp", resp, "http_error", err.Error())
		client = nil
		clientErr = errors.New("Unable to make HTTP Client Request")
		return
	}
	defer resp.Body.Close()

	// Get Response
	logger.Info("msg", "Got HTTP Response")
	if resp.StatusCode != http.StatusOK {
		logger.Error("msg", "HTTP Response status != 200 OK", "resp", resp.Status, "credentials", "invalid")
		client = nil
		clientErr = errors.New("Invalid Credentials")
		return
	} else {
		logger.Info("msg", "HTTP Response status == 200 OK", "resp", resp.Status, "credentials", "valid")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	jsonBody := string(body)

	// Create Map from Body of Response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonBody), &responseMap)

	if err != nil {
		logger.Error("msg", "Unable to extract response from body", "unmarshal_error", err)
		client = nil
		clientErr = errors.New("Unable to read HTTP Response")
		return
	}

	var accessToken = responseMap["accessToken"].(string)

	client = &Client{
		AccessToken: accessToken,
	}
	clientErr = nil

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

func (c *Client) POST(url string, payload []byte) (map[string]interface{}, error) {
	logger := GetLogger()
	logger.Info("msg", "Making API POST Request", "url", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("X-CloudKnox-Access-Token", c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("resp", resp, "http_error", err.Error())
		return nil, errors.New("Unable to make HTTP Client Request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Error("msg", "HTTP Response status != 200 OK", "resp", resp.Status, "resource_attributes", "invalid")
		return nil, errors.New("Invalid API Response | Please Check Resource Attributes")
	} else {
		logger.Info("msg", "HTTP Response status == 200 OK", "resp", resp.Status, "resource_attributes", "valid")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	jsonBody := string(body)

	// Create Map from Body of Response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonBody), &responseMap)

	return responseMap, err
}
