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

	g.POST("", func(c *gin.Context) {

	})
}
