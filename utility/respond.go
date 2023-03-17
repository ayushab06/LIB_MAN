package utility

import (
	"encoding/json"
	"lib_man/models"
	"net/http"
)

func Respond(statusCode int, message string, w *http.ResponseWriter, success bool) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(statusCode)
	res := models.Response{Success: success, Message: message}
	data, _ := json.Marshal(res)
	(*w).Write(data)
}

func RespondStruct(data []byte, w *http.ResponseWriter, success bool) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusAccepted)
	(*w).Write(data)
}

func RespondBooks(books []models.Books, w *http.ResponseWriter, success bool) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusAccepted)
	res := models.BookResponse{Success: success, Books: books}
	data, err := json.Marshal(res)
	if err != nil {
		Respond(http.StatusBadRequest, "Unexpected data format", w, false)
		return
	}
	(*w).Write(data)
}
