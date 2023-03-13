package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"
)

func SearchBookName(w http.ResponseWriter, r *http.Request) {
	status := utility.AuthToken(w, r)
	if !status {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.Respond(http.StatusBadRequest, "wrong format", &w, false)
		panic(err)
	}
	type search struct {
		Key_word string
	}
	var s search
	err = json.Unmarshal(body, &s)
	if err != nil {
		panic(err)
	}
	books, err := models.GetBookByName(s.Key_word)
	if err != nil {
		utility.Respond(http.StatusNotFound, err.Error(), &w, false)
		return
	}
	utility.RespondBooks(books, &w, true)
	if err != nil {
		fmt.Println(err.Error())
		utility.Respond(500, "some more error", &w, false)
	}
}
