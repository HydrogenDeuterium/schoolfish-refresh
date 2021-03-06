package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"schoolfish-refresh/model"
	"schoolfish-refresh/service"
	"schoolfish-refresh/util"
)

func Auth(g *gin.RouterGroup, db model.DBGroup) {
	g.GET("", func(c *gin.Context) {
		//获取的验证码和邮箱绑定
		email := c.Query("email")
		if email == "" {
			util.ReturnError(c, "请提供邮箱！")
		}
		get, _ := db.RedisGet(email)
		if get != "" {
			util.ReturnError(c, "获取验证码过于频繁！")
			return
		}

		// 往redis存数据，自动定期失效
		code := db.RedisSet(email)
		if code == "" {
			util.ReturnInternal(c)
			return
		}
		err := service.Sendmail(email, fmt.Sprintf("验证码为%s", code))
		if err != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, nil)
	})

	g.DELETE("", func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			util.ReturnError(c, "请提供邮箱！")
			return
		}
		db.RedisDel(email)
		util.ReturnGood(c, nil)
	})

	g.POST("", func(c *gin.Context) {
		email := c.DefaultPostForm("email", "")
		if email == "" {
			util.ReturnError(c, "请提供邮箱！")
			return
		}

		pwd := c.PostForm("password")
		//code := c.DefaultQuery("code","123456")
		//redis, err := db.RedisGet(email)
		//if err != nil || redis != code {
		//	ReturnError(c, "验证码无效或已过期！")
		//	return
		//}
		user := &model.Users{}
		where := db.Mysql.Where("email=?", email)
		//where := db.Mysql.Model(&model.User{}).Where("email=?", email)
		db := where.First(user)
		if db.RecordNotFound() {
			util.ReturnError(c, "用户不存在！")
			return
		}
		if db.Error != nil {
			util.ReturnInternal(c)
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(user.Hashed), []byte(pwd)) != nil {
			util.ReturnError(c, "用户名与密码不匹配！")
			return
		}
		token, err := util.GetToken(user.Uid, email, pwd)
		if err != nil {
			util.ReturnInternal(c)
		}
		util.ReturnGood(c, token)
	})
}
