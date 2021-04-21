package integration_test

import (
	"testing"
	"hithub.com/JIeeiroSst/go-app/domain"
	"hithub.com/JIeeiroSst/go-app/delivery/http"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

var (
	config WebConfig = WebConfig{
		ServerAddress:  "localhost",
		ServerPort:     8080,
		ContextTimeout: 20,
		MysqlConfig: MysqlConfig{
			Url:"root:1234@tcp(localhost:3306)/db?parseTime=True",
		},
	}
)

type WebConfig struct {
	ServerAddress  string      `envconfig:"SERVER_ADDRESS"`
	ServerPort     int         `envconfig:"SERVER_PORT"`
	ContextTimeout int64       `envconfig:"CONTEXT_TIMEOUT"`
	MysqlConfig    MysqlConfig `envconfig:"MYSQL"`
}

type MysqlConfig struct {
	Url string 
}

func SetupHTTPServer(t *testing.T) {
	e := echo.New()
	http.NewHandler(e,makeMockRepo(t))
	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()
}

func makeMockRepo(t *testing.T) domain.Repository{
	ctrl := gomock.NewController(t)
	repo:=NewMockRepository(ctrl)

	repo.EXPECT().CreateArticle(gomock.Eq(createArticle)).Return(nil)
	repo.EXPECT().ArticleAll().Return(listArticle,nil)
	repo.EXPECT().ArticleById(gomock.Eq(id)).Return(Article,nil)
	repo.EXPECT().UpdateArticle(gomock.Eq(id),gomock.Eq(UpdateArticle)).Return(nil)
	repo.EXPECT().DeleteArticle(gomock.Eq(id)).Return(nil)

	repo.EXPECT().AuthorById(gomock.Eq(id)).Return(author,nil)
	repo.EXPECT().CreateAuthor(gomock.Eq(createAuthor)).Return(nil)
	repo.EXPECT().AuthorAll().Return(authors,nil)
	repo.EXPECT().UpdateAuthor(gomock.Eq(id),gomock.Eq(updateAuthor)).Return(nil)
	repo.EXPECT().DeleteAuthor(gomock.Eq(id)).Return(nil)

	return repo 
}

var (
	id = 1
	createAuthor= domain.Author{
		ID: 1,
		Name: "chmp",
	}
	updateAuthor= domain.Author{
		Name: "chmp1",
	}
	UpdateArticle=domain.Article{
		Title: "chmp",
		Body: "chmp",
		AuthorId: 1,
	}
	createArticle=domain.Article{
		ID: 1,
		Title: "chmp",
		Body: "chmp",
		AuthorId: 1,
	}
	listArticle=[]domain.Article{
		{
			ID: 1,
			Title: "Name",
			Body: "name",
			AuthorId: 1,
			Author: domain.Author{
				ID: 1,
				Name:"chmp",
		},
	},
	}
	Article=domain.Article{
		ID: 1,
		Title: "Name",
		Body: "name",
		AuthorId: 1,
		Author: domain.Author{
			ID: 1,
			Name:"chmp",
		},
	}
	authors=[]domain.Author{
		{
			ID: 1,
			Name: "chmp",
		},
	}
	author = domain.Author{
		ID: 1,
		Name: "chmp",
	}
)