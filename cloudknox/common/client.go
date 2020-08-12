package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	config "github.com/go-akka/configuration"
	"github.com/mitchellh/go-homedir"
)

/* Private Functions */
func credentialsToJSON(credentials *Credentials) []byte {
	c, _ := json.Marshal(credentials)
	return c
}

func createNewRequest(method, url string, body io.Reader, client *Client) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	if client.AccessToken != "" {
		req.Header.Add("X-CloudKnox-Access-Token", client.AccessToken)
	}
	if client.AccessKey != "" {
		req.Header.Add("X-CloudKnox-Access-Key", client.AccessKey)
	}
	if client.ServiceAccountID != "" {
		req.Header.Add("X-CloudKnox-Service-Account-Id", client.ServiceAccountID)
	}

	req.Header.Add("X-CloudKnox-Timestamp", time.Now().Format(time.RFC3339))

	req.Header.Add("User-Agent", "CloudKnoxTerraformProvider/1.0.0")
	return req, nil
}

func getBaseURLFromConfig(path string) (*url.URL, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	text := string(content)
	conf := config.ParseString(text)
	return url.Parse(conf.GetString("api.base_url"))
}

func (c *Client) getRelativeURL(urlPath string) string {
	relativeURL, _ := url.Parse(urlPath)
	return c.BaseURL.ResolveReference(relativeURL).String()
}

/* Public Functions */

//POST allows clients to make HTTP POST Requests
func (c *Client) POST(route string, payload []byte) (map[string]interface{}, error) {
	logger := GetLogger()
	postURL := c.getRelativeURL(route)
	logger.Debug("msg", "making API POST request", "url", postURL)

	req, err := createNewRequest(http.MethodPost, postURL, bytes.NewBuffer(payload), c)
	if err != nil {
		logger.Error("Failed To Create HTTP Request", "http_error", err.Error())
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logger.Error("resp", resp, "http_error", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Error("msg", "HTTP Response status != 200 OK", "resp", resp.Status, "resource_attributes", "invalid")
		return nil, fmt.Errorf("Invalid API Response | Please Check Resource Attributes")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	response := make(map[string]interface{})
	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// NewClient creates a CloudKnox API Client used to interface with the API
func NewClient(credentials *Credentials) (*Client, error) {
	if credentials == nil {
		return nil, fmt.Errorf("credentials not found")
	}
	logger := GetLogger()
	logger.Info("msg", "building CloudKnox client object")
	homeDir, _ := homedir.Dir()
	apiConfigurationPath := homeDir + "//.cloudknox//api.conf"
	baseURL, err := getBaseURLFromConfig(apiConfigurationPath)

	if err != nil {
		return nil, err
	}

	client := &Client{
		BaseURL:    baseURL,
		httpClient: http.DefaultClient,
	}

	response, err := client.POST("/api/v2/service-account/authenticate", credentialsToJSON(credentials))
	if err != nil {
		logger.Error("msg", "failed to read http response", "unmarshal_error", err)
		return nil, err
	}
	client.AccessToken = response["accessToken"].(string)
	client.AccessKey = response["accessKey"].(string)
	client.ServiceAccountID = credentials.ServiceAccountID

	return client, nil
}
