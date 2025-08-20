package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

// 定义用户相关路由
type UserHandler struct {
}

func (u *UserHandler) SignUp(context *gin.Context) {
	type SignUpRep struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	var req SignUpRep
	//判断数据段是否正确
	if err := context.Bind(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})

	//数据库操作

}
func (u *UserHandler) Login(context *gin.Context) {

}

func (u *UserHandler) Edit(context *gin.Context) {

}

func (u *UserHandler) Profile(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
	})
}
