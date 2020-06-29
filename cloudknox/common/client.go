package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"terraform-provider-cloudknox/cloudknox/utils"
)

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

	url := GetConfiguration().BaseURL + GetConfiguration().Routes.Auth

	// Request Configuration
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
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

	// dump, _ := httputil.DumpRequest(req, true)

	// logger.Info("auth dump", dump)

	// Get Response
	logger.Info("msg", "Got HTTP Response")
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		logger.Error("msg", "HTTP Response status != 200 OK", "resp", resp.Status, "credentials", "invalid")
		client = nil
		clientErr = fmt.Errorf("Error During Authentication, Server Responded With %s", resp.Status)
		return
	} else {
		logger.Info("msg", "HTTP Response status == 200 OK", "resp", resp.Status, "credentials", "valid")
	}
	jsonBody := string(body)

	// Create Map from Body of Response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonBody), &responseMap)

	if err != nil {
		logger.Error("msg", "Unable to extract response from body", "unmarshal_error", err)
		logger.Error("body", body)
		client = nil
		clientErr = errors.New("Unable to read HTTP Response")
		return
	}

	var accessToken = responseMap["accessToken"].(string)

	client = &Client{
		AccessToken: accessToken,
	}
	clientErr = nil

	// logger.Debug("access_token", accessToken)

	return
}

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

	req_dump, _ := httputil.DumpRequest(req, true)
	logger.Debug("request_dump", utils.Truncate(string(req_dump), 30, true))

	client := &http.Client{}
	resp, err := client.Do(req)

	response_dump, _ := httputil.DumpResponse(resp, true)
	logger.Debug("response_dump", utils.Truncate(string(response_dump), 30, true))

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

	if err != nil {
		err = fmt.Errorf(string(body))
	}

	return responseMap, err
}
