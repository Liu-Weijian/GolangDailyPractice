package web

import (
	"GoProject/webook/internal/domain"
	"GoProject/webook/internal/service"
	"fmt"
	regexp "github.com/dlclark/regexp2"
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
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	svc         *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:         svc,
	}

}

func (u *UserHandler) SignUp(context *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	var req SignUpReq
	//判断数据段是否正确
	if err := context.Bind(&req); err != nil {
		fmt.Println("错误")
		context.String(http.StatusOK, "系统错误")
		return
	}
	fmt.Println(req.Email)
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		context.String(http.StatusOK, "邮箱格式错误")
		return
	}

	if !ok {
		context.String(http.StatusOK, "邮箱格式错误")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		context.String(http.StatusOK, "密码格式错误")
		return
	}
	if !ok {
		context.String(http.StatusOK, "密码格式错误")
		return
	}

	//调用service
	err = u.svc.Signup(context, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		context.String(http.StatusOK, "注册失败")
		return
	}

	context.String(http.StatusOK, "注册成功")

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
