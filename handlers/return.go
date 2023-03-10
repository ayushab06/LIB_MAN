package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"net/http"

	"github.com/beego/beego/orm"
)

func Return(myOrm *orm.Ormer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		o := (*myOrm)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		type returnBook struct {
			User_id int
			Book_id int
		}
		var b returnBook
		err = json.Unmarshal(body, &b)
		book := models.Books{Id: b.Book_id}
		user := models.Users{Id: b.User_id}
		o.Read(&book)
		o.Read(&user)
		book.Remaining_stock = book.Remaining_stock + 1
		bookings := models.GetBooking(b.User_id, b.Book_id, myOrm)
		fmt.Println(bookings.Id)
		bookings.Status = false
		o.Update(&bookings, "status")
		o.Update(&book, "remaining_stock")
		utility.Respond(200, "book return was successfull", &w, true)
	}
}
