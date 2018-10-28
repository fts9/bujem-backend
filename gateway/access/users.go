package access

import (
	"bujem/common/utility"
	"encoding/json"
	"fmt"
)

const (
	UsersServiceBaseURL = "http://localhost:9102"
	CreateUserEndPoint  = "users"
)

func CreateUser(user map[string]*json.RawMessage) {
	fmt.Println(getCreateUserURL())
}

func getCreateUserURL() string {
	builder := utility.NewURL(UsersServiceBaseURL)
	return builder.Version("v1").EndPoint(CreateUserEndPoint).Build()
}
