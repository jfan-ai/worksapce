package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler 我准备在它上面定义跟用户有关的路由

type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	const (
		emailRegexRattern    = "[\\w]+@[A-Za-z]+(\\.[A-Za-z0-9]+){1,2}"
		passwordRegexPattern = "^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$"
	)
	emailExp := regexp.MustCompile(emailRegexRattern, regexp.None)
	PasswordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: PasswordExp,
	}
}

func registerUsersRoutes(server *gin.Engine) {
	var u *UserHandler
	u = NewUserHandler()
	ug := server.Group("/users")
	ug.POST("signup", u.Signup)
	ug.POST("login", u.Login)
	ug.POST("edit", u.Edit)
	ug.GET("profile", u.Profile)
}

func (u *UserHandler) Signup(c *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "邮箱格式错误")
		return
	}

	if req.Password != req.ConfirmPassword {
		c.String(http.StatusOK, "两次输入的密码不一致")
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "密码格式错误")
		return
	}
	c.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) Login(c *gin.Context) {
	// 实现登录逻辑
}

func (u *UserHandler) Edit(c *gin.Context) {
	// 实现编辑逻辑
}

func (u *UserHandler) Profile(c *gin.Context) {
	// 实现获取个人资料逻辑
}
