package equipment

import (
	"net/http"
	"os"
	"prototype/api/controller/equipment/request"
	"prototype/api/controller/equipment/response"
	"prototype/api/middleware"
	"prototype/constant"
	"prototype/domain"
	"prototype/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EquipmentController struct {
	equipmentUseCase domain.EquipmentUseCaseInterface
	categoryUseCase  domain.CategoryEquipmentUseCaseInterface
}

func (uc *EquipmentController) GetAll(c echo.Context) error {
	res, err := uc.equipmentUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrDataNotFound.Error()})
	}

	equipmentResponses := make([]*response.EquipmentResponse, 0)
	for _, equip := range res {
		equipmentResponses = append(equipmentResponses, response.FromUseCase(equip))
	}
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Sucsess", equipmentResponses))
}

func (uc *EquipmentController) PostEquipment(c echo.Context) error {
	var equip request.EquipmentRequest
	if err := c.Bind(&equip); err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrFetchData.Error()})
	}

	imageFile, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrGetImage.Error()})
	}

	// Upload a image
	folder := os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
	imageURL, err := middleware.GetImage(imageFile, folder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrFetchImage.Error()})
	}

	// Retrieve category by id
	category, err := uc.categoryUseCase.GetById(equip.CategoryId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}

	// Save data to struct Equipment
	newEquipment := domain.Equipment{
		Name:        equip.Name,
		CategoryId:  equip.CategoryId,
		Category:    *category,
		Description: equip.Description,
		Image:       imageURL,
		Price:       equip.Price,
		Stock:       equip.Stock,
	}

	// Post equipment to database
	resp, err := uc.equipmentUseCase.PostEquipment(&newEquipment)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), domain.NewErrorResponse(err.Error()))
	}

	equipmentResponse := response.FromUseCase(&resp)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Create data Success", equipmentResponse))
}

func (uc *EquipmentController) DeleteEquipment(c echo.Context) error {
	equipmentID := c.Param("id")
	id, err := strconv.Atoi(equipmentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}

	equipment := uc.equipmentUseCase.DeleteEquipment(id)
	if equipment != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrDeleteData.Error()})
	}

	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Delete Sucsess", equipment))
}

func (uc *EquipmentController) GetById(c echo.Context) error {
	equipmentID := c.Param("id")
	id, err := strconv.Atoi(equipmentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BaseErrorResponse{Status: false, Message: constant.ErrFetchData.Error()})
	}

	equipment, err := uc.equipmentUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.BaseErrorResponse{Status: false, Message: constant.ErrById.Error()})
	}
	equipmentResponse := response.FromUseCase(equipment)

	equipmentResponses := make([]*response.EquipmentResponse, 0)
	equipmentResponses = append(equipmentResponses, response.FromUseCase(equipment))

	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Sucsess", equipmentResponse))
}

func NewEquipmentController(equipmentUseCase domain.EquipmentUseCaseInterface, category domain.CategoryEquipmentUseCaseInterface) *EquipmentController {
	return &EquipmentController{
		equipmentUseCase: equipmentUseCase,
		categoryUseCase:  category,
	}
}
