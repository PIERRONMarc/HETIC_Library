package repositories

import (
	"bytes"
	"encoding/json"
	"hetic-library/models"
	"hetic-library/services/elasticsearch"
	"io"
	"net/http"
)

func FindBooks(input string) (*http.Response, error) {
	client := &http.Client{}
	var bookQuery models.BookQueryRequest

	var body io.Reader
	if input != "" {
		bookQuery.Query.QueryString.Query = input
		requestBody, err := json.Marshal(bookQuery)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(requestBody)
	} else {
		body = nil
	}

	request, err := http.NewRequest("GET", elasticsearch.GetUrlWithIndex(elasticsearch.DefaultIndex)+"/_search", body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	return response, err
}

func CreateBook(book models.Book) (*http.Response, error) {
	requestBody, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(elasticsearch.GetUrlWithIndex(elasticsearch.DefaultIndex)+"/_doc", "application/json", bytes.NewBuffer(requestBody))
	return response, err
}

func UpdateBook(book models.Book, id string) (*http.Response, error) {
	updateRequest := elasticsearch.UpdateRequest{Document: book}
	requestBody, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(elasticsearch.GetUrlWithIndex(elasticsearch.DefaultIndex)+"/_update/"+id, "application/json", bytes.NewBuffer(requestBody))
	return response, err
}

func DeleteBook(id string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", elasticsearch.GetUrlWithIndex(elasticsearch.DefaultIndex)+"/_doc/"+id, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func DeleteAllBooks() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", elasticsearch.GetUrlWithIndex(elasticsearch.DefaultIndex), nil)
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}
