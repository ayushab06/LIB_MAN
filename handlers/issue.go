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

func Issue(w http.ResponseWriter, r *http.Request) {
	myOrm := orm.NewOrm()
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
	book := models.Books{Id: currBook.Book_id}
	myOrm.Read(&book)
	book.Remaining_stock = book.Remaining_stock - 1
	myOrm.Update(&book, "remaining_stock")
	currBook.Issue_date = time.Now()
	currBook.Status = true
	err = currBook.InsertToDB()
	if err != nil {
		utility.Respond(500, "some more error", &w, false)
	}

}
