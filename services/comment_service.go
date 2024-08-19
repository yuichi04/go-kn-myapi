package services

import (
	"fmt"
	"go-kn-myapi/models"
	"go-kn-myapi/repositories"
)

func PostCommentServices(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err)
		return models.Comment{}, err
	}
	defer db.Close()

	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
