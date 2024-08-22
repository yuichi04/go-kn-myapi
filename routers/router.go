package routers

import (
	"go-kn-myapi/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(c *controllers.MyAppController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", controllers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", c.PostArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/list", c.GetArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", c.GetArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", c.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/commnet", c.PostCommentHandler).Methods(http.MethodPost)

	return r
}
