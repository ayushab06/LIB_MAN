package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	status := utility.AuthToken(w, r)
	if !status {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.Respond(http.StatusBadRequest, "wrong format", &w, false)
		return
	}
	var b models.Books
	err = json.Unmarshal(body, &b)
	if err != nil {
		utility.Respond(http.StatusBadRequest, "wrong format", &w, false)
		return
	}
	alr := models.GetBookByExactName(b.Book_name)
	if alr.Book_name != "" {
		alr.Remaining_stock += b.Remaining_stock
		err = alr.Update()
	} else {
		err = b.InsertToDB()
	}
	if err != nil {
		utility.Respond(http.StatusInternalServerError, "some error at our end", &w, false)
	} else {
		utility.Respond(http.StatusAccepted, "book added successfully", &w, true)
	}
}
