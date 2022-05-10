package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "lol",
	})
}

func GetOne(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"id": id,
	})
}

func Form(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.Json("name")
	message := c.PostForm("message")

	fmt.Printf("id %s page %s name %s message: %s", id, page, name, message)
}
