package equipment

import (
	"net/http"
	"prototype/api/controller/equipment/request"
	"prototype/domain"
	"prototype/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EquipmentController struct {
	equipmentUseCase domain.EquipmentUseCaseInterface
}

func (uc *EquipmentController) GetAll(c echo.Context) error {
	res, err := uc.equipmentUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (uc *EquipmentController) PostEquipment(c echo.Context) error {
	var equip request.EquipmentRequest
	if err := c.Bind(&equip); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newEquipment := domain.Equipment{
		Name:        equip.Name,
		Category:    equip.Category,
		Description: equip.Description,
		Image:       equip.Image,
	}

	resp, err := uc.equipmentUseCase.PostEquipment(&newEquipment)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), domain.NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Create data Success", resp))
}

func (uc *EquipmentController) DeleteEquipment(c echo.Context) error {
	equipmentID := c.Param("id")
	id, err := strconv.Atoi(equipmentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "equipment id not found"})
	}

	equipment := uc.equipmentUseCase.DeleteEquipment(id)
	if equipment != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid"})
	}
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Delete Sucsess", nil))
}

func (uc *EquipmentController) GetById(c echo.Context) error {
	equipmentID := c.Param("id")
	id, err := strconv.Atoi(equipmentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "equipment id not found"})
	}

	equipment, err := uc.equipmentUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid"})
	}

	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Sucsess", equipment))
}

func NewEquipmentController(equipmentUseCase domain.EquipmentUseCaseInterface) *EquipmentController {
	return &EquipmentController{
		equipmentUseCase: equipmentUseCase,
	}
}
