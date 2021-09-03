package controller

import (
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/middleware"
	"schoolfish-refresh/model"
	"schoolfish-refresh/util"
	"strconv"
)

func Comment(g *gin.RouterGroup, db model.DBGroup) {

	g.GET("/products/:pid", func(c *gin.Context) {
		pid, err := strconv.Atoi(c.Param("pid"))
		if err != nil || pid <= 0 {
			util.ReturnError(c, "pid格式不正确！")
			return
		}
		var comments []model.Comment
		err = db.Mysql.Model(model.Comment{}).Where("product=?", pid).Find(&comments).Error
		if err != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, comments)
	})

	g.GET(":cid/response", func(c *gin.Context) {

	})

	g.POST("/products/:pid", middleware.LogonRequire(db), func(c *gin.Context) {
		uid, _ := c.Get("uid")
		pid, err := strconv.Atoi(c.Param("pid"))
		if err != nil || pid <= 0 {
			util.ReturnError(c, "pid格式不正确！")
			return
		}
		text := c.PostForm("text")
		comment := model.Comment{
			Product:     uint(pid),
			Commentator: uid.(uint),
			ResponseTo:  0,
			Text:        text,
		}
		err = db.Mysql.Model(&model.Comment{}).Create(&comment).Error

		if err != nil {
			util.ReturnInternal(c)
			return
		}
		result := db.Mysql.Model(&model.Comment{}).Where("commentator=?", uid).First(&comment)
		if result.Error != nil || result.RecordNotFound() {
			util.ReturnInternal(c)
		}
		util.ReturnGood(c, comment)
	})

	g.GET("/:cid", func(c *gin.Context) {
		cid, err := strconv.Atoi(c.Param("cid"))
		if err != nil || cid <= 0 {
			util.ReturnError(c, "cid格式不正确！")
			return
		}
		comment := model.Comment{}
		result := db.Mysql.Model(&model.Comment{}).Where("cid=?", cid).First(&comment)
		if result.Error != nil || result.RecordNotFound() {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, comment)
	})

	g.POST("/:cid", middleware.LogonRequire(db), func(c *gin.Context) {

	})

	g.PUT("/:cid", middleware.LogonRequire(db), func(c *gin.Context) {
		util.ReturnGood(c, "暂未实现！")
	})

	g.DELETE("/:cid", middleware.LogonRequire(db), func(c *gin.Context) {

	})
}
