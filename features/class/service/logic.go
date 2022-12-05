package service

import (
	"be13/project/features/class"
	"errors"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classRepository class.RepositoryInterface
	validate        *validator.Validate
}

func NewClass(repo class.RepositoryInterface) class.ServiceInterface {
	return &classService{
		classRepository: repo,
		validate:        validator.New(),
	}
}

// AddClass implements class.ServiceInterface
func (service *classService) AddClass(input class.Core) error {
	err := service.classRepository.AddClass(input)
	if err != nil {
		return errors.New("failed to add class")
	}
	return nil
}

// DeleteClass implements class.ServiceInterface
func (*classService) DeleteClass(id int) error {
	panic("unimplemented")
}

// GetAllClass implements class.ServiceInterface
func (service *classService) GetAllClass() (data []class.Core, err error) {
	data, err = service.classRepository.GetAllClass()
	return data, err
}

// GetClassbyId implements class.ServiceInterface
func (service *classService) GetClassbyId(name string) (class.Core, error) {
	data, errData := service.classRepository.GetClassbyId(name)
	if errData != nil {
		return class.Core{}, errors.New("error get id")
	}
	return data, nil
}

// UpdateClass implements class.ServiceInterface
func (*classService) UpdateClass(id int, input class.Core) error {
	panic("unimplemented")
}
