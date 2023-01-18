package main

import (
	"github.com/brian-ding/azure-manager/internal/pkg/refreship"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/refresh", refreship.RefreshHandler)
	r.GET("/refresh/:id", refreship.CheckHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
