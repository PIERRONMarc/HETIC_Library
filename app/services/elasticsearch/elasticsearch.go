package elasticsearch

const BaseUrl = "http://elasticsearch:9200"
const DefaultIndex = "hetic-library"

type SearchResults struct {
	Hits struct {
		MaxScore float64  `json:"max_score" binding:"required"`
		Results  []Result `json:"hits" binding:"required"`
	} `json:"hits" binding:"required"`
}

type Result struct {
	ID   string                 `json:"_id" binding:"required"`
	Data map[string]interface{} `json:"_source" binding:"required"`
}

// Object used to update a document at esBaseUrl/index/_update/id
type UpdateRequest struct {
	Document interface{} `json:"doc"` // the document to update
}

func GetUrlWithIndex(index string) string {
	return BaseUrl + "/" + index
}
