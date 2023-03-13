package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"
)

func Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var u models.Users
		err = json.Unmarshal(body, &u)
		if err != nil {
			panic(err)
		}
		users, _ := models.GetUserByEmail(u.Email)
		if len(users) != 0 {
			utility.Respond(http.StatusBadRequest, "the email already exists", &w, false)
			return
		}
		users, _ = models.GetUserByMobile(u.Mobile)
		if len(users) != 0 {
			utility.Respond(http.StatusBadRequest, "the mobile already exists", &w, false)
			return
		}
		err = u.InsertToDB()
		if err != nil {
			utility.Respond(http.StatusInternalServerError, "something wrong at our end", &w, false)
			return
		}
	}
}
