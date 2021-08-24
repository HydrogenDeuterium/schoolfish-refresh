package main

import (
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/models"
)

func main() {
	DB := models.DBInit()
	defer DB.Close()

	r := gin.Default()
	defer func(r *gin.Engine) {
		_ = r.Run()
	}(r) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
