package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"
	"time"

	"github.com/beego/beego/orm"
)

func Issue(myOrm *orm.Ormer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := utility.AuthToken(w, r)
		if !status {
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var currBook models.Bookings
		err = json.Unmarshal(body, &currBook)
		if err != nil {
			panic(err)
		}
		currBook.Issue_date = time.Now()
		currBook.Status=true
		err = currBook.InsertToDB(myOrm)
		if err != nil {
			utility.Respond(500, "some more error", &w, false)
		}
	}
}
