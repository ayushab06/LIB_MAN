package main

import (
	"fmt"
	"lib_man/handlers"

	//"lib_man/models"
	"net/http"
	"os"

	"github.com/beego/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	key   []byte
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	key = []byte(os.Getenv("ENC_KEY"))
	port := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	conn_str := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, port, dbName)
	err = orm.RegisterDataBase("default",
		"mysql",
		conn_str)
	if err != nil {
		fmt.Println("there is some error:", err)
	}
}
func main() {
	myOrm := orm.NewOrm()
	http.HandleFunc("/user/register", handlers.Register(&myOrm))
	http.HandleFunc("/user/login", handlers.Login(&myOrm))
	http.HandleFunc("/user/search/bookname", handlers.SearchBookName(&myOrm))
	http.HandleFunc("/user/search/category", handlers.SearchCategory(&myOrm))
	http.HandleFunc("/user/issue", handlers.Issue(&myOrm))
	http.HandleFunc("/user/return", handlers.Return(&myOrm))
	http.HandleFunc("/admin/addbook", handlers.AddBook(&myOrm))
	http.ListenAndServe(":8080", nil)
}
