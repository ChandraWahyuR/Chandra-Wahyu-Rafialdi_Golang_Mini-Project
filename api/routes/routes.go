package routes

import (
	"prototype/api/controller/equipment"
	"prototype/api/controller/rent"
	user "prototype/api/controller/user"
	"prototype/api/middleware/authorization"
	"prototype/constant"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	SignUpUser     *user.UserController
	EquipmentRoute *equipment.EquipmentController
	RentRoute      *rent.RentController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	e.POST("/register", r.SignUpUser.Register)
	e.POST("/login", r.SignUpUser.Login)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(constant.PrivateKeyJWT())))

	// User
	eAuth.GET("/equipment", r.EquipmentRoute.GetAll)
	eAuth.GET("/equipment/:id", r.EquipmentRoute.GetById)
	// Admin
	eAuth.POST("/admin/equipment", r.EquipmentRoute.PostEquipment, authorization.OnlyAdmin)
	eAuth.DELETE("/admin/equipment/:id", r.EquipmentRoute.DeleteEquipment, authorization.OnlyAdmin)
}
