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

var client *Client
var clientErr error

func credentialsToJSON(credentials *Credentials) []byte {
	c, _ := json.Marshal(credentials)
	return c
}

/* Private Functions */
func buildClient(credentials *Credentials, configurationType string) {
	log := GetLogger()
	log.Info("Using " + configurationType + " to request API Access Token")

	// Make POST Request for API Token

	// Setup HTTP Request

	// Parameters
	url := "https://olympus.aws-staging.cloudknox.io/api/v2/service-account/authenticate"
	var jsonBytes = credentialsToJSON(credentials)

	// Request Configuration
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	// Setup Client and Make Request
	hclient := &http.Client{}
	resp, err := hclient.Do(req)
	if err != nil {
		log.Info(err)
		client = nil
		clientErr = errors.New("Unable to make HTTP Client Request")
		return
	}
	defer resp.Body.Close()

	// Get Response
	log.Println("response Status:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Info(resp.Status)
		client = nil
		clientErr = errors.New("Invalid Credentials")
		log.Info("Please Check Credentials")
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	jsonBody := string(body)

	// Create Map from Body of Response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonBody), &responseMap)

	if err != nil {
		log.Info(err)
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
	return client, clientErr
}

func ValidateClient(client *Client) error {
	if client != nil {
		if client.AccessToken != "" {
			return nil
		}

		return errors.New("No Access Token")
	}

	return errors.New("No Valid Client")
}
