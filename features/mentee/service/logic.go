package service

import (
	"be13/project/features/mentee"
	"errors"

	"github.com/go-playground/validator/v10"
)

type menteeService struct {
	menteeRepository mentee.RepositoryInterface
	validate         *validator.Validate
}

func NewMentee(repo mentee.RepositoryInterface) mentee.ServiceInterface {
	return &menteeService{
		menteeRepository: repo,
		validate:         validator.New(),
	}
}

// AddMentee implements mentee.ServiceInterface
func (service *menteeService) AddMentee(input mentee.Core) error {
	err := service.menteeRepository.AddMentee(input)
	if err != nil {
		return errors.New("failed to add class")
	}
	return nil
}

// DeleteMentee implements mentee.ServiceInterface
func (service *menteeService) DeleteMentee(id int) (mentee.Core, error) {
	data, err := service.menteeRepository.DeleteMentee(id)
	return data, err
}

// GetAllmentee implements mentee.ServiceInterface
func (service *menteeService) GetAllmentee() (data []mentee.Core, err error) {
	data, err = service.menteeRepository.GetAllmentee()
	return data, err
}

// GetMentebyParam implements mentee.ServiceInterface
func (service *menteeService) GetMentebyParam(name string, status string, category string, class uint) ([]mentee.Core, error) {
	data, err := service.menteeRepository.GetMentebyParam(name, status, category, class)
	return data, err
}

// UpdateMentee implements mentee.ServiceInterface
func (service *menteeService) UpdateMentee(id int, input mentee.Core) error {
	err := service.menteeRepository.UpdateMentee(id, input)
	if err != nil {
		return errors.New("id not found")
	}
	return nil
}
