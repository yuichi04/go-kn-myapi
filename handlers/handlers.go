package handlers

import (
	"encoding/json"
	"go-kn-myapi/models"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world\n")
}

func GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	err := json.NewDecoder(req.Body).Decode(&reqArticle)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusInternalServerError)
	}
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	err := json.NewDecoder(req.Body).Decode(&reqComment)
	if err != nil {
		http.Error(w, "fail to decode json\n", http.StatusInternalServerError)
	}
	comment := reqComment
	json.NewEncoder(w).Encode(comment)
}
