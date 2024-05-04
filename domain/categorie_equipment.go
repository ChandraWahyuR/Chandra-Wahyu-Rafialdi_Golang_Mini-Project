package domain

type CategoryEquipment struct {
	ID   int
	Name string
}

type CategoryEquipmentRepositoryInterface interface {
	PostCategoryEquipment(equip *CategoryEquipment) error
	GetAll() ([]*CategoryEquipment, error)
	DeleteCategoryEquipment(ID int) error
	GetById(ID int) (*CategoryEquipment, error)
}

// Logic related to user activity
type CategoryEquipmentUseCaseInterface interface {
	PostCategoryEquipment(*CategoryEquipment) (CategoryEquipment, error)
	GetAll() ([]*CategoryEquipment, error)
	DeleteCategoryEquipment(ID int) error
	GetById(ID int) (*CategoryEquipment, error)
}
