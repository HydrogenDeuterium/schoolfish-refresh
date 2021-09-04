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
		var comments []model.Comments
		err = db.Mysql.Model(model.Comments{}).Where("product=?", pid).Limit(10).Find(&comments).Error
		if err != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, comments)
	})

	g.POST("/products/:pid", middleware.LogonRequire(db), func(c *gin.Context) {
		uid, _ := c.Get("uid")
		pid, err := strconv.Atoi(c.Param("pid"))
		if err != nil || pid <= 0 {
			util.ReturnError(c, "pid格式不正确！")
			return
		}
		text := c.PostForm("text")
		comment := model.Comments{
			Product:     uint(pid),
			Commentator: uid.(uint),
			ResponseTo:  nil,
			Text:        &text,
		}
		err = db.Mysql.Model(&model.Comments{}).Create(&comment).Error

		if err != nil {
			util.ReturnInternal(c)
			return
		}
		result := db.Mysql.Model(&model.Comments{}).Where("commentator=?", uid).First(&comment)
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
		comment := model.Comments{}
		result := db.Mysql.Model(&model.Comments{}).Where("cid=?", cid).First(&comment)
		if result.Error != nil || result.RecordNotFound() {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, comment)
	})

	g.PUT("/:cid", middleware.LogonRequire(db), func(c *gin.Context) {
		util.ReturnGood(c, "暂未实现！")
	})

	g.DELETE("/:cid", middleware.LogonRequire(db), func(c *gin.Context) {
		uid, _ := c.Get("uid")
		cid, err := strconv.Atoi(c.Param("cid"))
		if err != nil || cid <= 0 {
			util.ReturnError(c, "cid格式不正确！")
			return
		}
		comment := model.Comments{}
		result := db.Mysql.Model(&model.Comments{}).Where("cid=?", cid).First(&comment)
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		if comment.Commentator != uid {
			util.ReturnError(c, "无权操作！")
			return
		}
		result = db.Mysql.Model(&model.Comments{}).Where("cid=?", cid).Delete(&comment)
		if result.Error != nil {
			util.ReturnInternal(c)
		}
		util.ReturnGood(c, comment)
	})

	g.GET(":cid/response", func(c *gin.Context) {
		cid, err := strconv.Atoi(c.Param("cid"))
		if err != nil || cid <= 0 {
			util.ReturnError(c, "pid格式不正确！")
			return
		}
		var comments []model.Comments
		err = db.Mysql.Model(model.Comments{}).Where("response_to=?", cid).Find(&comments).Error
		if err != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, comments)
	})

	g.POST("/:cid/response", middleware.LogonRequire(db), func(c *gin.Context) {
		uid, _ := c.Get("uid")
		cid, err := strconv.Atoi(c.Param("cid"))
		if err != nil || cid <= 0 {
			util.ReturnError(c, "cid格式不正确！")
			return
		}
		var commentTo model.Comments
		err = db.Mysql.Model(&commentTo).Where("").First(&commentTo).Error
		if err != nil {
			util.ReturnInternal(c)
			return
		}
		text := c.PostForm("text")
		uc := uint(cid)
		comment := model.Comments{
			Product:     commentTo.Product,
			Commentator: uid.(uint),
			ResponseTo:  &uc,
			Text:        &text,
		}
		err = db.Mysql.Model(&model.Comments{}).Create(&comment).Error

		if err != nil {
			util.ReturnInternal(c)
			return
		}
		result := db.Mysql.Model(&model.Comments{}).Where("commentator=?", uid).First(&comment)
		if result.Error != nil || result.RecordNotFound() {
			util.ReturnInternal(c)
		}
		util.ReturnGood(c, comment)
	})
}
