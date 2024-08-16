package main

import (
	"go-kn-myapi/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.GetArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/list", handlers.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodGet)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	// http.ListenAndServe関数の第二引数にはサーバの中で使うルータを指定する
	// nilの場合はGoのHTTPサーバがデフォルトで持っているルータが自動的に採用される
	log.Fatal(http.ListenAndServe(":8080", r))
}
