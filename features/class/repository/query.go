package repository

import (
	"be13/project/features/class"
	"errors"

	"gorm.io/gorm"
)

type classRepository struct {
	db *gorm.DB
}

func NewClass(db *gorm.DB) class.RepositoryInterface {
	return &classRepository{
		db: db,
	}
}

// AddClass implements class.RepositoryInterface
func (repo *classRepository) AddClass(input class.Core) error {
	classGorm := CoretoModel(input)
	tx := repo.db.Create(&classGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil
}

// DeleteClass implements class.RepositoryInterface
func (*classRepository) DeleteClass(id int) error {
	panic("unimplemented")
}

// GetAllClass implements class.RepositoryInterface
func (repo *classRepository) GetAllClass() (data []class.Core, err error) {
	var class []Class
	tx := repo.db.Find(&class)
	if tx.Error != nil {
		return nil, tx.Error
	}
	res := toCoreList(class) // model to core
	return res, nil
}

// GetClassbyId implements class.RepositoryInterface
func (repo *classRepository) GetClassbyId(name string) (class.Core, error) {
	classes := Class{}
	tx := repo.db.Where("name", name).Find(&classes)
	if tx.Error != nil {
		return class.Core{}, tx.Error
	}
	result := classes.ModeltoCore()
	return result, nil
}

// UpdateClass implements class.RepositoryInterface
func (*classRepository) UpdateClass(id int, input class.Core) error {
	panic("unimplemented")
}
