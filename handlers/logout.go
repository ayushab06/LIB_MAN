package handlers

import (
	"lib_man/utility"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	status := utility.AuthToken(w, r)
	if !status {
		return
	}
	cookie := http.Cookie{Name: "token", Value: "", Expires: time.Now().Add(-time.Hour)}
	http.SetCookie(w, &cookie)
	utility.Respond(http.StatusAccepted, "logout was successfull", &w, true)
}
