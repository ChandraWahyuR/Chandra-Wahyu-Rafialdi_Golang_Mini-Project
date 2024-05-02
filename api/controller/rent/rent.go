package rent

import (
	"net/http"
	"prototype/api/controller/rent/request"
	"prototype/domain"
	"prototype/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RentController struct {
	rentusecase domain.RentUseCaseInterface
}

func (uc *RentController) GetAll(c echo.Context) error {
	res, err := uc.rentusecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (uc *RentController) PostRent(c echo.Context) error {
	var rent request.RentRequest
	if err := c.Bind(&rent); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	rentData := domain.Rent{
		Quantity:  rent.Quantity,
		Total:     rent.Total,
		DateStart: rent.DateStart,
		Duration:  rent.Duration,
	}

	resp, err := uc.rentusecase.PostRent(&rentData)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), domain.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Create data Success", resp))
}

func (uc *RentController) DeleteRent(c echo.Context) error {
	rentID := c.Param("id")
	id, err := strconv.Atoi(rentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "rent id not found"})
	}
	rent := uc.rentusecase.DeleteRent(id)

	if rent != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid"})
	}
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

	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Sucsess", rent))
}

func NewRentController(rentusecase domain.RentUseCaseInterface) *RentController {
	return &RentController{
		rentusecase: rentusecase,
	}
}
