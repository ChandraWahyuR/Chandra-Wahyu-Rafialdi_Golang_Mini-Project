package equipment

import (
	"net/http"
	"prototype/api/controller/equipment/request"
	"prototype/api/controller/equipment/response"
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

	equipmentResponses := make([]*response.EquipmentResponse, 0)
	for _, equip := range res {
		equipmentResponses = append(equipmentResponses, response.FromUseCase(equip))
	}
	return c.JSON(http.StatusOK, equipmentResponses)
}

func (uc *EquipmentController) PostEquipment(c echo.Context) error {
	var equip request.EquipmentRequest
	if err := c.Bind(&equip); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	imageURL, err := FetchImage(equip.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch equipment image"})
	}

	// Save to structi Equipment
	newEquipment := domain.Equipment{
		Name:        equip.Name,
		Category:    equip.Category,
		Description: equip.Description,
		Image:       imageURL,
		Price:       equip.Price,
	}

	resp, err := uc.equipmentUseCase.PostEquipment(&newEquipment)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), domain.NewErrorResponse(err.Error()))
	}

	// Show response with format from response folder
	equipmentResponse := response.FromUseCase(&resp)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Create data Success", equipmentResponse))
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

	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Delete Sucsess", equipment))
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
	equipmentResponse := response.FromUseCase(equipment)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Sucsess", equipmentResponse))
}

func NewEquipmentController(equipmentUseCase domain.EquipmentUseCaseInterface) *EquipmentController {
	return &EquipmentController{
		equipmentUseCase: equipmentUseCase,
	}
}
