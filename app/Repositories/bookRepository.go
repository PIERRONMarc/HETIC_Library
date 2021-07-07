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
	response, err := http.Post(elasticsearch.GetUrl(elasticsearch.DefaultIndex, elasticsearch.BookType), "application/json",  bytes.NewBuffer(requestBody))

	return response, err
}