package models

import (
	"log"
	"time"

	"github.com/beego/beego/orm"
)

func init() {
	orm.RegisterModel(new(Bookings))
}

type Bookings struct {
	Id         int       `orm:"column(id);auto"`
	Book_id    int       `orm:"column(book_id);null"`
	User_id    int       `orm:"column(user_id);null"`
	Issue_date time.Time `orm:"column(issue_date);null"`
	Status     bool      `orm:"column(status);null"`
}

func (booking *Bookings) InsertToDB(myOrm *orm.Ormer) error {
	_, err := (*myOrm).Insert(booking)
	if err != nil {
		log.Fatal("Error in Insert: ", err)
	}
	return err
}
func GetBooking(u_id int, b_id int, myOrm *orm.Ormer) (b Bookings) {
	(*myOrm).QueryTable(new(Bookings)).Filter("user_id", u_id).Filter("book_id", b_id).Filter("status", 1).One(&b)
	return b
}
