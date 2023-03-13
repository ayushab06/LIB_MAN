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
	Mobile     string `orm:"column(mobile);null"`
}

func init() {
	orm.RegisterModel(new(Users))
}

func (u *Users) InsertToDB() error {
	myOrm := orm.NewOrm()
	_, err := myOrm.Insert(u)
	if err != nil {
		log.Fatal("Error in Insert: ", err)
	}
	return err
}

func GetUserByEmail(email string) ([]Users, error) {
	o := orm.NewOrm()
	var users []Users
	_, err := o.QueryTable("users").Filter("email", email).All(&users)
	return users, err
}

func GetUserByMobile(mobile string) ([]Users, error) {
	o := orm.NewOrm()
	var users []Users
	_, err := o.QueryTable("users").Filter("mobile", mobile).All(&users)
	return users, err
}
