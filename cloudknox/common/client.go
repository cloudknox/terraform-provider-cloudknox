package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	config "github.com/go-akka/configuration"
	"github.com/mitchellh/go-homedir"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func credentialsToJSON(credentials *Credentials) []byte {
	c, _ := json.Marshal(credentials)
	return c
}

func createNewRequest(method, url string, body io.Reader, accessToken string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if accessToken != "" {
		req.Header.Add("X-CloudKnox-Access-Token", accessToken)
	}
	req.Header.Add("User-Agent", "CloudKnoxTerraformProvider/1.0.0")
	return req, nil
}

func getBaseUrlFromConfig(path string) (*url.URL, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	text := string(content)
	conf := config.ParseString(text)
	return url.Parse(conf.GetString("api.base_url"))
}

func (c *Client) getRelativeUrl(urlPath string) string {
	relativeURL, _ := url.Parse(urlPath)
	return c.BaseURL.ResolveReference(relativeURL).String()
}

func (c *Client) POST(route string, payload []byte) (map[string]interface{}, error) {
	logger := GetLogger()
	postUrl := c.getRelativeUrl(route)
	logger.Debug("msg", "making API POST request", "url", postUrl)
	req, err := createNewRequest(
		http.MethodPost, postUrl, bytes.NewBuffer(payload), c.AccessToken,
	)
	if err != nil {
		logger.Error("Failed To Create Http Request", "http_error", err.Error())
		return nil, errors.New("Unable to make HTTP Client Request")
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logger.Error("resp", resp, "http_error", err.Error())
		return nil, errors.New("Unable to make HTTP Client Request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		logger.Error("msg", "HTTP Response status != 200 OK", "resp", resp.Status, "resource_attributes", "invalid")
		return nil, errors.New("Invalid API Response | Please Check Resource Attributes")
	}

	body, _ := ioutil.ReadAll(resp.Body)
	response := make(map[string]interface{})
	err = json.Unmarshal([]byte(string(body)), &response)
	if err != nil {
		err = fmt.Errorf(string(body))
	}
	return response, err
}

func NewClient(credentials *Credentials)(*Client, error){
	if credentials == nil {
		return nil, fmt.Errorf("credentials not found")
	}
	logger := GetLogger()
	logger.Info("msg", "building CloudKnox client object", "config_type")
	homeDir, _ := homedir.Dir()
	apiConfigurationPath := homeDir + "//.cloudknox//api.conf"
	baseURL, err := getBaseUrlFromConfig(apiConfigurationPath)

	if err != nil {
		return nil, err
	}

	client := &Client{
		BaseURL: baseURL,
		httpClient:  http.DefaultClient,
	}

	response, err := client.POST("api/v2/service-account/authenticate", credentialsToJSON(credentials))
	if err != nil {
		logger.Error("msg", "failed to read http response", "unmarshal_error", err)
		return nil, err
	}
	client.AccessToken = response["accessToken"].(string)
	return client, nil
}