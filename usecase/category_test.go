package usecase_test

import (
	"prototype/domain"
	"prototype/usecase"
	mock_data "prototype/usecase/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostCategoryEquipment(t *testing.T) {
	mockRepo := &mock_data.MockCategoryRepository{
		Categories: make(map[int]*domain.CategoryEquipment),
	}
	useCase := usecase.NewCategoryUseCase(mockRepo)

	category := &domain.CategoryEquipment{Name: "Test Category"}
	expectedCat := category
	createdCat, err := useCase.PostCategoryEquipment(category)
	assert.NoError(t, err)
	assert.EqualValues(t, expectedCat, createdCat)

	_, err = useCase.PostCategoryEquipment(&domain.CategoryEquipment{})
	assert.Error(t, err, "Expected an error for empty input")
}

func TestGetAllCategories(t *testing.T) {
	mockRepo := new(mock_data.MockCategoryRepository)
	equipUseCase := usecase.NewCategoryUseCase(mockRepo)

	mockData := []*domain.CategoryEquipment{{ID: 1, Name: "Test 1"}, {ID: 2, Name: "Test 2"}}
	mockRepo.On("GetAll").Return(mockData, nil)
	result, err := equipUseCase.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, mockData, result)
}

func TestDeleteCategoryEquipment(t *testing.T) {
	mockRepo := &mock_data.MockCategoryRepository{
		Categories: map[int]*domain.CategoryEquipment{
			1: {ID: 1, Name: "Category 1"},
		},
	}
	useCase := usecase.NewCategoryUseCase(mockRepo)

	err := useCase.DeleteCategoryEquipment(1)
	assert.NoError(t, err)

	// Test deleting non-existent category
	err = useCase.DeleteCategoryEquipment(2)
	assert.Error(t, err, "Expected an error for deleting non-existent category")
}

func TestGetCategoryById(t *testing.T) {
	mockRepo := &mock_data.MockCategoryRepository{
		Categories: map[int]*domain.CategoryEquipment{
			1: {ID: 1, Name: "Category 1"},
		},
	}
	useCase := usecase.NewCategoryUseCase(mockRepo)

	cat, err := useCase.GetById(1)
	assert.NoError(t, err)
	assert.NotNil(t, cat)

	_, err = useCase.GetById(2)
	assert.Error(t, err, "Expected an error for getting non-existent category")
}
