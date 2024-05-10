package category

import (
	"net/http"

	"prototype/api/controller/category/request"
	"prototype/api/controller/category/response"
	"prototype/domain"
	"prototype/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase domain.CategoryEquipmentUseCaseInterface
}

func (uc *CategoryController) GetAll(c echo.Context) error {
	res, err := uc.categoryUseCase.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	equipmentResponses := make([]*response.CategoryResponse, 0)
	for _, equip := range res {
		equipmentResponses = append(equipmentResponses, response.FromUseCase(equip))
	}
	return c.JSON(http.StatusOK, equipmentResponses)
}

func (uc *CategoryController) PostCategory(c echo.Context) error {
	var equip request.CategoryRequest
	if err := c.Bind(&equip); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newEquipment := domain.CategoryEquipment{
		Name: equip.Name,
	}

	resp, err := uc.categoryUseCase.PostCategoryEquipment(&newEquipment)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), domain.NewErrorResponse(err.Error()))
	}

	equipmentResponse := response.FromUseCase(&resp)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Create data Success", equipmentResponse))
}

func (uc *CategoryController) DeleteCategory(c echo.Context) error {
	equipmentID := c.Param("id")
	id, err := strconv.Atoi(equipmentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "equipment id not found"})
	}

	equipment := uc.categoryUseCase.DeleteCategoryEquipment(id)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Delete Sucsess", equipment))
}

func (uc *CategoryController) GetById(c echo.Context) error {
	categoryID := c.Param("id")
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "category id not found"})
	}

	category, err := uc.categoryUseCase.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid"})
	}
	resp, err := uc.categoryUseCase.PostCategoryEquipment(category)
	return c.JSON(http.StatusOK, domain.NewSuccessResponse("Get Data Sucsess", resp))
}

func NewCategoryController(categoryUseCase domain.CategoryEquipmentUseCaseInterface) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUseCase,
	}
}
