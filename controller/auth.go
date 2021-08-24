package controller

import (
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/models"
)

func Auth(g *gin.RouterGroup, db models.DBs) {
	g.GET("", func(c *gin.Context) {
		//获取的验证码和邮箱绑定
		email := c.Query("email")
		if db.RedisGet(email) != "" {
			//returnError(c, "获取验证码过于频繁！")
			return
		}

		// 往redis存数据，自动定期失效
		err := db.RedisSet(email)
		if err != nil {
			//returnInternal(c, "redis写入失败")
			return
		}
		//err = service.SendMail(email, fmt.Sprintf("验证码为%s", code))
		//if err != nil {
		//	returnInternal(c, "邮件发送失败")
		//	return
		//}
		//returnGood(c, nil)
	})
}
