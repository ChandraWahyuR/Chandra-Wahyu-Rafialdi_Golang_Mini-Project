package rentconfirm

import (
	"net/http"
	"prototype/api/controller/rent_confirm/request"
	"prototype/api/controller/rent_confirm/response"
	md "prototype/api/middleware"
	"prototype/domain"
	"strconv"

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

	// Mendapatkan data rent dari user berdasarkan userID
	rents, err := uc.rentUseCase.GetUserID(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User id from rent not found"})
	}
	var convertedRents []domain.Rent
	for _, rent := range rents {
		convertedRents = append(convertedRents, *rent)
	}
	status := domain.StatusPending
	confirmData := domain.RentConfirm{
		UserId:        userID,
		Duration:      conf.Duration,
		PaymentMethod: conf.PaymentMethod,
		Delivery:      &conf.Delivery,
		Address:       conf.Address,
		Status:        status,
		Rents:         convertedRents,
	}

	// Mengonversi slice []*domain.Rent menjadi []domain.Rent
	for i, rent := range rents {
		confirmData.Rents[i] = *rent
	}

	// Menyimpan rent_confirm
	confirmedRent, err := uc.rentconfirmUseCase.PostRentConfirm(&confirmData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to confirm rent"})
	}

	rentResponse := response.FromUseCase(&confirmedRent)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Create data Success", rentResponse))
}

func (uc *RentConfirmController) GetById(c echo.Context) error {
	confirmId := c.Param("id")
	id, err := strconv.Atoi(confirmId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id not found"})
	}

	rentConfirm, err := uc.rentconfirmUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid"})
	}

	return c.JSON(http.StatusOK, rentConfirm)
}

func (uc *RentConfirmController) GetAll(c echo.Context) error {
	res, err := uc.rentconfirmUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	respon := make([]*response.RentConfirmRespond, 0)
	for _, respond := range res {
		respon = append(respon, response.FromUseCase(respond))
	}
	return c.JSON(http.StatusOK, respon)
}

func NewRentConfirmController(confirm domain.RentConfirmUseCaseInterface, rent domain.RentUseCaseInterface) *RentConfirmController {
	return &RentConfirmController{
		rentconfirmUseCase: confirm,
		rentUseCase:        rent,
	}
}
