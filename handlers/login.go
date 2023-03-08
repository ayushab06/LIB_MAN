package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"

	"github.com/beego/beego/orm"
	"github.com/gorilla/sessions"
)

func Login(store *sessions.CookieStore, myOrm *orm.Ormer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var u models.Users
		err = json.Unmarshal(body, &u)
		if err != nil {
			utility.RespondWithError(500, "thet", &w)
		}
		err = u.InsertToDB(myOrm)
		if err != nil {
			utility.RespondWithError(500, "some more error", &w)
		}
		
		session, _ := store.Get(r, "cookie-name")
		session.Values["authenticated"] = true
		session.Save(r, w)
	}
}
