package rentconfirm

import (
	"net/http"
	"prototype/api/controller/rent_confirm/request"
	"prototype/api/controller/rent_confirm/response"
	md "prototype/api/middleware"
	"prototype/domain"

	"github.com/labstack/echo/v4"
)

type RentConfirmController struct {
	rentconfirmUseCase domain.RentConfirmUseCaseInterface
	rentUseCase        domain.RentUseCaseInterface
}

func (uc *RentConfirmController) PostRentConfirm(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	userID, _, _, err := md.ExtractToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	var conf request.RentConfirmRequest
	if err := c.Bind(&conf); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	rents, err := uc.rentUseCase.GetUserID(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User id from rent not found"})
	}

	confirmData := domain.RentConfirm{
		UserId:        userID,
		PaymentMethod: conf.PaymentMethod,
		Delivery:      &conf.Delivery,
		Address:       conf.Address,
		Status:        conf.Status,
		Rents:         rents,
	}

	confirmedRent, err := uc.rentconfirmUseCase.PostRentConfirm(&confirmData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to confirm rent"})
	}

	// Mengonversi rent confirm yang berhasil ke dalam bentuk response
	rentResponse := response.FromUseCase(&confirmedRent)

	// Mengembalikan data rent confirm yang telah di-confirm
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Create data Success", rentResponse))
}

func NewRentConfirmController(confirm domain.RentConfirmUseCaseInterface, rent domain.RentUseCaseInterface) *RentConfirmController {
	return &RentConfirmController{
		rentconfirmUseCase: confirm,
		rentUseCase:        rent,
	}
}
