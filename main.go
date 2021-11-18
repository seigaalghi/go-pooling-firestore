package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seigaalghi/go-pooling-firestore/controller"
	"github.com/seigaalghi/go-pooling-firestore/database"
	"github.com/seigaalghi/go-pooling-firestore/middlewares"
)

func main() {
	server := gin.Default()
	v1 := server.Group("api/v1")
	v2 := server.Group("api/v2")

	con1, err := database.Connect("connection1")
	if err != nil {
		fmt.Println(err.Error())
	}
	con2, err := database.Connect("connection2")
	if err != nil {
		fmt.Println(err.Error())
	}

	v1.Use(middlewares.DbSetter("dbV1", con1))
	v1.GET("/", controller.GetArticle)

	v2.GET("/", middlewares.DbSetter("dbV2", con2), controller.GetArticles)
	server.Run(":5000")
}
