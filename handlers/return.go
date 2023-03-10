package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"

	"github.com/beego/beego/orm"
)

func Return(myOrm *orm.Ormer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var b models.Books
		err = json.Unmarshal(body, &b)
		if err != nil {
			panic(err)
		}
		err = b.InsertToDB(myOrm)
		if err != nil {
			utility.Respond(500, "some more error", &w, false)
		}
	}
}
