package controller

import "github.com/gin-gonic/gin"

func returnJson(c *gin.Context, code int, data interface{}, message string) {
	c.JSON(200, gin.H{
		"code": code,
		"data": data,
		"msg":  message,
	})
}

//好的返回值不用给信息，能得到数据就行
func returnGood(c *gin.Context, data interface{}) {
	returnJson(c, 200, data, "请求成功")
}

//坏的返回值不用给数据，告诉前端哪里有问题
func returnError(c *gin.Context, msg string) {
	returnJson(c, 400, nil, msg)
}

//内部错误啥都不用给
func returnInternal(c *gin.Context) {
	returnJson(c, 500, nil, "")
}
