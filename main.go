package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world\n")
}

func postArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func postNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func getArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Artilce No.1")
}

func getArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List")
}

func postCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article/1", getArticleHandler)
	http.HandleFunc("/article/list", getArticleListHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
