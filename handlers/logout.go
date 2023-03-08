package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func Logout(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")
		session.Values["authenticated"] = false
		session.Save(r, w)
	}
}
