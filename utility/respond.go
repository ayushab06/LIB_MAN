package utility

import (
	"encoding/json"
	"fmt"
	"lib_man/models"
	"net/http"
)

func Respond(statusCode int, message string, w *http.ResponseWriter, success bool) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusBadRequest)
	res := models.Response{Success: success, Message: message}
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println("there was some error marshalling the data")
	}
	(*w).Write(data)
}

func RespondBooks(books []models.Books, w *http.ResponseWriter, success bool) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusAccepted)
	res := models.BookResponse{Success: success, Books: books}
	data, err := json.Marshal(res)
	if err != nil {
		fmt.Println("there was some error marshalling the data")
	}
	(*w).Write(data)
}