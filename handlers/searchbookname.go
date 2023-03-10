package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"

	"github.com/beego/beego/orm"
)

func SearchBookName(myOrm *orm.Ormer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		books, err := models.GetBookByName(s.Key_word, myOrm)
		if err != nil {
			fmt.Println("there was some error with parsing")
		}
		utility.RespondBooks(books, &w, true)
		if err != nil {
			fmt.Println(err.Error())
			utility.Respond(500, "some more error", &w, false)
		}
	}
}
