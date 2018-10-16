package utility

const defaultProtocol = "http"

type Url interface {
	NewUrl(baseUrl string) url
}

type url struct {
	protocol    string
	domain      string
	port        string
	endPoint    string
	pathParams  map[string]string
	queryParams map[string]string
}

func NewUrl(baseUrl string) url {
	return &url{}
}
