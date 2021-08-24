package main

import (
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/models"
)

func main() {
	rdb := models.RedisInit()
	mdb := models.MysqlInit()

	models.RedisTest()

	defer rdb.Close()
	defer mdb.Close()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
