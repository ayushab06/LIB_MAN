package models

import (
	"log"

	"github.com/beego/beego/orm"
)

type Books struct {
	Id              int    `orm:"column(id);auto"`
	Book_name       string `orm:"column(book_name);null"`
	Author_name     string `orm:"column(author_name);null"`
	Category_name   string `orm:"column(category_name);null"`
	Details         string `orm:"column(details);null"`
	Remaining_stock int    `orm:"column(remaining_stock);null"`
	Price           int    `orm:"column(price);null"`
}

func init() {
	orm.RegisterModel(new(Books))
}

func (b *Books) InsertToDB(myOrm *orm.Ormer) error {
	_, err := (*myOrm).Insert(b)
	if err != nil {
		log.Fatal("Error in Insert: ", err)
	}
	return err
}
