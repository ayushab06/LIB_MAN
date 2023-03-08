package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

func Secret(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "cookie-name")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Print secret message
		fmt.Fprintln(w, "The cake is a lie!")
	}
}
