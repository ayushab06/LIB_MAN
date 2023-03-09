package utility

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func CheckStatus(store *sessions.CookieStore, w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "cookie-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		Respond(http.StatusForbidden, "not allowed", &w, false)
		return false
	}
	return true

}
