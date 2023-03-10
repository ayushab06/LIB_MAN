package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"

	"github.com/beego/beego/orm"
)

func Register(myOrm *orm.Ormer) http.HandlerFunc {
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
		err = u.InsertToDB(myOrm)
		if err != nil {
			utility.Respond(http.StatusInternalServerError, "something wrong at our end", &w, false)
			return
		}
	}
}
