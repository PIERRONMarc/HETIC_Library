package elasticsearch

const BaseUrl = "http://elasticsearch:9200"
const DefaultIndex = "hetic-library"

func GetUrlWithIndex(index string) string {
	return BaseUrl + "/" + index 
}