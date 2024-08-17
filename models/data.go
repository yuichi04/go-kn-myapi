package models

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "comment1 message.",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "comment2 message.",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID:          1,
		Title:       "article1 title",
		Contents:    "article1 contents.",
		UserName:    "yuichi",
		NiceNum:     100,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt:   time.Now(),
	}

	Article2 = Article{
		ID:        2,
		Title:     "article2 title",
		Contents:  "article2 contents.",
		UserName:  "yuichi",
		NiceNum:   120,
		CreatedAt: time.Now(),
	}
)
