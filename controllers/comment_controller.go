package controllers

import (
	"encoding/json"
	"go-kn-myapi/controllers/services"
	"go-kn-myapi/models"
	"net/http"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "Fail to decode json\n", http.StatusInternalServerError)
		return
	}

	newComment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "Fail to internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newComment)
}
