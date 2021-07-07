package elasticsearch

const BaseUrl = "http://elasticsearch:9200"
const DefaultIndex = "hetic-library"
const BookType = "book"

func GetUrl(index string, documentType string) string {
	return BaseUrl + "/" + index + "/" + documentType
}