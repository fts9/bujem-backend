package controller

import (
	"bujem/users/application"
	"bujem/users/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Listen() {
	log.Println("Initializing application service")
	application.Initialize()

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

	log.Println("Listening on localhost port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HandleCreateUser(response http.ResponseWriter, request *http.Request) {
	var user model.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	_ = application.CreateUser(&user)
	json.NewEncoder(response).Encode(&user)
}

func HandleUpdateUser(response http.ResponseWriter, request *http.Request) {
	var user model.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	_ = application.UpdateUser(&user)
	json.NewEncoder(response).Encode(&user)
}

func HandleGetUser(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	user, _ := application.GetUser(id)
	json.NewEncoder(response).Encode(&user)
}

func HandleDeleteUser(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	_ = application.DeleteUser(id)
}
