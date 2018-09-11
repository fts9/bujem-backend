package controller

import (
	"bujem/common/utility"
	"bujem/users/application"
	"bujem/users/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const entity = "User"

// Listen configures the routings and begins listening on its designated port
func Listen() {
	log.Println("Configuring routings for Users controller")
	router := mux.NewRouter()
	log.Println("Mapping end-point [GET]/user/{id}")
	router.HandleFunc("/user/{id}", HandleGetUser).Methods("GET")

	log.Println("Mapping end-point [POST]/user")
	router.HandleFunc("/user", HandleCreateUser).Methods("POST")

	log.Println("Mapping end-point [PUT]/user")
	router.HandleFunc("/user", HandleUpdateUser).Methods("PUT")

	log.Println("Mapping end-point [DELETE]/user/{id}")
	router.HandleFunc("/user/{id}", HandleDeleteUser).Methods("DELETE")

	log.Println("Listening on localhost port 9102")
	log.Fatal(http.ListenAndServe(":9102", router))
}

// HandleCreateUser accepts and decodes a create user request and passes it to the application service
func HandleCreateUser(response http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		utility.HandleBadRequest(err, response)
		return
	}

	user, err = application.CreateUser(&user)
	if err != nil {
		utility.HandleInternalError(entity, utility.OperationCreate, err, response)
		return
	}

	json.NewEncoder(response).Encode(&user)
}

// HandleUpdateUser accepts and decodes an update user request and passes it to the application service
func HandleUpdateUser(response http.ResponseWriter, request *http.Request) {
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		utility.HandleBadRequest(err, response)
		return
	}

	user, err = application.UpdateUser(&user)
	if err != nil {
		utility.HandleInternalError(entity, utility.OperationUpdate, err, response)
		return
	}

	json.NewEncoder(response).Encode(&user)
}

// HandleGetUser accepts a get user request and passes it to the application service
func HandleGetUser(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utility.HandleBadRequest(err, response)
		return
	}

	user, err := application.GetUser(id)
	if err != nil {
		utility.HandleInternalError(entity, utility.OperationGet, err, response)
		return
	}

	json.NewEncoder(response).Encode(&user)
}

// HandleDeleteUser accepts a delete user request and passes it to the application service
func HandleDeleteUser(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utility.HandleBadRequest(err, response)
		return
	}

	err = application.DeleteUser(id)
	if err != nil {
		utility.HandleInternalError(entity, utility.OperationDelete, err, response)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}
