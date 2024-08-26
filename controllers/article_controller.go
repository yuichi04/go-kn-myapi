package controllers

import (
	"encoding/json"
	"go-kn-myapi/controllers/services"
	"go-kn-myapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) GetArticleHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query paramter", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "Fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "invalid Query paramter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "Fail to internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusInternalServerError)
		return
	}

	newArticle, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "Fail to internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newArticle)
}

func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusInternalServerError)
		return
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "Fail to internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}
