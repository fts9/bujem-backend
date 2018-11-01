package access

import (
	"bujem/common/utility"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

const (
	version                          = "v1"
	userSingleResourceEndPoint       = "user"
	userSingleResourceWithIDEndPoint = "user/{id}"
	userIDPathParameter              = "id"
)

func CreateUser(user map[string]*json.RawMessage) {

}

func getCreateUserURL() string {
	return getBaseUserURLBuilder().EndPoint(userSingleResourceEndPoint).Build()
}

func getReadUserURL(userID int64) string {
	userIDString := strconv.FormatInt(userID, 10)
	return getBaseUserURLBuilder().EndPoint(userSingleResourceWithIDEndPoint).PathParam(userIDPathParameter, userIDString).Build()
}

func getUpdateUserURL(userID int64) string {
	userIDString := strconv.FormatInt(userID, 10)
	return getBaseUserURLBuilder().EndPoint(userSingleResourceWithIDEndPoint).PathParam(userIDPathParameter, userIDString).Build()
}

func getDeleteUserURL(userID int64) string {
	userIDString := strconv.FormatInt(userID, 10)
	return getBaseUserURLBuilder().EndPoint(userSingleResourceWithIDEndPoint).PathParam(userIDPathParameter, userIDString).Build()
}

func getBaseUserURLBuilder() utility.URLBuilder {
	builder := utility.NewURL(UsersServiceBaseURL)
	builder.Version(version)
	return builder
}

func getHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 10,
	}
}
