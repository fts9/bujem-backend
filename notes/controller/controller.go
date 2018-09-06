package controller

import (
	"bujem/notes/application"
	"bujem/notes/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Listen intialises the router and begins listening for requests on the designated port
func Listen() {
	log.Println("Configuring router")
	router := mux.NewRouter()

	log.Println("Mapping end-point [GET]/note/{id}")
	router.HandleFunc("/note/{id}", HandleFindNoteByID).Methods("GET")

	log.Println("Mapping end-point [POST]/note/{id}")
	router.HandleFunc("/note", HandleCreateNote).Methods("POST")

	log.Println("Mapping end-point [PUT]/note")
	router.HandleFunc("/note", HandleUpdateNote).Methods("PUT")

	log.Println("Mapping end-point [DELETE]/note/{id}")
	router.HandleFunc("/note/{id}", HandleDeleteNoteByID).Methods("DELETE")

	log.Println("Listening on port 9101")
	log.Fatal(http.ListenAndServe(":9101", router))
}

// HandleCreateNote decodes the create request and passes it on to the application layer
func HandleCreateNote(response http.ResponseWriter, request *http.Request) {
	var note model.Note
	err := json.NewDecoder(request.Body).Decode(&note)
	if err != nil {
		handleBadRequest(err, response)
		return
	}

	err = application.Create(&note)
	if err != nil {
		handleInternalError("create", err, response)
		return
	}

	json.NewEncoder(response).Encode(&note)
}

// HandleUpdateNote decodes the update request and passes it on to the application layer
func HandleUpdateNote(response http.ResponseWriter, request *http.Request) {
	var note model.Note
	err := json.NewDecoder(request.Body).Decode(&note)
	if err != nil {
		handleBadRequest(err, response)
		return
	}

	err = application.Update(&note)
	if err != nil {
		handleInternalError("update", err, response)
		return
	}

	json.NewEncoder(response).Encode(&note)
}

// HandleFindNoteByID decodes the find request and passes it on to the application layer
func HandleFindNoteByID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		handleBadRequest(err, response)
		return
	}

	note, err := application.FindByID(id)
	if err != nil {
		handleInternalError("get", err, response)
		return
	}

	json.NewEncoder(response).Encode(&note)
}

// HandleDeleteNoteByID decodes the delete request and passes it on to the application layer
func HandleDeleteNoteByID(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		handleBadRequest(err, response)
		return
	}

	err = application.DeleteByID(id)
	if err != nil {
		handleInternalError("delete", err, response)
		return
	}
	response.WriteHeader(http.StatusNoContent)
}

func handleBadRequest(err error, response http.ResponseWriter) {
	log.Println("Could not decode request:")
	log.Println(err)
	response.WriteHeader(http.StatusBadRequest)
}

func handleInternalError(operation string, err error, response http.ResponseWriter) {
	log.Printf("An error occurred trying to %s a Note record:", operation)
	log.Println(err)
	response.WriteHeader(http.StatusServiceUnavailable)
}
