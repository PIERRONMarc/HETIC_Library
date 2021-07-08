package elasticsearch

const BaseUrl = "http://elasticsearch:9200"
const DefaultIndex = "hetic-library"

// Object used to update a document at esBaseUrl/index/_update/id
type UpdateRequest struct {
	Document interface{} `json:"doc"` // the document to update
}

func GetUrlWithIndex(index string) string {
	return BaseUrl + "/" + index 
}