package repositories_test

import (
	"testing"

	"go-kn-myapi/models"
	"go-kn-myapi/repositories"
	"go-kn-myapi/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

// SelectArticleDetail関数のテスト
func TestSelectArticleDetail(t *testing.T) {
	// 1. 期待する結果を定義
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			// 2. テスト対象の関数を実行
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			// 3. 期待する結果と関数の実行結果を比較
			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Contents: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {
	// 1. 期待する結果を定義
	expectedNum := len(testdata.ArticleTestData)

	// 2. テスト対象の関数を実行
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	// 3. 期待する結果と関数の実行結果を比較
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

// InsertArticle関数のテスト
func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}

	expectedArticleID := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.ID != expectedArticleID {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleID, newArticle.ID)
	}

	// TestInsertArticle だけの個別の後処理
	t.Cleanup(func() {
		const sqlStr = `
			DELETE FROM articles
			WHERE title = ? AND contents = ? AND username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

// UpdateNiceNum関数のテスト
func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	row, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}
	expectedNum := row.NiceNum + 1

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	if expectedNum != got.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expectedNum)
	}
}
