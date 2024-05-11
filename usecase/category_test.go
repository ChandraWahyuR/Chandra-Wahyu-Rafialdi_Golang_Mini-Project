package usecase_test

import (
	"prototype/domain"
	"prototype/usecase"
	"prototype/usecase/mock"
	mock_data "prototype/usecase/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostCategoryEquipment(t *testing.T) {
	mockRepo := &mock.MockCategoryRepository{
		Categories: make(map[int]*domain.CategoryEquipment),
	}
	useCase := usecase.NewCategoryUseCase(mockRepo)

	category := &domain.CategoryEquipment{Name: "Test Category"}
	mockRepo.On("PostCategoryEquipment", category).Return(*category, nil)

	createdCat, err := useCase.PostCategoryEquipment(category)

	assert.NoError(t, err)
	assert.Equal(t, category, &createdCat)

	_, err = useCase.PostCategoryEquipment(&domain.CategoryEquipment{})
	assert.Error(t, err, "Expected an error for empty input")
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
