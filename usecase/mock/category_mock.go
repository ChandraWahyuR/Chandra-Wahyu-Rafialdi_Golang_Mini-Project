package mock

import (
	"errors"
	"prototype/domain"

	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	mock.Mock
	Categories map[int]*domain.CategoryEquipment
}

func NewMockCategoryRepository() *MockCategoryRepository {
	return &MockCategoryRepository{
		Categories: make(map[int]*domain.CategoryEquipment),
	}
}

func (m *MockCategoryRepository) PostCategoryEquipment(equip *domain.CategoryEquipment) error {
	if equip == nil {
		return errors.New("category is nil")
	}
	m.Categories[equip.ID] = equip

	return nil
}

func (m *MockCategoryRepository) GetAll() ([]*domain.CategoryEquipment, error) {
	categories := make([]*domain.CategoryEquipment, 0, len(m.Categories))
	for _, cat := range m.Categories {
		categories = append(categories, cat)
	}
	return categories, nil
}

func (m *MockCategoryRepository) DeleteCategoryEquipment(id int) error {
	if _, ok := m.Categories[id]; !ok {
		return errors.New("category not found")
	}
	delete(m.Categories, id)
	return nil
}

func (m *MockCategoryRepository) GetById(id int) (*domain.CategoryEquipment, error) {
	cat, ok := m.Categories[id]
	if !ok {
		return nil, errors.New("category not found")
	}
	return cat, nil
}
