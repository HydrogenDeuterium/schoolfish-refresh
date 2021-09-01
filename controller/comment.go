package controller

import (
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/middleware"
	"schoolfish-refresh/model"
	"schoolfish-refresh/util"
)

func Comment(g *gin.RouterGroup, db model.DBGroup) {
	g.GET(":cid", func(c *gin.Context) {

	})

	g.GET("/products/:pid", func(c *gin.Context) {

	})

	g.GET(":cid/response", func(c *gin.Context) {

	})

	g.POST("/comments/:pid", middleware.LogonRequire(db), func(c *gin.Context) {

	})

	g.POST("/:cid", middleware.LogonRequire(db), func(c *gin.Context) {

	})

	g.PUT("/:cid", middleware.LogonRequire(db), func(c *gin.Context) {
		util.ReturnGood(c, "暂未实现！")
	})

	g.DELETE("/:cid", middleware.LogonRequire(db), func(c *gin.Context) {

	})
}
