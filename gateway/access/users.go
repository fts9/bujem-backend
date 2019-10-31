package access

import (
	"bujem/common/utility"
	"bytes"
	"encoding/json"
	"io/ioutil"
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

func CreateUser(user map[string]*json.RawMessage) map[string]*json.RawMessage {
	jsonString, err := json.Marshal(user)
	response, err := http.Post(getCreateUserURL(), ContentTypeJSON, bytes.NewBuffer(jsonString))
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	bodyContent, err := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(bodyContent, user)
	if err != nil {
		return nil
	}
	return user
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
