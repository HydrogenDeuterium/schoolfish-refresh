package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"schoolfish-refresh/models"
	"schoolfish-refresh/service"
)

func Auth(g *gin.RouterGroup, db models.DBs) {
	g.GET("", func(c *gin.Context) {
		//获取的验证码和邮箱绑定
		email := c.Query("email")
		if email == "" {
			returnError(c, "请提供邮箱！")
		}
		if db.RedisGet(email) != "" {
			returnError(c, "获取验证码过于频繁！")
			return
		}

		// 往redis存数据，自动定期失效
		code := db.RedisSet(email)
		if code == "" {
			returnInternal(c)
			return
		}
		err := service.Sendmail(email, fmt.Sprintf("验证码为%s", code))
		if err != nil {
			returnInternal(c)
			return
		}
		returnGood(c, nil)
	})

	g.DELETE("", func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			returnError(c, "请提供邮箱！")
			return
		}
		db.RedisDel(email)
		returnGood(c, nil)
	})

	g.POST("", func(c *gin.Context) {
		email := c.PostForm("email")
		pwd := c.PostForm("password")
		code := c.Query("code")
		if db.RedisGet(email) != code {
			returnError(c, "验证码无效或已过期！")
			return
		}
		var user *models.Users
		db.Mysql.Where("email=?", email).Find(user)
		if user == nil {
			returnError(c, "用户不存在！")
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(user.Hashed), []byte(pwd)) != nil {
			returnError(c, "用户名与密码不匹配！")
			return
		}
		//TODO 实现计算token
		token := "123456"
		returnGood(c, token)
	})
}
