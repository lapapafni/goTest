package main

import (
	"awesomeProject/controllers"
	"github.com/gin-gonic/gin"
)

func app() {
	router := gin.Default()
	router.GET("/ping", controllers.Index)
	router.GET("/:id", controllers.GetOne)
	router.POST("/", controllers.Form)
	router.Run()
}
