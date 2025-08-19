package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	//静态路由
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	server.POST("/post1", func(c *gin.Context) {
		c.String(http.StatusOK, "post 1")
	})

	//参数路由 1 RESTful  /get2/name
	server.GET("/get2/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "get2,参数路由: %s", name)
	})

	//参数路由 2	/order?id=
	server.GET("/order", func(c *gin.Context) {
		id := c.Query("id")
		c.String(http.StatusOK, "order,参数路由: %s", id)
	})

	//通配符路由
	server.GET("/get3/*.html", func(c *gin.Context) {
		page := c.Param(".html")
		c.String(http.StatusOK, "get3.html,通配符路由 %s", page)
	})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := server.Run(":8080")
	if err != nil {
		return
	}

}
