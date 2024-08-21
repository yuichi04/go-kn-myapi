package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"go-kn-myapi/controllers"
	"go-kn-myapi/services"

	"github.com/gorilla/mux"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Fail to connect DB")
		return
	}
	s := services.NewMyAppService(db)
	c := controllers.NewMyAppController(s)

	r := mux.NewRouter()

	r.HandleFunc("/hello", controllers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", c.GetArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/list", c.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", c.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/nice", c.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", c.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	// http.ListenAndServe関数の第二引数にはサーバの中で使うルータを指定する
	// nilの場合はGoのHTTPサーバがデフォルトで持っているルータが自動的に採用される
	log.Fatal(http.ListenAndServe(":8080", r))
}
