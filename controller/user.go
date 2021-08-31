package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"schoolfish-refresh/model"
)

func User(g *gin.RouterGroup, db model.DBGroup) {
	g.POST("", func(c *gin.Context) {
		email := c.DefaultPostForm("email", "")
		if email == "" {
			returnError(c, "提供邮箱！")
			return
		}
		find := db.Mysql.Where("email = ?", email).First(&model.User{}).RecordNotFound()
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
		user := &model.User{
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

	g.GET("/", JWTAuthMiddleware(), func(c *gin.Context) {
		uid, exist := c.Get("uid")
		if exist == false {
			returnInternal(c)
			return
		}
		user := model.User{}
		if db.Mysql.Where("uid=?", uid).First(&user).RecordNotFound() {
			returnError(c, "用户未注册!")
			return
		}
		data := Struct2Map(user)
		returnGood(c, data)

	})

	g.PUT("/", JWTAuthMiddleware(), func(c *gin.Context) {
		//uid := c.Param("uid")
		uid, exist := c.Get("uid")
		if exist == false {
			returnInternal(c)
			return
		}
		if db.Mysql.Where("uid=?", uid).First(&model.User{}).RecordNotFound() {
			returnError(c, "用户未注册!")
		}
		userMap := c.PostFormMap("user")
		user := &model.User{
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
		returnGood(c, Struct2Map(user))
	})
}
