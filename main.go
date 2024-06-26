package main

import (
	controllersCt "prototype/api/controller/category"
	controllersEq "prototype/api/controller/equipment"
	controllersRent "prototype/api/controller/rent"
	controllersConf "prototype/api/controller/rent_confirm"
	controllers "prototype/api/controller/user"
	"prototype/api/routes"
	"prototype/config"
	"prototype/drivers"
	"prototype/repository"
	"prototype/usecase"

	_ "prototype/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Chandra Mini Project
// @version 1.0
// @description Mini Project tentang penyewaan alat-alat bertema lingkungan mulai dari alat perawatan tanaman, kebersihan lingkungan, alat hiking dan camping, dan alat proses daur ulang.
// @contact.url http://www.swagger.io/support
// @chandrawahyurafialdi.email support@swagger.io
// @BasePath /v2
func main() {
	config.LoadFileEnv()
	config.InitConfigDb()
	db := drivers.ConnectDB(config.InitConfigDb())

	e := echo.New()

	userRepo := repository.NewUserRepo(db)
	equipmentRepo := repository.NewEquipmentRepo(db)
	rentRepo := repository.NewRentRepo(db)
	categoryRepo := repository.NewCategoryRepo(db)
	confirmRentRepo := repository.NewRentConfirmRepo(db)

	userUseCase := usecase.NewUserUseCase(userRepo)
	equipmentUseCase := usecase.NewEquipmentUseCase(equipmentRepo)
	rentUseCase := usecase.NewRentUseCase(rentRepo)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepo)
	rentConfirmUseCase := usecase.NewRentConfirmUseCase(confirmRentRepo)

	userController := controllers.NewUserController(userUseCase)
	equipmentController := controllersEq.NewEquipmentController(equipmentUseCase, categoryUseCase)
	RentController := controllersRent.NewRentController(rentUseCase, equipmentUseCase, userUseCase)
	categoryController := controllersCt.NewCategoryController(categoryUseCase)
	rentConfirmController := controllersConf.NewRentConfirmController(rentConfirmUseCase, rentUseCase, userUseCase, equipmentUseCase)

	// Chat bot
	routes := routes.RouteController{
		SignUpUser:        userController,
		EquipmentRoute:    equipmentController,
		RentRoute:         RentController,
		CategoryEquipment: categoryController,
		RentConfirm:       rentConfirmController,
	}
	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}
