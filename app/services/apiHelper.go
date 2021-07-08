package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HttpResponseToJson(response *http.Response) (map[string]interface{}, error) {
	defer response.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(response.Body)

	var jsonResponse map[string]interface{}
	err := json.Unmarshal(bodyBytes, &jsonResponse)
	return jsonResponse, err
}