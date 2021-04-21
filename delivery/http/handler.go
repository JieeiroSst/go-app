package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"hithub.com/JIeeiroSst/go-app/domain"
	"hithub.com/JIeeiroSst/go-app/log"
	"hithub.com/JIeeiroSst/go-app/utils"
)

var (
	not_found      		string = "no found"
	empty          		string = "empty"
	create_message 		string = "create successfully"
	delete_message 		string = "delete successfully"
	update_message 		string = "update successfully"
	create_message_err 	string = "create successfully"
	delete_message_err 	string = "delete successfully"
	update_message_err 	string = "update successfully"
	find_all			string = "find successfully"
	find_by_id			string = "find successfully"
)

type Handler struct{
	repo domain.Repository
}
//mockgen -source=api.go -destination=./integration_test/mocks_test.go -package=integration_test
func NewHandler(e *echo.Echo,handler domain.Repository) {
	h:=Handler{
		repo: handler,
	}
	authorGroup:=e.Group("/author")
	authorGroup.POST("",h.CreateAuthor)
	authorGroup.PUT("/:id",h.UpdateAuthor)
	authorGroup.DELETE("/:id",h.DeleteAuthor)
	authorGroup.GET("/",h.AuthorAll)
	authorGroup.GET("/:id",h.AuthorById)

	articleGroup:=e.Group("/article")
	articleGroup.POST("",h.CreateArticle)
	articleGroup.PUT("/:id",h.UpdateArticle)
	articleGroup.DELETE("/:id",h.DeleteArticle)
	articleGroup.GET("/",h.ArticleAll)
	articleGroup.GET("/:id",h.ArticleById)
}

func (h *Handler) CreateArticle(e echo.Context) error {
	article:=domain.Article{
		ID: utils.StringToInt(e.FormValue("id")),
		Title: e.FormValue("title"),
		Body: e.FormValue("body"),
		AuthorId: utils.StringToInt(e.FormValue("author_id")),
	}	
	err:=h.repo.CreateArticle(article)
	if err!=nil {
		log.InitZapLog().Sugar().Errorf(create_message_err+"%s",err)
		return e.String(http.StatusInternalServerError,create_message_err)
	}
	log.InitZapLog().Info(create_message)
	return e.String(http.StatusOK,create_message)
}

func (h *Handler) ArticleAll(e echo.Context) error {
	aritcles,err:=h.repo.ArticleAll()
	if err!=nil{
		log.InitZapLog().Sugar().Errorf("server find all error %s",err)
		return e.String(http.StatusInternalServerError,not_found)
	}
	log.InitZapLog().Info(find_all)
	return e.JSON(http.StatusOK,aritcles)
}

func (h *Handler) ArticleById(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	article,err:=h.repo.ArticleById(id)
	if err!=nil{
		log.InitZapLog().Sugar().Errorf(empty+"%s",err)
		return e.String(http.StatusInternalServerError,empty)
	}
	log.InitZapLog().Info(find_all)
	return e.JSON(http.StatusOK,article)
}

func (h *Handler) UpdateArticle(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	article:=domain.Article{
		Title: e.FormValue("title"),
		Body: e.FormValue("body"),
		AuthorId: utils.StringToInt(e.FormValue("author_id")),
	}
	err:=h.repo.UpdateArticle(id,article)
	if err!=nil{
		log.InitZapLog().Sugar().Errorf(update_message_err+"%s",err)
		return e.String(http.StatusInternalServerError,update_message_err)
	}
	log.InitZapLog().Info(update_message)
	return e.String(http.StatusOK,update_message)
}

func (h *Handler) DeleteArticle(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	err:=h.repo.DeleteArticle(id)
	if err!=nil{
		log.InitZapLog().Sugar().Errorf(delete_message_err+"%s",err)
		return e.String(http.StatusInternalServerError,delete_message_err)
	}
	log.InitZapLog().Info(delete_message)
	return e.String(http.StatusOK,delete_message)
}

func (h *Handler) CreateAuthor(e echo.Context) error{
	author:=domain.Author{
		ID: utils.StringToInt(e.FormValue("id")),
		Name: e.FormValue("name"),
	}
	err:=h.repo.CreateAuthor(author)
	if err!=nil {
		log.InitZapLog().Sugar().Errorf(create_message_err+"%s",err)
		return e.String(http.StatusInternalServerError,create_message_err)
	}
	log.InitZapLog().Info(create_message)
	return e.String(http.StatusOK,create_message)
}

func (h *Handler) AuthorAll(e echo.Context) error {
	authors,err:=h.repo.AuthorAll()
	if err!=nil{
		log.InitZapLog().Sugar().Errorf(empty+"%s",err)
		return e.String(http.StatusInternalServerError,empty)
	}
	log.InitZapLog().Info(find_all)
	return e.JSON(http.StatusOK,authors)
}

func (h *Handler) AuthorById(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
    author,err:=h.repo.AuthorById(id)
	if err!=nil {
		log.InitZapLog().Sugar().Errorf(empty+"%s",err)
		return e.String(http.StatusInternalServerError,empty)
	}
	log.InitZapLog().Info(find_by_id)
	return e.JSON(http.StatusOK,author)
}

func (h *Handler) UpdateAuthor(e echo.Context) error{
	id:=utils.StringToInt(e.Param("id"))
    author:=domain.Author{
		Name: e.FormValue("name"),
	}
	err:=h.repo.UpdateAuthor(id,author)
	if err!=nil {
		log.InitZapLog().Sugar().Errorf(update_message_err+"%s",err)
		return e.String(http.StatusInternalServerError,update_message_err)
	}
	log.InitZapLog().Info(update_message)
	return e.String(http.StatusOK,update_message)
}

func (h *Handler) DeleteAuthor(e echo.Context) error {
	id:=utils.StringToInt(e.Param("id"))
	err:=h.repo.DeleteAuthor(id)
	if err!=nil {
		log.InitZapLog().Sugar().Errorf(delete_message_err+"%s",err)
		return e.String(http.StatusInternalServerError,delete_message_err)
	}
	log.InitZapLog().Info(delete_message)
	return e.String(http.StatusOK,delete_message)
}