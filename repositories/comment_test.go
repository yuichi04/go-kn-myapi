package repositories_test

import (
	"go-kn-myapi/models"
	"go-kn-myapi/repositories"
	"testing"
)

// InsertComment関数のテスト
func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "testest",
	}
	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n",
			expectedCommentID, newComment.CommentID)
	}
	t.Cleanup(func() {
		const sqlStr = `
			DELETE FROM comments
			WHERE article_id = ? AND message = ?
		`
		testDB.Exec(sqlStr, comment.ArticleID, comment.Message)
	})
}

// SelectCommentList関数のテスト
func TestSelectCommentList(t *testing.T) {
	// 1. 期待する結果を定義
	expectedArticleID := 1

	// 2. テスト対象の関数を実行
	got, err := repositories.SelectCommentList(testDB, expectedArticleID)
	if err != nil {
		t.Fatal(err)
	}

	// 3. 期待する結果とテスト結果を比較
	for _, comment := range got {
		if comment.ArticleID != expectedArticleID {
			t.Errorf("want comment of articleID %d but got %d",
				expectedArticleID, comment.ArticleID)
		}
	}
}
