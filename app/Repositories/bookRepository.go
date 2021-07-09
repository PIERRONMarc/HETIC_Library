package repositories

import (
	"bytes"
	"encoding/json"
	"hetic-library/models"
	"hetic-library/services/elasticsearch"
	"net/http"
)

func CreateBook(book models.Book) (*http.Response, error) {
	requestBody, err := json.Marshal(book)
	response, err := http.Post(elasticsearch.GetUrlWithIndex(elasticsearch.DefaultIndex) + "/_doc", "application/json",  bytes.NewBuffer(requestBody))
	return response, err
}

func UpdateBook(book models.Book, id string) (*http.Response, error) {
	updateRequest := elasticsearch.UpdateRequest{Document: book}
	requestBody, err := json.Marshal(updateRequest)
	response, err := http.Post(elasticsearch.GetUrlWithIndex(elasticsearch.DefaultIndex) + "/_update/" + id, "application/json",  bytes.NewBuffer(requestBody))
	return response, err
}

func DeleteBook(book models.Book, id string) (*http.Response, error) {
	requestBody, err := json.Marshal(book)
	response, err := http.Post(elasticsearch.GetUrlWithIndex(elasticsearch.DefaultIndex) + "/_doc/" + id, "application/json",  bytes.NewBuffer(requestBody))
	return response, err
}

func DeleteAllBooks(book models.Book) (*http.Response, error) {
	requestBody, err := json.Marshal(book)
	response, err := http.Post(elasticsearch.GetUrlWithIndex(elasticsearch.DefaultIndex), "application/json",  bytes.NewBuffer(requestBody))
	return response, err
}
