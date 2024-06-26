package rent

import (
	"net/http"
	"prototype/api/controller/rent/request"
	"prototype/api/controller/rent/response"
	md "prototype/api/middleware"
	"prototype/constant"
	"prototype/domain"
	"prototype/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RentController struct {
	rentusecase      domain.RentUseCaseInterface
	equipmentusecase domain.EquipmentUseCaseInterface
	userusecase      domain.UseCaseInterface
}

func (uc *RentController) GetAll(c echo.Context) error {
	res, err := uc.rentusecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "Data still Empty"})
	}

	respon := make([]*response.RentResponse, 0)
	for _, respond := range res {
		respon = append(respon, response.FromUseCase(respond))
	}
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessGetData, respon))
}

// @Tags Rent
// @Summary Get All Data Rent
// @Description Get All Data Rent
// @ID Post-Rent
// @Produce json
// @Success 200 {object} response.RentResponse
// @Failure 400
// @Router /rent [post]
func (uc *RentController) PostRent(c echo.Context) error {
	// Take Jwt token for define user id
	// then insert it to rent user_id
	token := c.Request().Header.Get("Authorization")
	// fmt.Println("Received token:", token)
	userID, _, _, err := md.ExtractToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, domain.BaseErrorResponse{Status: false, Message: "Error while extracting token"})
	}

	var rent request.RentRequest
	if err := c.Bind(&rent); err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrFetchData.Error()})
	}

	// Took equipment data
	equipment, err := uc.equipmentusecase.GetById(rent.EquipmentId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "Equipment not found"})
	}

	// Check equipment stock
	if rent.Quantity > equipment.Stock {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "Requested quantity exceeds from available stock"})
	}

	user, err := uc.userusecase.GetByID(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "User not found"})
	}

	totalRent := equipment.Price * rent.Quantity
	rentData := domain.Rent{
		UserId:      userID,
		User:        *user,
		EquipmentId: rent.EquipmentId,
		Equipment:   *equipment,
		Quantity:    rent.Quantity,
		Total:       totalRent,
	}

	resp, err := uc.rentusecase.PostRent(&rentData)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), domain.NewErrorResponse(err.Error()))
	}

	// Update equipment stock
	equipment.Stock -= rent.Quantity
	_, err = uc.equipmentusecase.UpdateEquipment(rent.EquipmentId, equipment)
	if err != nil {

		// Rollback data rent if fail
		uc.rentusecase.DeleteRent(resp.ID)
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: "Failed to update equipment stock"})
	}

	response := response.FromUseCase(&resp)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessInsert, response))
}

func (uc *RentController) DeleteRent(c echo.Context) error {
	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}

	// Get rent data before deleting
	rentToDelete, err := uc.rentusecase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetDataID.Error()})
	}

	equipment, err := uc.equipmentusecase.GetById(rentToDelete.EquipmentId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetDataID.Error()})
	}
	equipment.Stock += rentToDelete.Quantity

	// Update equipment stock
	_, err = uc.equipmentusecase.UpdateEquipment(rentToDelete.EquipmentId, equipment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrUpdateData.Error()})
	}
	// Delete rent data
	err = uc.rentusecase.DeleteRent(id)

	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessDelete, nil))
}

func (uc *RentController) GetById(c echo.Context) error {
	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}

	rent, err := uc.rentusecase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetDataID.Error()})
	}
	resp, err := uc.rentusecase.PostRent(rent)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessGetData, resp))
}

// @Tags Rent
// @Summary Update Data Rent
// @Description Update Data Rent
// @ID Update-Rent
// @Produce json
// @Success 200 {object} response.RentResponse
// @Failure 400
// @Router /rent/user{id} [put]
func (uc *RentController) UpdateRent(c echo.Context) error {
	// Bind to db
	var rent request.RentRequest
	if err := c.Bind(&rent); err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrFetchData.Error()})
	}

	// Get ID
	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetDataID.Error()})
	}

	// Take equipment id from table
	// And equipment took the id to accsess price from table equipment id
	rentToUpdate, err := uc.rentusecase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "id not found"})
	}

	// Check if rent is confirmed, if yes then it cannot be updated again
	if rentToUpdate.RentConfirmID != 0 {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "Data Cannot be update"})
	}

	equipment, err := uc.equipmentusecase.GetById(rentToUpdate.EquipmentId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrDataNotFound.Error()})
	}

	// Check if stock is got taken or decrease
	// Calculate difference in quantity
	quantityDifference := rent.Quantity - rentToUpdate.Quantity

	equipment.Stock -= quantityDifference

	if equipment.Stock < 0 {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "Insufficient equipment stock"})
	}

	// Update data
	totalRent := equipment.Price * rent.Quantity
	rentData := domain.Rent{
		Quantity: rent.Quantity,
		Total:    totalRent,
	}

	updateRent, err := uc.rentusecase.UpdateRent(id, &rentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, constant.ErrUpdateData)
	}
	_, err = uc.equipmentusecase.UpdateEquipment(rentToUpdate.EquipmentId, equipment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrUpdateData.Error()})
	}
	respon := response.FromUseCase(updateRent)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessUpdate, respon))
}

// @Tags Rent
// @Summary Get All Data Rent For User
// @Description Get All Data Rent For User
// @ID Get-User-Rent
// @Produce json
// @Success 200 {object} response.RentResponse
// @Failure 400
// @Router /rent/user [get]
func (uc *RentController) GetByUserID(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	userID, _, _, err := md.ExtractToken(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(constant.ErrById.Error()))
	}

	rents, err := uc.rentusecase.GetUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(constant.ErrFindData.Error()))
	}

	rentResponses := make([]*response.RentResponse, len(rents))
	for i, rent := range rents {
		rentResponses[i] = response.FromUseCase(rent)
	}
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessInsert, rentResponses))
}

func NewRentController(rentusecase domain.RentUseCaseInterface, equipment domain.EquipmentUseCaseInterface, user domain.UseCaseInterface) *RentController {
	return &RentController{
		rentusecase:      rentusecase,
		equipmentusecase: equipment,
		userusecase:      user,
	}
}
