package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ClientParameters struct {
	ServiceAccountID      string
	AccessKey             string
	SecretKey             string
	SharedCredentialsFile string
	Profile               string
}

type Client struct {
	AccessToken string
}

/* Private Variables */
var client *Client
var clientErr error

/* Private Functions */
func buildClient(sai string, ak string, sk string, configurationType string) {
	log := GetLogger()
	log.Info("Using " + configurationType + " to request API Access Token")

	// Make POST Request for API Token

	// Setup HTTP Request
	url := "https://olympus.aws-staging.cloudknox.io/api/v2/service-account/authenticate"
	var jsonStr = []byte(fmt.Sprintf(`{
		"serviceAccountId": "%s",
		"accessKey": "%s",
		"secretKey": "%s"
	  }`, sai, ak, sk))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// Setup Client and Make Request
	hclient := &http.Client{}
	resp, err := hclient.Do(req)
	if err != nil {
		log.Info(err)
		client = nil
		clientErr = errors.New("Client Error")
		return
	}
	defer resp.Body.Close()

	// Get Response
	log.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	jsonBody := string(body)

	// Create Map from Body of Response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonBody), &responseMap)

	if err != nil {
		log.Info(err)
		client = nil
		clientErr = errors.New("Client Error")
		return
	}

	var accessToken = responseMap["accessToken"].(string)

	client = &Client{
		AccessToken: accessToken,
	}

	log.Println("Access Token:", client.AccessToken)

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
