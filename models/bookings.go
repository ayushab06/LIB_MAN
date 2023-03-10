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
