package main

import (
	"github.com/labstack/echo/v4"
	repo "hithub.com/JIeeiroSst/go-app/repository/mysql"
	"hithub.com/JIeeiroSst/go-app/config"
	"hithub.com/JIeeiroSst/go-app/delivery/http"
)

func main(){
	e:=echo.New()
    repo:=repo.NewMysqlConn(&config.Config.MysqlConfig)
	http.NewHandler(e,repo)
	e.Logger.Fatal(e.Start(":8080"))
}