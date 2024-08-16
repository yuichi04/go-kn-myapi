package main

import (
	"go-kn-myapi/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article/1", handlers.GetArticleHandler)
	r.HandleFunc("/article/list", handlers.GetArticleListHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	// http.ListenAndServe関数の第二引数にはサーバの中で使うルータを指定する
	// nilの場合はGoのHTTPサーバがデフォルトで持っているルータが自動的に採用される
	log.Fatal(http.ListenAndServe(":8080", r))
}
