package routes

import (
	user "prototype/api/controller/user"
	"prototype/api/middleware/authorization"
	"prototype/constant"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	SignUpUser *user.UserController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	e.POST("/register", r.SignUpUser.Register)
	e.POST("/login", r.SignUpUser.Login)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(constant.PrivateKeyJWT())))

	// User

	// Admin
	eAuth.POST("/admin/equpment", r.SignUpUser.Login, authorization.OnlyAdmin)
}
