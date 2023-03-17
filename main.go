package main

import (
	"fmt"
	"lib_man/handlers"
	"net/http"
	"os"

	"github.com/beego/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var key []byte

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
	http.HandleFunc("/user/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/search", handlers.Search)
	http.HandleFunc("/issue", handlers.Issue)
	http.HandleFunc("/return", handlers.Return)
	http.HandleFunc("/addbook", handlers.AddBook)
	http.HandleFunc("/logout", handlers.Logout)
	http.ListenAndServe(":8080", nil)
}
