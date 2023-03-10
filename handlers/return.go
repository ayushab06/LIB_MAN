package handlers

import (
	"encoding/json"
	"io/ioutil"
	"lib_man/models"
	"lib_man/utility"
	"math"
	"net/http"
	"time"

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
		t := time.Now().Sub(bookings.Issue_date)
		bookings.Status = false
		o.Update(&bookings, "status")
		o.Update(&book, "remaining_stock")
		type res struct {
			Success  bool
			Message  string
			Payment  int
			Duration int
		}
		days := int(math.Ceil(t.Hours() / float64(24)))
		fees := days
		if days > 30 {
			fees += (days - 30)
		}
		rp := res{
			Success:  true,
			Message:  "Successfull",
			Payment:  fees,
			Duration: days,
		}
		data, _ := json.Marshal(rp)
		utility.RespondStruct(data, &w, true)
	}
}
