package integration_test

import (
	"testing"	
	"time"
)


var (
	not_found          string = "no found"
	empty              string = "empty"
	create_message     string = "create successfully"
	delete_message     string = "delete successfully"
	update_message     string = "update successfully"
	create_message_err string = "create successfully"
	delete_message_err string = "delete successfully"
	update_message_err string = "update successfully"
	find_all           string = "find successfully"
	find_by_id         string = "find successfully"
)

func TestItemtHttpHandler(t *testing.T) {
	SetupHTTPServer(t)
	time.Sleep(time.Second)

	//author
	t.Run("AuthorById", AuthorById)
	t.Run("AuthorByAll",AuthorByAll)
	t.Run("CreateAuthor",CreateAuthor)
	t.Run("UpdateAuthor",UpdateAuthor)
	t.Run("DeleteAuthor",DeleteAuthor)

	//Article
	t.Run("CreateArticle",CreateArticle)
	t.Run("UpdateArticle1",UpdateArticle1)
	t.Run("DeleteArticle",DeleteArticle)
	t.Run("ArticleAll",ArticleAll)
	t.Run("ArticleById",ArticleById)
}
