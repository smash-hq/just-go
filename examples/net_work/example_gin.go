// Package net_work
// Description todo
// Author smash_hq
// Created 2022/12/8 11:55
package net_work

import (
	"github.com/gin-gonic/gin"
	"h7/examples/database"
)

// 声明一组路由
var (
	Engine     = gin.Default()
	RouteGroup = Engine.Group("/h7")
)

func TestGin() {
	data := database.SelectData()
	RouteGroup.GET("/user", func(context *gin.Context) {
		context.JSONP(200, data.ToJSON())
	})
	Engine.Run("localhost:8080")
}
