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
	// articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	// if err != nil {
	// 	http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	// 	return
	// }
	// resString := fmt.Sprintf("Article No.%d\n", articleID)
	// io.WriteString(w, resString)
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// queryMap := req.URL.Query()

	// var page int
	// if p, ok := queryMap["page"]; ok && len(p) > 0 {
	// 	var err error
	// 	page, err = strconv.Atoi(p[0])
	// 	if err != nil {
	// 		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	// 		return
	// 	}
	// } else {
	// 	page = 1
	// }

	// resString := fmt.Sprintf("Article List (page %d)\n", page)
	// io.WriteString(w, resString)
	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		http.Error(w, "fait to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "Posting Nice...\n")
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// io.WriteString(w, "Posting Comment...")
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
