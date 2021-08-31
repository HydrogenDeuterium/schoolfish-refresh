package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"schoolfish-refresh/model"
	"strconv"
)

func Product(g *gin.RouterGroup, db model.DBGroup) {
	g.GET("", func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			fmt.Println(err.Error())
			returnInternal(c)
			return
		}
		pageSize := 20

		var products []model.Product

		err = db.Mysql.Limit(pageSize).Offset(page*pageSize - page).Find(&products).Error
		if err != nil {
			fmt.Println(err.Error())
			returnInternal(c)
		}
		returnGood(c, products)
	})

	g.GET("/user/:uid", func(c *gin.Context) {
		uid, exists := c.Get("uid")
		if exists == false {
			returnInternal(c)
			return
		}
		user := model.User{}
		result := db.Mysql.Where("uid=?", uid).First(&user)
		if result.RecordNotFound() {
			returnError(c, "用户不存在！")
			return
		}
		if result.Error != nil {
			returnInternal(c)
			return
		}
		var products []model.Product
		result = db.Mysql.Where("owner=?", uid).Find(&products)
		if result.Error != nil {
			returnInternal(c)
			return
		}
		returnGood(c, products)
	})

}
