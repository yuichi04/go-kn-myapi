package services

import (
	"go-kn-myapi/models"
	"go-kn-myapi/repositories"
)

func (s *MyAppService) PostCommentServices(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
