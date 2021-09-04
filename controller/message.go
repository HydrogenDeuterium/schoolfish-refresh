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
		if self == "" {
			util.ReturnError(c, "不正确的uid！")
			return
		}
		util.ReturnGood(c, nil)
	})

	c.GET("/:uid", func(c *gin.Context) {
		self, _ := c.Get("uid")
		if self == "" {
			util.ReturnError(c, "不正确的uid！")
			return
		}
		util.ReturnGood(c, nil)
	})

	c.POST("/:uid", func(c *gin.Context) {
		self, _ := c.Get("uid")
		self, err := strconv.Atoi(self.(string))
		if err != nil {
			util.ReturnError(c, "不正确的uid！")
			return
		}
		other := c.Param("uid")
		uid, _ := strconv.Atoi(other)
		text := c.Param("text")
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
		util.ReturnGood(c, nil)
	})

}
