package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"schoolfish-refresh/model"
	"schoolfish-refresh/service"
)

func Auth(g *gin.RouterGroup, db model.DBGroup) {
	g.GET("", func(c *gin.Context) {
		//获取的验证码和邮箱绑定
		email := c.Query("email")
		if email == "" {
			returnError(c, "请提供邮箱！")
		}
		get, _ := db.RedisGet(email)
		if get != "" {
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
		email := c.DefaultPostForm("email", "")
		if email == "" {
			returnError(c, "请提供邮箱！")
			return
		}

		pwd := c.PostForm("password")
		//code := c.DefaultQuery("code","123456")
		//redis, err := db.RedisGet(email)
		//if err != nil || redis != code {
		//	returnError(c, "验证码无效或已过期！")
		//	return
		//}
		user := &model.User{}
		where := db.Mysql.Where("email=?", email)
		//where := db.Mysql.Model(&model.User{}).Where("email=?", email)
		db := where.First(user)
		if db.RecordNotFound() {
			returnError(c, "用户不存在！")
			return
		}
		if db.Error != nil {
			returnInternal(c)
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(user.Hashed), []byte(pwd)) != nil {
			returnError(c, "用户名与密码不匹配！")
			return
		}
		//TODO 实现计算token
		token, err := getToken(email, pwd)
		if err != nil {
			returnInternal(c)
		}
		returnGood(c, token)
	})
}
