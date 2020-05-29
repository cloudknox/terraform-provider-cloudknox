package apiHandler

import (
	"bytes"
	"cloudknox/terraform-provider-cloudknox/common"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/log/level"
)

func MakePOSTRequest(accessToken string, url string, payload []byte) (map[string]interface{}, error) {
	logger := common.GetLogger()
	level.Info(logger).Log("msg", "Making API POST Request", "url", url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("X-CloudKnox-Access-Token", accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		level.Error(logger).Log("resp", resp, "http_error", err.Error())
		return nil, errors.New("Unable to make HTTP Client Request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		level.Error(logger).Log("msg", "HTTP Response status != 200 OK", "resp", resp.Status, "resource_attributes", "invalid")
		return nil, errors.New("Invalid API Response | Please Check Resource Attributes")
	} else {
		level.Info(logger).Log("msg", "HTTP Response status == 200 OK", "resp", resp.Status, "resource_attributes", "valid")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	jsonBody := string(body)

	// Create Map from Body of Response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonBody), &responseMap)

	return responseMap, err
}
