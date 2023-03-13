package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var b models.Books
	err = json.Unmarshal(body, &b)
	if err != nil {
		panic(err)
	}
	err = b.InsertToDB()
	if err != nil {
		utility.Respond(500, "some more error", &w, false)
	} else {
		utility.Respond(200, "book added successfully", &w, true)
	}
}
