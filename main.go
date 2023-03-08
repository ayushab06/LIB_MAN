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

	"github.com/gorilla/sessions"
)

var (
	key   []byte
	store = sessions.NewCookieStore(key)
)

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	key = []byte(os.Getenv("ENC_KEY"))
	//fmt.Println(key)
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
	http.HandleFunc("/secret", handlers.Secret(store))
	http.HandleFunc("/user/signup", handlers.SignUp(store, &myOrm))
	http.HandleFunc("/user/login", handlers.Login(store, &myOrm))
	http.HandleFunc("/admin/addbook", handlers.AddBook(store, &myOrm))

	http.HandleFunc("/logout", handlers.Logout(store))
	http.ListenAndServe(":8080", nil)
}
