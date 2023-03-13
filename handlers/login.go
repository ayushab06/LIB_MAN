package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/utility"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utility.Respond(http.StatusBadRequest, "wrong format", &w, false)
		return
	}
	type loginS struct {
		Password string
	}
	var l loginS
	err = json.Unmarshal(body, &l)
	if err != nil {
		utility.Respond(http.StatusBadRequest, "wrong format", &w, false)
		return
	}
	err = godotenv.Load()
	pass := os.Getenv("LIB_MAN_PASS")
	if l.Password == pass {
		str, _ := utility.GenerateToken()
		cookie := http.Cookie{
			Name:    "token",
			Value:   str,
			Expires: time.Now().Add(time.Hour * 8),
		}
		http.SetCookie(w, &cookie)
		return
	} else {
		utility.Respond(http.StatusBadRequest, "wrong password", &w, false)
		return
	}
}
