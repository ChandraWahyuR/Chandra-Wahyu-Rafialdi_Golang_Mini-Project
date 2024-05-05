package routes

import (
	"prototype/api/controller/equipment"
	"prototype/api/controller/equipment/category"
	"prototype/api/controller/rent"

	rentconfirm "prototype/api/controller/rent_confirm"
	user "prototype/api/controller/user"
	"prototype/api/middleware/authorization"
	"prototype/constant"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

type RouteController struct {
	SignUpUser        *user.UserController
	EquipmentRoute    *equipment.EquipmentController
	RentRoute         *rent.RentController
	CategoryEquipment *category.CategoryController
	RentConfirm       *rentconfirm.RentConfirmController
}

func (r *RouteController) InitRoute(e *echo.Echo) {
	e.POST("/register", r.SignUpUser.Register)
	e.POST("/login", r.SignUpUser.Login)

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(constant.PrivateKeyJWT())))

	// User
	eAuth.GET("/equipment", r.EquipmentRoute.GetAll)
	eAuth.GET("/equipment/:id", r.EquipmentRoute.GetById)

	// # rent
	eAuth.POST("/rent", r.RentRoute.PostRent)
	eAuth.PUT("/rent/:id", r.RentRoute.UpdateRent)
	eAuth.GET("/rent/user/:id", r.RentRoute.GetByUserID)

	// # rent confirm
	eAuth.POST("/confirm", r.RentConfirm.PostRentConfirm)
	eAuth.GET("/confirm", r.RentConfirm.GetAll)
	eAuth.GET("/confirm/:id", r.RentConfirm.GetById)

	// Admin
	eAuth.POST("/admin/equipment", r.EquipmentRoute.PostEquipment, authorization.OnlyAdmin)
	eAuth.DELETE("/admin/equipment/:id", r.EquipmentRoute.DeleteEquipment, authorization.OnlyAdmin)

	// Category
	eAuth.GET("/equipment/category", r.CategoryEquipment.GetAll)
	eAuth.POST("/admin/equipment/category", r.CategoryEquipment.PostCategory, authorization.OnlyAdmin)
	eAuth.DELETE("/admin/equipment/category/:id", r.CategoryEquipment.DeleteCategory, authorization.OnlyAdmin)

	// # rent
	eAuth.GET("/rent", r.RentRoute.GetAll, authorization.OnlyAdmin)
	eAuth.GET("/rent/:id", r.RentRoute.GetById, authorization.OnlyAdmin)
	eAuth.DELETE("/rent/:id", r.RentRoute.DeleteRent, authorization.OnlyAdmin)

	// # rent confirm

}
