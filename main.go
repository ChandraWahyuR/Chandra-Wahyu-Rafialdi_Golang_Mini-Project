package main

import (
	controllersEq "prototype/api/controller/equipment"
	controllersCt "prototype/api/controller/equipment/category"
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
	categoryRepo := repository.NewCategoryRepo(db)

	userUseCase := usecase.NewUserUseCase(userRepo)
	equipmentUseCase := usecase.NewEquipmentUseCase(equipmentRepo)
	rentUseCase := usecase.NewRentUseCase(rentRepo)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)

	userController := controllers.NewUserController(userUseCase)
	equipmentController := controllersEq.NewEquipmentController(equipmentUseCase, categoryUseCase)
	RentController := controllersRent.NewRentController(rentUseCase, equipmentUseCase)
	categoryController := controllersCt.NewCategoryController(categoryUseCase)

	routes := routes.RouteController{
		SignUpUser:        userController,
		EquipmentRoute:    equipmentController,
		RentRoute:         RentController,
		CategoryEquipment: categoryController,
	}

	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
