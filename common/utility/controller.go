package utility

import (
	"log"
	"net/http"
)

const (
	// OperationCreate is the adjective to describe creating a new record
	OperationCreate = "create"
	// OperationUpdate is the adjective to describe updating an existing record
	OperationUpdate = "update"
	// OperationGet is the adjective to describe getting an existing record
	OperationGet = "get"
	// OperationDelete is the adjective to describe deleting an existing record
	OperationDelete = "delete"
)

// HandleBadRequest logs details of the error encountered and returns an HTTP 400 status
func HandleBadRequest(err error, response http.ResponseWriter) {
	log.Println("Could not decode request:")
	log.Println(err)
	response.WriteHeader(http.StatusBadRequest)
}

// HandleInternalError logs details of the error encountered and returns an HTTP 503 status
func HandleInternalError(entity string, operation string, err error, response http.ResponseWriter) {
	log.Printf("An error occurred trying to %s a %s record:", operation, entity)
	log.Println(err)
	response.WriteHeader(http.StatusServiceUnavailable)
}
