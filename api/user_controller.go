package api

import (
	"goframework/lib"
	"goframework/model"

	"github.com/gin-gonic/gin"
)

func (s *Server) Profile(c *gin.Context) (data interface{}, error *lib.Error) {
	// claims, flag := c.Get("claims")
	// if flag {

	// }
	// uid := claims["uid"]
	uid := 1 //todo
	data, err := s.UserService.FindById(uid)
	if err != nil {
		error = lib.NewInternalError(err.Error())
	}

	return
}

type RegisterParam struct {
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" binding:"required,min=6,max=10" label:"密码"`
}

func (s *Server) Register(c *gin.Context) (data interface{}, error *lib.Error) {
	var param RegisterParam
	if err := c.ShouldBind(&param); err != nil {
		error = lib.NewAutoParamError(err)
		return
	}
	user := &model.User{
		Email:    param.Email,
		Password: param.Password,
	}
	user, err := s.UserService.Register(user)
	if err != nil {
		error = lib.NewInternalError(err.Error())
		return
	}
	data, err = lib.CreateToken(int(user.ID))
	if err != nil {
		error = lib.NewInternalError(err.Error())
		return
	}
	return
}

type LoginParam struct {
	Email    string `json:"email" binding:"required,email" label:"邮箱"`
	Password string `json:"password" binding:"required,min=6,max=10" label:"密码"`
}

func (s *Server) Login(c *gin.Context) (data interface{}, error *lib.Error) {
	var param LoginParam
	if err := c.ShouldBind(&param); err != nil {
		error = lib.NewAutoParamError(err)
		return
	}
	user, err := s.UserService.Login(param.Email, param.Password)
	if err != nil {
		error = lib.NewInternalError(err.Error())
		return
	}
	data, err = lib.CreateToken(int(user.ID))
	if err != nil {
		error = lib.NewInternalError(err.Error())
		return
	}
	return
}

func (s *Server) Logout(c *gin.Context) (data interface{}, error *lib.Error) {
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjA4OTY2NzAsImlkIjo1fQ.SzGEjRkYdVhnbweObtaNs-Gy39dUUItqAnhU3Cm2ejM"

	return
}

type RefreshTokenParam struct {
	RefreshToken string `json:"refresh_token" binding:"required" label:"token"`
}

func (s *Server) RefreshToken(c *gin.Context) (data interface{}, error *lib.Error) {
	var param RefreshTokenParam
	if err := c.ShouldBind(&param); err != nil {
		error = lib.NewAutoParamError(err)
		return
	}
	data, err := lib.RefreshToken(param.RefreshToken)
	if err != nil {
		error = lib.NewInternalError(err.Error())
		return
	}
	return
}
