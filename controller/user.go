package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"schoolfish-refresh/models"
)

func User(g *gin.RouterGroup, db models.DBGroup) {
	g.POST("", func(c *gin.Context) {
		email := c.DefaultPostForm("email", "")
		if email == "" {
			returnError(c, "提供邮箱！")
			return
		}
		find := db.Mysql.Where("email = ?", email).First(&models.Users{}).RecordNotFound()
		if find == false {
			returnError(c, "用户已注册!")
			return
		}
		rawPassword := c.DefaultPostForm("password", "")
		if rawPassword == "" {
			returnError(c, "提供密码！")
			return
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
		if err != nil {
			returnInternal(c)
			return
		}
		user := &models.Users{
			Username: c.DefaultPostForm("username", ""),
			Email:    email,
			Hashed:   string(hashed),
			Avatar:   c.DefaultPostForm("avatar", ""),
			Info:     c.DefaultPostForm("info", ""),
			Profile:  c.DefaultPostForm("profile", ""),
			Location: c.DefaultPostForm("location", ""),
		}
		db.Mysql.Create(user)
		returnGood(c, user)
	})

	g.GET("/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		var user *models.Users
		if db.Mysql.Where("uid=?", uid).First(user).RecordNotFound() {
			returnError(c, "用户未注册!")
			return
		}
		data, err := json.Marshal(user)
		if err != nil {
			returnInternal(c)
			return
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
