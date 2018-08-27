package controller

import (
	"bujem/users/application"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Listen() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", HandleGetUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func HandleGetUser(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])
	user, _ := application.GetUser(id)
	json.NewEncoder(response).Encode(&user)
}
