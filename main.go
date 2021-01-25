package main

import (req "demo/RequestHandle"
	"github.com/gin-gonic/gin"
	"demo/middlewarelog"
	"github.com/sirupsen/logrus"

)

func main(){
	route:=gin.New()
	var log = logrus.New()
	route.Use(middlewarelog.Logger(log))
	route.POST("/mail",req.HandleEmailRequest)
	route.Run(":8080")
}

