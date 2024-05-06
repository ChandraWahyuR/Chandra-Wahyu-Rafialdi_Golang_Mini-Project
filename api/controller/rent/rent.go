package rent

import (
	"net/http"
	"prototype/api/controller/rent/request"
	"prototype/api/controller/rent/response"
	md "prototype/api/middleware"
	"prototype/domain"
	"prototype/utils"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RentController struct {
	rentusecase      domain.RentUseCaseInterface
	equipmentusecase domain.EquipmentUseCaseInterface
}

func (uc *RentController) GetAll(c echo.Context) error {
	res, err := uc.rentusecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	respon := make([]*response.RentResponse, 0)
	for _, respond := range res {
		respon = append(respon, response.FromUseCase(respond))
	}
	return c.JSON(http.StatusOK, respon)
}

func (uc *RentController) PostRent(c echo.Context) error {
	// Take Jwt token for define user id
	// then insert it to rent user_id
	token := c.Request().Header.Get("Authorization")
	// fmt.Println("Received token:", token)
	userID, _, _, err := md.ExtractToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}

	var rent request.RentRequest
	if err := c.Bind(&rent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Took equipment data
	equipment, err := uc.equipmentusecase.GetById(rent.EquipmentId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Equipment not found"})
	}
	totalRent := equipment.Price * rent.Quantity
	rentData := domain.Rent{
		UserId:      userID,
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
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Create data Success", response))
}

func (uc *RentController) DeleteRent(c echo.Context) error {
	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "rent id not found"})
	}
	rent := uc.rentusecase.DeleteRent(id)

	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Delete Sucsess", rent))
}

func (uc *RentController) GetById(c echo.Context) error {
	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "rent id not found"})
	}

	rent, err := uc.rentusecase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid"})
	}
	resp, err := uc.rentusecase.PostRent(rent)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Sucsess", resp))
}

func (uc *RentController) UpdateRent(c echo.Context) error {
	// Bind to db
	var rent request.RentRequest
	if err := c.Bind(&rent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Get ID
	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "rent id not found"})
	}

	// Take equipment id from table
	// And equipment took the id to accsess price from table equipment id
	rentToUpdate, err := uc.rentusecase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to fetch rent data"})
	}

	equipment, err := uc.equipmentusecase.GetById(rentToUpdate.EquipmentId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Equipment not found"})
	}

	// Update data
	totalRent := equipment.Price * rent.Quantity
	rentData := domain.Rent{
		Quantity: rent.Quantity,
		Total:    totalRent,
	}

	updateRent, err := uc.rentusecase.UpdateRent(id, &rentData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	respon := response.FromUseCase(updateRent)
	return c.JSON(http.StatusOK, respon)
}

func (uc *RentController) GetByUserID(c echo.Context) error {
	userID := c.Param("id")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user_id format"})
	}

	rents, err := uc.rentusecase.GetUserID(parsedUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch rent data"})
	}
	rentResponses := make([]*response.RentResponse, len(rents))

	for i, rent := range rents {
		rentResponses[i] = response.FromUseCase(rent)
	}
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Sucsess", rentResponses))
}

func NewRentController(rentusecase domain.RentUseCaseInterface, equipment domain.EquipmentUseCaseInterface) *RentController {
	return &RentController{
		rentusecase:      rentusecase,
		equipmentusecase: equipment,
	}
}
