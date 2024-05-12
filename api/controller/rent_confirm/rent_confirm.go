package rentconfirm

import (
	"net/http"
	"os"
	"prototype/api/controller/rent_confirm/request"
	"prototype/api/controller/rent_confirm/response"
	md "prototype/api/middleware"
	"prototype/constant"
	"prototype/domain"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type RentConfirmController struct {
	rentconfirmUseCase domain.RentConfirmUseCaseInterface
	rentUseCase        domain.RentUseCaseInterface
	//
	userusecase      domain.UseCaseInterface
	equipmentusecase domain.EquipmentUseCaseInterface
}

// @Tags Confirmation Rental
// @Summary Post Data Confirmiation
// @Description Post Data Confirmiation
// @ID Post-Confirmiation
// @Produce json
// @Success 200 {object} response.RentConfirmRespond
// @Failure 400
// @Router /confirm [post]
func (uc *RentConfirmController) PostRentConfirm(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	userID, _, _, err := md.ExtractToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, domain.BaseErrorResponse{Status: false, Message: "Unauthorized"})
	}

	var conf request.PostConfirmRequest
	if err := c.Bind(&conf); err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrFetchData.Error()})
	}

	imageFile, err := c.FormFile("payment_method")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetImage.Error()})
	}
	// Upload a image
	folder := os.Getenv("CLOUDINARY_UPLOAD_BUKTI")
	imageURL, err := md.GetImage(imageFile, folder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrFetchImage.Error()})
	}

	// Logic for delivery same as input
	delivery := conf.Delivery
	if delivery {
		if conf.Address == "" {
			return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "Address is required for delivery"})
		}
	}

	// Took Rent data that need to be confirmed
	rents, err := uc.rentUseCase.GetUnconfirmedRents(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "Failed to get unconfirmed rents"})
	}

	// Check if data rent is avaible
	if len(rents) == 0 {
		return c.JSON(http.StatusNotFound, domain.BaseErrorResponse{Status: false, Message: "Error, rent data not found"})
	}

	rentsData := make([]domain.Rent, len(rents))
	for i, rent := range rents {
		rentsData[i] = *rent
	}

	// Data from equipemt and user for show data in response
	user, err := uc.userusecase.GetByID(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "User not found"})
	}

	status := domain.StatusPending
	confirmData := domain.RentConfirm{
		UserId:        userID,
		User:          *user,
		Duration:      conf.Duration,
		PaymentMethod: imageURL,
		Delivery:      &conf.Delivery,
		Address:       conf.Address,
		DateStart:     time.Now(),
		Status:        status,
		Rents:         rentsData,
	}

	for i, rent := range rents {
		confirmData.Rents[i] = *rent
	}

	// Tambahan
	equipementsData := make([]domain.Rent, len(rents))
	for i, rent := range rents {
		rentData := *rent
		equipment, err := uc.equipmentusecase.GetById(rentData.EquipmentId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: "Failed to get equipment"})
		}

		rentData.Equipment = *equipment
		rentsData[i] = rentData
	}

	confirmedRent, err := uc.rentconfirmUseCase.PostRentConfirm(&confirmData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: "Failed to confirm rent"})
	}
	confirmData.Rents = equipementsData
	//

	rentResponse := response.FromUseCase(&confirmedRent)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Create data Success", rentResponse))
}

func (uc *RentConfirmController) GetById(c echo.Context) error {
	confirmId := c.Param("id")
	id, err := strconv.Atoi(confirmId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "id not found"})
	}

	rentConfirm, err := uc.rentconfirmUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}
	rentConfirmResponse := response.FromUseCase(rentConfirm)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Success", rentConfirmResponse))
}

// @Tags Confirmation Rental
// @Summary Get All Data Confirmiation
// @Description Get All Data Confirmiation
// @ID Get-All-Confirmiation
// @Produce json
// @Success 200 {object} response.RentConfirmRespond
// @Failure 400
// @Router /confirm [get]
func (uc *RentConfirmController) GetAll(c echo.Context) error {
	res, err := uc.rentconfirmUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetDatabase.Error()})
	}

	respon := make([]*response.RentConfirmRespond, 0)
	for _, respond := range res {
		respon = append(respon, response.FromUseCase(respond))
	}
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Success", respon))
}

// New Feature
// Get User Rents Confirmation by User id

// @Tags Confirmation Rental
// @Summary Get All Data Confirmation for User
// @Description User Get All Data Confirmiation
// @ID Get-All-Confirmiation-For-User
// @Produce json
// @Success 200 {object} response.RentConfirmRespond
// @Failure 400
// @Router /confirm/user [get]
func (uc *RentConfirmController) FindRentConfirmByUserId(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	// Ekstract userID from JWT
	userID, _, _, err := md.ExtractToken(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(constant.ErrById.Error()))
	}

	rentConfirms, err := uc.rentconfirmUseCase.FindRentConfirmByUserId(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.NewErrorResponse(constant.ErrFindData.Error()))
	}

	// Conversi Data rent confirm to rent confirm response
	rentResponses := make([]*response.RentConfirmRespond, len(rentConfirms))
	for i, rent := range rentConfirms {
		rentResponses[i] = response.FromUseCase(rent)
	}

	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Success", rentResponses))
}

// Admin
// @Tags Confirmation Rental
// @Summary Confirmation Data Rental
// @Description Confirmation Data Rental by Admin
// @ID Confirmiation-Rental
// @Produce json
// @Success 200 {object} response.RentConfirmRespond
// @Failure 400
// @Router /admin/confirm/{id} [put]
func (uc *RentConfirmController) ConfirmAdmin(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	adminID, _, _, err := md.ExtractToken(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(constant.ErrById.Error()))
	}

	var conf request.RentConfirmRequest
	if err := c.Bind(&conf); err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetDatabase.Error()})
	}

	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}

	rentConfirm, err := uc.rentconfirmUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetDataFromId.Error()})
	}

	if rentConfirm.Status != domain.StatusPending {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "rent confirmation has already been confirmed"})
	}

	now := time.Now()
	returnDate := now.Add(time.Duration(conf.Duration) * 7 * 24 * time.Hour)

	rentConfirmData := domain.RentConfirm{
		AdminId:    adminID,
		Status:     conf.Status,
		ReturnTime: returnDate,
	}

	confirmedRent, err := uc.rentconfirmUseCase.ConfirmAdmin(id, &rentConfirmData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: "failed to confirm rent"})
	}

	rentResponse := response.FromUseCase(confirmedRent)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessUpdate, rentResponse))
}

// @Tags Confirmation Rental
// @Summary Cancel Data Confirmation for User
// @Description User Can cancel confirmation before it got reviewed by admin
// @ID Cancel-Data-Confirmation-For-User
// @Produce json
// @Success 200 {object} response.RentConfirmRespond
// @Failure 400
// @Router /confirm/user/{id} [delete]
func (uc *RentConfirmController) CancelRentConfirmByUserId(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	userID, _, _, err := md.ExtractToken(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(constant.ErrById.Error()))
	}

	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}

	rentConfirm, err := uc.rentconfirmUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetDataID.Error()})
	}

	// If user didnt have data rent confirmation
	if rentConfirm.UserId != userID {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "rent confirmation not found"})
	}

	conf := uc.rentconfirmUseCase.CancelRentConfirmByUserId(id, userID)
	if conf != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: "failed to cancel rent confirmation"})
	}

	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Rent confirmation cancelled successfully", nil))
}

// Rental Info

// @Tags Info data of the user who rented
// @Summary Get All Info Returned and not returned
// @Description Admin can see all info user that have Returned and not returned
// @ID Info-Rerutned
// @Produce json
// @Success 200 {object} response.RentalInfoRespond
// @Failure 400
// @Router /admin/info [get]
func (uc *RentConfirmController) GetAllInfoRental(c echo.Context) error {
	res, err := uc.rentconfirmUseCase.GetAllInfoRental()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: err.Error()})
	}

	respon := make([]*response.RentalInfoRespond, 0)
	for _, respond := range res {
		respon = append(respon, response.FromUseCaseInfo(respond))
	}
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Success", respon))
}

// @Tags Info data of the user who rented
// @Summary Update Data Returned and not returned equipment
// @Description Admin can confirm that user have returned equipment and not returned it
// @ID Confirmation-Rerutned-data
// @Produce json
// @Success 200 {object} response.RentalInfoRespond
// @Failure 400
// @Router /admin/info{id} [put]
func (uc *RentConfirmController) ConfirmReturnRental(c echo.Context) error {
	var conf request.RentConfirmRequest
	if err := c.Bind(&conf); err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "Error getting data"})
	}

	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}

	rentConfirm, err := uc.rentconfirmUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}

	rentConfirm.Status = conf.Status

	confirmedRent, err := uc.rentconfirmUseCase.ConfirmReturnRental(id, rentConfirm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: "Failed to confirm rent"})
	}

	// Convert confirmed rent to response format
	rentResponse := response.FromUseCaseInfo(confirmedRent)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessUpdate, rentResponse))
}

func NewRentConfirmController(confirm domain.RentConfirmUseCaseInterface, rent domain.RentUseCaseInterface, user domain.UseCaseInterface, equipment domain.EquipmentUseCaseInterface) *RentConfirmController {
	return &RentConfirmController{
		rentconfirmUseCase: confirm,
		rentUseCase:        rent,
		//
		userusecase:      user,
		equipmentusecase: equipment,
	}
}
