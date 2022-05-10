package main

import (
	"awesomeProject/controllers"
	"github.com/gin-gonic/gin"
)

func app(port string) {
	router := gin.Default()
	router.GET("/ping", controllers.Index)
	router.GET("/:id", controllers.GetOne)
	router.Run(port)
}
