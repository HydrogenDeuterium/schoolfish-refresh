package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"schoolfish-refresh/middleware"
	"schoolfish-refresh/model"
	"schoolfish-refresh/util"
	"strconv"
)

func Product(g *gin.RouterGroup, db model.DBGroup) {
	g.GET("", func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			fmt.Println(err.Error())
			util.ReturnInternal(c)
			return
		}
		pageSize := 20

		var products []model.Product

		err = db.Mysql.Limit(pageSize).Offset(page*pageSize - page).Find(&products).Error
		if err != nil {
			fmt.Println(err.Error())
			util.ReturnInternal(c)
		}
		util.ReturnGood(c, products)
	})

	g.GET("/users/:uid", func(c *gin.Context) {
		uid := c.Param("uid")

		var user = model.User{}
		result := db.Mysql.Where("uid=?", uid).First(&user)
		if result.RecordNotFound() {
			util.ReturnError(c, "用户不存在！")
			return
		}
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		var products []model.Product
		result = db.Mysql.Where("owner=?", uid).Find(&products)
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, products)
	})

	g.GET("/:pid", func(c *gin.Context) {
		pid := c.Param("pid")
		var product model.Product
		result := db.Mysql.Where("pid=?", pid).First(&product)
		if result.RecordNotFound() {
			util.ReturnError(c, "货物不存在！")
			return
		}
		if err := result.Error; err != nil {
			log.Println(err)
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, product)
	})

	g.POST("", middleware.LogonRequire(), func(c *gin.Context) {

	})
}
