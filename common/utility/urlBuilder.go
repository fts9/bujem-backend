package utility

import (
	"regexp"
)

const (
	defaultProtocol      = "http"
	protocolDetectRegex  = ".+:\\/\\/.+"
	trailingSlashRegex   = "^.+\\/$"
	forwardSlash         = "/"
	queryParamStartRegex = "\\?"
	queryParamStart      = "?"
	queryParamAdditional = "&"
)

// URLBuilder provides a builder tool for constructing URLs
type URLBuilder interface {
	Protocol(string) URLBuilder
	Version(string) URLBuilder
	EndPoint(string) URLBuilder
	PathParam(string, string) URLBuilder
	QueryParam(string, string) URLBuilder
	Build() string
}

type urlBuilder struct {
	protocol    string
	domain      string
	version     string
	endPoint    string
	pathParams  map[string]string
	queryParams map[string]string
}

func NewURL(baseURL string) URLBuilder {
	var domainValue string
	if !urlSpecifiesProtocol(baseURL) {
		domainValue = defaultProtocol
	}

	return &urlBuilder{
		protocol:    domainValue,
		domain:      baseURL,
		version:     "",
		endPoint:    "",
		pathParams:  make(map[string]string),
		queryParams: make(map[string]string)}
}

func (builder *urlBuilder) Version(version string) URLBuilder {
	builder.version = version
	return builder
}

func (builder *urlBuilder) Protocol(protocol string) URLBuilder {
	builder.protocol = protocol
	return builder
}

func (builder *urlBuilder) EndPoint(endPoint string) URLBuilder {
	builder.endPoint = endPoint
	return builder
}

func (builder *urlBuilder) PathParam(paramName string, paramValue string) URLBuilder {
	builder.pathParams[paramName] = paramValue
	return builder
}

func (builder *urlBuilder) QueryParam(paramName string, paramValue string) URLBuilder {
	builder.queryParams[paramName] = paramValue
	return builder
}

func (builder *urlBuilder) Build() string {
	var urlValue string
	if len(builder.protocol) > 0 {
		urlValue += builder.protocol
	}

	urlValue += builder.domain
	appendSlash(&urlValue)

	if len(builder.version) > 0 {
		urlValue += builder.version
		appendSlash(&urlValue)
	}

	if len(builder.endPoint) > 0 {
		urlValue += builder.endPoint
	}

	if len(builder.pathParams) > 0 {
		for paramName, paramValue := range builder.pathParams {
			substitutePathParam(&urlValue, paramName, paramValue)
		}
	}

	if len(builder.pathParams) > 0 {
		for paramName, paramValue := range builder.queryParams {
			appendQueryParam(&urlValue, paramName, paramValue)
		}
	}

	return urlValue
}

func appendSlash(url *string) {
	if !hasTrailingSlash(*url) {
		*url += forwardSlash
	}
}

func urlSpecifiesProtocol(url string) bool {
	protocolDetectMatcher := regexp.MustCompile(protocolDetectRegex)
	return protocolDetectMatcher.MatchString(url)
}

func hasTrailingSlash(url string) bool {
	trailingSlashMatcher := regexp.MustCompile(trailingSlashRegex)
	return trailingSlashMatcher.MatchString(url)
}

func substitutePathParam(url *string, paramName string, paramValue string) {
	paramName = "\\{" + paramName + "\\}"
	paramMatcher := regexp.MustCompile(paramName)
	paramMatcher.ReplaceAllString(*url, paramValue)
}

func appendQueryParam(url *string, paramName string, paramValue string) {
	queryParamStartMatcher := regexp.MustCompile(queryParamStartRegex)
	if queryParamStartMatcher.MatchString(*url) {
		*url += queryParamStart
	} else {
		*url += queryParamAdditional
	}
	*url += paramName + "=" + paramValue
}
