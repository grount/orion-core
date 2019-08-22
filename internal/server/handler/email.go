package handler

import (
	"net/http"
)

func EmailsHandler(writer http.ResponseWriter, request *http.Request) {
	respondJSON(writer, http.StatusOK, nil)
}
