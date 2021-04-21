package utils

import (
	"strconv"

	"hithub.com/JIeeiroSst/go-app/log"
)

func StringToInt(num string) int {
	numInt,err:=strconv.Atoi(num)
	if err!=nil {
		log.InitZapLog().Sugar().Errorf("string to int error %s",err)
	}
	log.InitZapLog().Info("string to int sucess")
	return numInt
}