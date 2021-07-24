package services

import (
	"encoding/json"
	"hetic-library/services/elasticsearch"
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

func HttpResponseToSearchResults(response *http.Response) (elasticsearch.SearchResults, error) {
	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var searchResults elasticsearch.SearchResults
	err := json.Unmarshal(bodyBytes, &searchResults)
	return searchResults, err
}
