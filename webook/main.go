package main

import (
	"GoProject/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"http://localhost:3000"},
		//AllowMethods:     []string{"POST", "PATCH"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//ExposeHeaders:    []string{},

		//cookie相关
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "https://localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	userHandler(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}

}

func userHandler(server *gin.Engine) {
	u := web.NewUserHandler()
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
