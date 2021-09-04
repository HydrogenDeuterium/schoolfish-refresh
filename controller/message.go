package controller

import (
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/middleware"
	"schoolfish-refresh/model"
	"schoolfish-refresh/util"
	"strconv"
)

func Message(c *gin.RouterGroup, DB model.DBGroup) {
	//私信都要登录
	c.Use(middleware.LogonRequire(DB))

	c.GET("", func(c *gin.Context) {
		self, _ := c.Get("uid")
		var messages []model.Messages
		db := DB.Mysql.Model(model.Messages{}).Limit(10)
		query := db.Where("`from` =?", self).Or("`to`=?", self)

		result := query.Debug().Find(&messages)
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, messages)
	})

	c.GET("/:uid", func(c *gin.Context) {
		self, _ := c.Get("uid")
		other := c.Param("uid")
		uid, err := strconv.Atoi(other)
		if err != nil || uid <= 0 {
			util.ReturnError(c, "不正确的uid！")
			return
		}
		var messages []model.Messages
		db := DB.Mysql.Model(model.Messages{}).Limit(10)
		users := []string{strconv.Itoa(int(self.(uint))), other}
		query := db.Where("`from` in (?)", users).Where("`to` in (?)", users).Order("mid")
		result := query.Find(&messages)
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, messages)
	})

	c.POST("/:uid", func(c *gin.Context) {
		self, _ := c.Get("uid")
		other := c.Param("uid")
		uid, err := strconv.Atoi(other)
		if err != nil || uid <= 0 || uint(uid) == self {
			util.ReturnError(c, "不正确的uid！")
			return
		}
		text := c.PostForm("text")
		message := model.Messages{
			From: self.(uint),
			To:   uint(uid),
			Text: &text,
		}
		err = DB.Mysql.Model(&message).Create(&message).Error
		if err != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, message)
	})
}
