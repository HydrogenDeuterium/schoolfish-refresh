package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"schoolfish-refresh/model"
	"schoolfish-refresh/util"
	"strings"
)

// LogonRequire 基于JWT的认证中间件
func LogonRequire(db model.DBGroup) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		//for k,v :=range c.Request.Header {
		//	fmt.Println(k,v)
		//}
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			util.ReturnError(c, "请求头中auth为空")
			c.Abort()
			return
		}
		//别问我为啥他要这么处理，非要在前面加个“Bearer ”
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.ReturnError(c, "请求头中auth格式有误")
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		claim, err := util.ParseToken(parts[1])
		if err != nil {
			util.ReturnError(c, "无效的Token")
			c.Abort()
			return
		}

		//确保信息是可靠的
		user := model.User{}
		result := db.Mysql.Where("uid=?", claim.Uid).First(&user)
		if result.RecordNotFound() {
			//仅当用户删除账号后，之前申请的jwt未失效前出现，少见，不用测试
			util.ReturnError(c, "用户未注册!")
			c.Abort()
			return
		}

		if result.Error != nil {
			util.ReturnInternal(c)
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("email", claim.Email)
		c.Set("uid", claim.Uid)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

func CORS() func(c *gin.Context) {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()
		c.Next()
	}

}
