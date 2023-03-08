package utility

import (
	"net/http"
)

func RespondWithError(statusCode int, message string, w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusBadRequest)
	jsonData := []byte(`{"success":"NO"}`)
	//append(jsonData,[]byte(`{"message":"there is"}`))
	(*w).Write(jsonData)
}
