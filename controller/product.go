package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"schoolfish-refresh/middleware"
	"schoolfish-refresh/model"
	"schoolfish-refresh/util"
	"strconv"
	"strings"
)

func Product(g *gin.RouterGroup, db model.DBGroup) {
	g.GET("", func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil {
			fmt.Println(err.Error())
			util.ReturnInternal(c)
			return
		}
		pageSize := 10

		var products []model.Products

		result := db.Mysql.Model(model.Products{}).Limit(pageSize).Offset(page*pageSize - pageSize).Find(&products)

		if result.RecordNotFound() {
			fmt.Println("你妈的为什么")
		}
		if result.Error != nil {
			fmt.Println(result.Error.Error())
			util.ReturnInternal(c)
		}
		util.ReturnGood(c, products)
	})

	g.GET("/users/:uid", func(c *gin.Context) {
		uid := c.Param("uid")

		var user = model.Users{}
		result := db.Mysql.Where("uid=?", uid).First(&user)
		if result.RecordNotFound() {
			util.ReturnError(c, "用户不存在！")
			return
		}
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		var products []model.Products
		result = db.Mysql.Where("owner=?", uid).Find(&products)
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, products)
	})

	g.GET("/:pid", func(c *gin.Context) {
		pid := c.Param("pid")
		var product model.Products
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

	g.POST("", middleware.LogonRequire(db), func(c *gin.Context) {
		uid, exist := c.Get("uid")
		if exist == false {
			util.ReturnInternal(c)
			return
		}
		var product model.Products
		//golang 不支持双引号 json
		productJson := strings.Replace(c.PostForm("product"), "'", "\"", -1)
		err := json.Unmarshal([]byte(productJson), &product)
		if err != nil {
			util.ReturnInternal(c)
			return
		}
		product.Owner = uid.(uint)

		result := db.Mysql.Create(&product)
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, product)

	})

	g.PUT("/:pid", middleware.LogonRequire(db), func(c *gin.Context) {
		uid, exist := c.Get("uid")
		if exist == false {
			util.ReturnInternal(c)
			return
		}

		pid := c.Param("pid")

		var product model.Products
		//golang 不支持双引号 json
		result := db.Mysql.Where("pid=?", pid).First(&product)
		if result.RecordNotFound() {
			util.ReturnError(c, "商品不存在！")
			return
		}
		productJson := strings.Replace(c.PostForm("product"), "'", "\"", -1)
		err := json.Unmarshal([]byte(productJson), &product)
		if err != nil {
			util.ReturnInternal(c)
			return
		}
		product.Owner = uid.(uint)

		query := db.Mysql.Where("pid=?", pid).Limit(1)
		//fmt.Println(query.Value)
		result = query.Model(&model.Products{}).Update(&product)
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, product)

	})

	g.DELETE("/:pid", middleware.LogonRequire(db), func(c *gin.Context) {
		uid, exist := c.Get("uid")
		if exist == false {
			util.ReturnInternal(c)
			return
		}

		pid := c.Param("pid")

		var product model.Products
		//golang 不支持双引号 json
		result := db.Mysql.Where("pid=?", pid).First(&product)
		if result.RecordNotFound() {
			util.ReturnError(c, "不能删除不存在的商品！")
			return
		}
		result = db.Mysql.Where("owner=?", uid).First(&product)
		if result.RecordNotFound() {
			util.ReturnError(c, "没有所有权！")
		}

		result = db.Mysql.Model(&model.Products{}).Delete(&product)
		if result.RowsAffected != 1 {
			//我寻思一次应该只能删掉一个
			result.Rollback()
			util.ReturnInternal(c)
		}
		if result.Error != nil {
			util.ReturnInternal(c)
			return
		}
		util.ReturnGood(c, product)

	})
}
