package controllers

import (
	"net/http"
	"prototype/api/controller/user/request"
	"prototype/api/controller/user/response"
	"prototype/api/middleware"
	"prototype/domain"
	"prototype/utils"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase domain.UseCaseInterface
}

func (uc *UserController) Register(c echo.Context) error {
	var userRegister request.UserRegister
	c.Bind(&userRegister)

	// Validate password with regex
	if !utils.ValidatePassword(userRegister.Password) {
		return c.JSON(http.StatusBadRequest, domain.NewErrorResponse("Invalid password format"))
	}

	// Confirm password user
	if userRegister.Password != userRegister.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, domain.NewErrorResponse("Password and confirm password do not match"))
	}

	// Hashing password
	hashedPassword, err := utils.HashPassword(userRegister.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.NewErrorResponse("Failed to hash password"))
	}

	userValid := userRegister.ToEntities()
	userValid.Password = hashedPassword

	// Register Data ke database
	result, err := uc.userUseCase.Register(userRegister.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), domain.NewErrorResponse(err.Error()))
	}

	user := response.FromUseCase(&result)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Register Success", user))
}

// Login
func (uc *UserController) Login(c echo.Context) error {
	var userLogin request.UserLogin
	c.Bind(&userLogin)
	result, err := uc.userUseCase.Login(userLogin.ToEntities())
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), domain.NewErrorResponse(err.Error()))
	}

	// Role In jwt because i'm got confuse to implement it
	userRoles := userLogin.Role
	if userLogin.Email == "admin@gmail.com" {
		userRoles = "admin"
	} else {
		userRoles = "user"
	}

	// JWT TOKEN
	token, _ := middleware.CreateTokenJWT(result.ID, result.Email, userRoles)
	result.Token = token

	user := response.LoginUseCase(&result)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Login Success", user))
}

func NewUserController(userUseCase domain.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}
