package web

import (
	"github.com/gin-gonic/gin"

	regexp "github.com/dlclark/regexp2"
)

type Userhandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler() *Userhandler {
	const (
		emailRegexRattern    = "[\\w]+@[A-Za-z]+(\\.[A-Za-z0-9]+){1,2}"
		passwordRegexPattern = "^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$"
	)

	emailExp := regexp.MustCompile(emailRegexRattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)

	return &Userhandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func registerUsersRoutes(server *gin.Engine) {
	var u *Userhandler
	u = NewUserHandler()
	ug := server.Group("/users")
	ug.POST("signup", u.Signup)
	ug.POST("login", u.Login)
	ug.POST("edit", u.Edit)
	ug.POST("profile", u.Profile)

}

func (u *Userhandler) Signup(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmpassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(200, "系统错误")
	}
	if !ok {
		ctx.String(200, "邮箱格式不正确")
	}

	if req.ConfirmPassword != req.Password {
		ctx.String(200, "两次密码输入不一致")
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(200, "系统错误")
	}
	if !ok {
		ctx.String(200, "密码格式不正确")
	}

}

func (u *Userhandler) Login(ctx *gin.Context) {
}

func (u *Userhandler) Edit(ctx *gin.Context) {
}

func (u *Userhandler) Profile(ctx *gin.Context) {
}
