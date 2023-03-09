package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"
	"time"

	"github.com/beego/beego/orm"
	"github.com/gorilla/sessions"
)

func SignUp(store *sessions.CookieStore, myOrm *orm.Ormer) http.HandlerFunc {
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
		err = u.InsertToDB(myOrm)
		if err != nil {
			utility.Respond(500, "internal server error", &w, false)
		}
		tknStr, err := utility.GenerateJWT("user")
		if err != nil {
			fmt.Println(err.Error())
			utility.Respond(http.StatusInternalServerError, "something wrong at our end", &w, false)
			return
		}
		expirationTime := time.Now().Add(time.Hour)
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tknStr,
			Expires: expirationTime,
		})
	}
}
