package api

import (
	"database/sql"
	"go-kn-myapi/controllers"
	"go-kn-myapi/services"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	s := services.NewMyAppService(db)
	ac := controllers.NewArticleController(s)
	cc := controllers.NewCommentController(s)
	r := mux.NewRouter()
	r.HandleFunc("/article", ac.PostArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/list", ac.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", ac.GetArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", ac.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/commnet", cc.PostCommentHandler).Methods(http.MethodPost)

	return r
}
