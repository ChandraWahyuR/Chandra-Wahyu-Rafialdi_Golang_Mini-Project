package main

import (
	controllers "prototype/api/controller/user"
	"prototype/api/routes"
	"prototype/config"
	"prototype/drivers"
	"prototype/repository"
	"prototype/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func main() {
	config.LoadFileEnv()
	config.InitConfigDb()
	db := drivers.ConnectDB(config.InitConfigDb())

	e := echo.New()

	userRepo := repository.NewUserRepo(db)

	userUseCase := usecase.NewUserUseCase(userRepo)

	userController := controllers.NewUserController(userUseCase)

	routes := routes.RouteController{
		SignUpUser: userController,
	}
	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}

// refrence
// https://dev.to/karanpratapsingh/connecting-to-postgresql-using-gorm-24fj

// Unique
// https://stackoverflow.com/questions/73701967/make-user-email-field-unique
