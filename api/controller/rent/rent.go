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

	response := response.FromUseCase(&resp)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessInsert, response))
}

func (uc *RentController) DeleteRent(c echo.Context) error {
	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}
	rent := uc.rentusecase.DeleteRent(id)

	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessDelete, rent))
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
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: "Serv"})
	}

	// Check if rent is confirmed, if yes then it cannot be updated again
	if rentToUpdate.RentConfirmID != 0 {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: "Data Cannot be update"})
	}

	equipment, err := uc.equipmentusecase.GetById(rentToUpdate.EquipmentId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrDataNotFound.Error()})
	}

	// Update data
	totalRent := equipment.Price * rent.Quantity
	rentData := domain.Rent{
		Quantity: rent.Quantity,
		Total:    totalRent,
	}

	updateRent, err := uc.rentusecase.UpdateRent(id, &rentData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, constant.ErrUpdateData)
	}
	respon := response.FromUseCase(updateRent)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse(constant.SuccessUpdate, respon))
}

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
