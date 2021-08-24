package main

import (
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/controller"
	"schoolfish-refresh/models"
)

func main() {
	DB := models.DBInit()
	defer DB.Close()

	r := gin.Default()
	defer func(r *gin.Engine) {
		_ = r.Run()
	}(r) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	controller.Auth(r.Group("/auth"), DB)
}
