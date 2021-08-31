package main

import (
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/controller"
	"schoolfish-refresh/middleware"
	"schoolfish-refresh/model"
)

func main() {
	DB := model.DBInit()
	defer DB.Close()

	r := gin.Default()

	r.Use(middleware.CORS())

	defer func(r *gin.Engine) {
		_ = r.Run()
	}(r) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	controller.Auth(r.Group("/auth"), DB)
	controller.User(r.Group("/users"), DB)
	controller.Product(r.Group("/products"), DB)
}
