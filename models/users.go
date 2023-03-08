package models

import (
	"log"

	"github.com/beego/beego/orm"
)

type Users struct {
	Id         int    `orm:"column(id);auto"`
	First_name string `orm:"column(first_name);null"`
	Last_name  string `orm:"column(last_name);null"`
	Email      string `orm:"column(email);null"`
	Password   string `orm:"column(password);null"`
}

func init() {
	orm.RegisterModel(new(Users))
}

func (u *Users) InsertToDB(myOrm *orm.Ormer) error {
	_, err := (*myOrm).Insert(u)
	if err != nil {
		log.Fatal("Error in Insert: ", err)
	}
	return err
}
