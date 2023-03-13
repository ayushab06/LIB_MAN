package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"
)

func SearchCategory(w http.ResponseWriter, r *http.Request) {
	status := utility.AuthToken(w, r)
	if !status {
		return
	}
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
		utility.Respond(http.StatusInternalServerError, "some more error", &w, false)
	}
}
