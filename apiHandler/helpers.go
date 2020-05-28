package apiHandler

import (
	"bytes"
	"cloudknox/terraform-provider-cloudknox/common"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func MakePOSTRequest(accessToken string, url string, payload []byte) (map[string]interface{}, error) {
	log := common.GetLogger()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	req.Header.Set("X-CloudKnox-Access-Token", accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Info(err)
		return nil, errors.New("Unable to make HTTP Client Request")
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Info(resp.Status)
		return nil, errors.New("Invalid Response")
	}
	body, _ := ioutil.ReadAll(resp.Body)
	jsonBody := string(body)

	// Create Map from Body of Response
	responseMap := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonBody), &responseMap)

	return responseMap, err
}
