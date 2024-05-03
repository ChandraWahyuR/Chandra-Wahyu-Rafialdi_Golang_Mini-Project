package main

import (
	controllersEq "prototype/api/controller/equipment"
	controllersRent "prototype/api/controller/rent"
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
	rentRepo := repository.NewRentRepo(db)

	userUseCase := usecase.NewUserUseCase(userRepo)
	equipmentUseCase := usecase.NewEquipmentUseCase(equipmentRepo)
	rentUseCase := usecase.NewRentUseCase(rentRepo)

	userController := controllers.NewUserController(userUseCase)
	equipmentController := controllersEq.NewEquipmentController(equipmentUseCase)
	RentController := controllersRent.NewRentController(rentUseCase, equipmentUseCase)

	routes := routes.RouteController{
		SignUpUser:     userController,
		EquipmentRoute: equipmentController,
		RentRoute:      RentController,
	}

	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
