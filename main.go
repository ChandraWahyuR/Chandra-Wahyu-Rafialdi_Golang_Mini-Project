package main

import (
	controllersEq "prototype/api/controller/equipment"
	controllers "prototype/api/controller/user"
	"prototype/api/routes"
	"prototype/config"
	"prototype/drivers"
	"prototype/repository"
	"prototype/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadFileEnv()
	config.InitConfigDb()
	db := drivers.ConnectDB(config.InitConfigDb())

	e := echo.New()

	userRepo := repository.NewUserRepo(db)
	equipmentRepo := repository.NewEquipmentRepo(db)

	userUseCase := usecase.NewUserUseCase(userRepo)
	equipmentUseCase := usecase.NewEquipmentUseCase(equipmentRepo)

	userController := controllers.NewUserController(userUseCase)
	equipmentController := controllersEq.NewEquipmentController(equipmentUseCase)

	routes := routes.RouteController{
		SignUpUser:     userController,
		EquipmentRoute: equipmentController,
	}
	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
