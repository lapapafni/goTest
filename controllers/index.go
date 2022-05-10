package controllers

import (
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
