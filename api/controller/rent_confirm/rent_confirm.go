package rentconfirm

import (
	"net/http"
	"prototype/api/controller/rent_confirm/request"
	"prototype/api/controller/rent_confirm/response"
	md "prototype/api/middleware"
	"prototype/constant"
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

	// Check if delivery is true
	if conf.Delivery != false && *&conf.Delivery == true {
		if conf.Address == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Address is required for delivery"})
		}
	}
	// Took Rent data that need to be confirmed
	rents, err := uc.rentUseCase.GetUnconfirmedRents(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Failed to get unconfirmed rents"})
	}

	// Check if data rent is avaible
	if len(rents) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Error, rent data not found"})
	}

	rentsData := make([]domain.Rent, len(rents))
	for i, rent := range rents {
		rentsData[i] = *rent
	}

	status := domain.StatusPending
	confirmData := domain.RentConfirm{
		UserId:        userID,
		Duration:      conf.Duration,
		PaymentMethod: conf.PaymentMethod,
		Delivery:      &conf.Delivery,
		Address:       conf.Address,
		Status:        status,
		Rents:         rentsData,
	}

	for i, rent := range rents {
		confirmData.Rents[i] = *rent
	}

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
	rentConfirmResponse := response.FromUseCase(rentConfirm)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Success", rentConfirmResponse))
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
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Success", respon))
}

// New Feature
// Get User Rents Confirmation by User id

func (uc *RentConfirmController) FindRentConfirmByUserId(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	// Ekstract userID from JWT
	userID, _, _, err := md.ExtractToken(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(constant.ErrById.Error()))
	}

	rentConfirms, err := uc.rentconfirmUseCase.FindRentConfirmByUserId(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(constant.ErrFindData.Error()))
	}

	// Conversi Data rent confirm to rent confirm response
	rentResponses := make([]*response.RentConfirmRespond, len(rentConfirms))
	for i, rent := range rentConfirms {
		rentResponses[i] = response.FromUseCase(rent)
	}

	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Success", rentResponses))
}

// Admin
func (uc *RentConfirmController) ConfirmAdmin(c echo.Context) error {
	var conf request.RentConfirmRequest
	if err := c.Bind(&conf); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Eror get data"})
	}

	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "rent id not found"})
	}

	rentConfirm, err := uc.rentconfirmUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid"})
	}

	if rentConfirm.Status != domain.StatusPending {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "rent confirmation has already been confirmed"})
	}

	rentData := domain.RentConfirm{
		Status: conf.Status,
	}

	confirmedRent, err := uc.rentconfirmUseCase.ConfirmAdmin(id, &rentData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to confirm rent"})
	}

	rentResponse := response.FromUseCase(confirmedRent)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessUpdate, rentResponse))
}

func NewRentConfirmController(confirm domain.RentConfirmUseCaseInterface, rent domain.RentUseCaseInterface) *RentConfirmController {
	return &RentConfirmController{
		rentconfirmUseCase: confirm,
		rentUseCase:        rent,
	}
}
