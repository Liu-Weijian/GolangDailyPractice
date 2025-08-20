package main

import (
	"GoProject/webook/internal/web"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	userHandler(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}

}

func userHandler(server *gin.Engine) {
	u := &web.UserHandler{}
	//非RESTful风格
	server.POST("/users/signup", u.SignUp)

	server.POST("/users/login", u.Login)

	server.POST("/users/edit", u.Edit)

	server.GET("/users/profile", u.Profile)

	//RESTful风格
	//注册
	//server.PUT("/users", func(context *gin.Context) {})
	//编辑
	//server.POST("/users/:id", func(context *gin.Context) {})
	//删除
	//server.delete("/users", func(context *gin.Context) {})

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
