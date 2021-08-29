package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/models"
)

func User(g *gin.RouterGroup, db models.DBGroup) {
	g.POST("", func(c *gin.Context) {
		userMap := c.PostFormMap("user")
		err := db.Mysql.Where("email = ?", userMap["email"])
		if err != nil {
			returnError(c, "用户已注册")
		}
		user := &models.Users{
			Username: userMap["username"],
			Email:    userMap["email"],
			Hashed:   userMap["hashed"],
			Avatar:   userMap["avatar"],
			Info:     userMap["info"],
			Profile:  userMap["profile"],
			Location: userMap["location"],
		}
		db.Mysql.Create(user)
	})

	g.GET("/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		var user *models.Users
		if db.Mysql.Where("uid=?", uid).First(user).RecordNotFound() {
			returnError(c, "用户未注册!")
		}
		data, err := json.Marshal(user)
		if err != nil {
			returnInternal(c)
		}
		returnGood(c, data)

	})

	g.PUT("/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		if db.Mysql.Where("uid=?", uid).First(&models.Users{}).RecordNotFound() {
			returnError(c, "用户未注册!")
		}
		userMap := c.PostFormMap("user")
		user := &models.Users{
			Username: userMap["username"],
			Email:    userMap["email"],
			Hashed:   userMap["hashed"],
			Avatar:   userMap["avatar"],
			Info:     userMap["info"],
			Profile:  userMap["profile"],
			Location: userMap["location"],
		}
		err := db.Mysql.Update("email = ?", userMap["email"])
		if err != nil {
			returnInternal(c)
		}
		db.Mysql.Create(user)
	})
}
