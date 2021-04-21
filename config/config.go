package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"hithub.com/JIeeiroSst/go-app/repository/mysql"
)

type WebConfig struct {
	PORT string 	`envconfig:"WEB_PORT" default:":4040"`
	MysqlConfig mysql.Config `envconfig:"WEB_MYSQL" default:"docker:docker@tcp(localhost:3306)/db?parseTime=True"`
}

var Config WebConfig

func init(){
	err:=envconfig.Process("",&Config)
	if err!=nil{
		panic(err)
	}
	fmt.Println(Config)
}